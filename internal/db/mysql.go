package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/common"
	"strings"
)

func GenerateMysqlSchema(dsn, database, table string) (*common.TableInfo, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	tableInfo := new(common.TableInfo)
	tableInfo.Name = table
	tableInfo.UpperName = strings.ToUpper(table)
	tableInfo.Columns = make([]*common.Column, 0)

	if len(table) == 0 {
		logrus.Warnf("table is empty, dsn: %s", dsn)
		return tableInfo, nil
	}

	//quryString := "SELECT COLUMN_NAME, COLUMN_KEY, DATA_TYPE, IS_NULLABLE, COLUMN_TYPE, COLUMN_DEFAULT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND table_name = ?"
	queryColumsRows, err := db.Query("SHOW FULL COLUMNS FROM " + table)

	if err != nil {
		return nil, err
	}
	defer queryColumsRows.Close()

	for queryColumsRows.Next() {
		c := new(common.Column)
		queryColumsRows.Scan(&c.Field, &c.Type, &c.Collation, &c.Null, &c.Key, &c.Default, &c.Extra, &c.Privileges, &c.Comment)
		tableInfo.Columns = append(tableInfo.Columns, c)
	}

	// query createTable
	//queryCreateTableRows, err := db.Query()
	queryCreateTableRow, err := db.Query("SHOW CREATE TABLE " + table)
	if err != nil {
		return nil, err
	}
	defer queryCreateTableRow.Close()

	tempTable := ""
	for queryCreateTableRow.Next() {
		queryCreateTableRow.Scan(&tempTable, &tableInfo.CreateTable)
	}

	// query Index
	queryIndexTableRow, err := db.Query("SHOW INDEX FROM " + table)
	if err != nil {
		return nil, err
	}

	tableInfo.SqlIndexes = make([]*common.SqlIndex, 0)
	defer queryIndexTableRow.Close()
	for queryIndexTableRow.Next() {
		sqlIndex := new(common.SqlIndex)
		queryIndexTableRow.Scan(&sqlIndex.TableName,
			&sqlIndex.NonUnique,
			&sqlIndex.KeyName,
			&sqlIndex.SeqInIndex,
			&sqlIndex.ColumnName,
			&sqlIndex.Collation,
			&sqlIndex.Cardinality,
			&sqlIndex.SubPart,
			&sqlIndex.Packed,
			&sqlIndex.Null,
			&sqlIndex.IndexType,
			&sqlIndex.Comment,
			&sqlIndex.IndexComment,
			&sqlIndex.Expression,
			&sqlIndex.Visible)
		if "PRIMARY" == sqlIndex.KeyName {
			tableInfo.PrimaryIndex = sqlIndex
		} else {
			tableInfo.SqlIndexes = append(tableInfo.SqlIndexes, sqlIndex)
		}
	}

	return tableInfo, nil
}
