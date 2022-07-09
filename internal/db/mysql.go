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

	tempTable := ""
	for queryCreateTableRow.Next() {
		queryCreateTableRow.Scan(&tempTable, &tableInfo.CreateTable)
	}

	return tableInfo, nil
}
