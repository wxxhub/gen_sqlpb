package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/common"
	"log"
	"regexp"
	"strings"
)

func GenerateMysqlSchema(dsn, database, table string) (*common.TableInfo, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if len(table) == 0 {
		return nil, fmt.Errorf("table is empty")
	}

	// query createTable
	queryCreateTableRow, err := db.Query("SHOW CREATE TABLE " + table)
	if err != nil {
		return nil, err
	}
	defer queryCreateTableRow.Close()

	createTableInfo := ""
	tn := ""
	for queryCreateTableRow.Next() {
		queryCreateTableRow.Scan(&tn, &createTableInfo)
	}

	tableInfo, err := parseCreateTable(table, createTableInfo)
	if err != nil {
		return nil, err
	}

	return tableInfo, err
}

func parseCreateTable(table, createTable string) (*common.TableInfo, error) {
	defer func() {

	}()

	if len(createTable) == 0 {
		return nil, fmt.Errorf("createTable empty")
	}

	tableInfo := new(common.TableInfo)
	tableInfo.Name = table
	tableInfo.CreateTable = createTable
	tableInfo.Columns = make([]*common.Column, 0)
	tableInfo.SqlIndexes = make([]*common.SqlIndex, 0)

	fieldMap := make(map[string]*common.Column)

	lines := strings.Split(createTable, "\n")
	endI := len(lines) - 1

	for i := 1; i < endI; i++ {
		trimLine := strings.TrimSpace(lines[i])
		if '`' == trimLine[0] { // 	字段
			colum := parseCreateTableField(trimLine)

			fieldMap[colum.Field] = colum
			tableInfo.Columns = append(tableInfo.Columns, colum)
		} else { // 索引
			idx := parseCreateTableIndexField(trimLine)
			if idx.Type == common.PrimaryIndexKey {
				tableInfo.PrimaryIndex = idx
			} else {
				tableInfo.SqlIndexes = append(tableInfo.SqlIndexes)
			}
		}
	}

	return tableInfo, nil
}

func parseCreateTableField(line string) *common.Column {
	fieldData := strings.Split(line, " ")

	if len(fieldData) < 3 {
		logrus.Panicf("field error")
	}

	colum := new(common.Column)
	colum.Field = strings.Trim(fieldData[0], "`")
	colum.Type = fieldData[1]
	//TODO colum.Collation
	upperLine := strings.ToUpper(line)

	if strings.Contains(upperLine, "UNSIGNED") {
		colum.Type = colum.Type + " unsigned"
	}

	colum.Null = "NO"
	if strings.Contains(upperLine, "NOT NULL") {
		colum.Null = "YES"
	}

	defaultValueReg := regexp.MustCompile("DEFAULT ([^ ]*)|DEFAULT (\"[^\"]\")|DEFAULT ('[^']*')")
	matchIndex := defaultValueReg.FindStringSubmatchIndex(upperLine)
	if len(matchIndex) > 0 {
		colum.Default = line[matchIndex[2]:matchIndex[3]]
	}

	commentValueReg := regexp.MustCompile("COMMENT '([^']*)'")
	matchIndex = commentValueReg.FindStringSubmatchIndex(upperLine)
	if len(matchIndex) > 0 {
		colum.Comment = line[matchIndex[2]:matchIndex[3]]
	}

	return colum
}

func parseCreateTableIndexField(line string) *common.SqlIndex {
	fieldData := strings.Split(line, " ")

	sqlIndex := new(common.SqlIndex)
	sqlIndex.Type = common.IndexKey

	if len(fieldData) < 3 {
		logrus.Panicf("field error")
	}

	upperTrimLine := strings.ToUpper(line)
	if strings.Contains(upperTrimLine, "KEY") {
		// 索引名及字段
		s := strings.Split(line, "(")
		if len(s) != 2 {
			log.Panic("index err:", line)
		}

		reg := regexp.MustCompile("`([a-zA-Z0-9_-]*)`")
		sqlIndex.KeyName = strings.Trim(reg.FindString(s[0]), "`")

		fields := reg.FindAllString(s[1], -1)
		for _, item := range fields {
			sqlIndex.ColumnName = append(sqlIndex.ColumnName, strings.Trim(item, "`"))
		}

		// 索引种类
		if strings.Contains(upperTrimLine, "PRIMARY KEY") {
			sqlIndex.Type = common.PrimaryIndexKey
			sqlIndex.KeyName = "PRIMARY"
		} else if strings.Contains(upperTrimLine, "UNIQUE KEY") {
			sqlIndex.Type = common.UniqueIndexKey
		}

		// 索引方式
		reg = regexp.MustCompile("USING ([^ ]*)")
		matchStrings := reg.FindStringSubmatch(upperTrimLine)
		if len(matchStrings) > 1 {
			sqlIndex.IndexType = matchStrings[1]
		}
	}

	return sqlIndex
}

/*
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

	tableInfo.PrimaryIndex, tableInfo.SqlIndexes = parseIndexByCreateTable(tableInfo.CreateTable)

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
*/
