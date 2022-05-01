package internal

import (
	"github.com/wxxhub/gen_sql_pb/internal/db"
	"html/template"
	"log"
	"os"
	"strings"

	_ "embed"
)

//go:embed template/proto.tpl
var protoTpl string

type Table struct {
	Name    string
	Columns map[string]string
}

type Content struct {
	Srv    string
	Tables []*Table
}

func GenProto(cols []*db.Columns, srv string, tableName string, savePath string) {
	table := &Table{
		Name:    tableName,
		Columns: genTableContent(cols),
	}

	content := &Content{
		Srv:    srv,
		Tables: []*Table{table},
	}

	//log.Println("cols:", cols)

	tmpl, err := template.New("test").Parse(protoTpl)
	if err != nil {
		panic(err)
	}
	//path := os.F
	f, err := os.OpenFile("test.proto", os.O_CREATE|os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	err = tmpl.Execute(f, content)
	if err != nil {
		panic(err)
	}
}

func genTableContent(cols []*db.Columns) map[string]string {
	m := make(map[string]string)
	for _, item := range cols {
		itemType := strings.Split(item.Type, "(")[0]
		switch itemType {
		case "char", "varchar", "text", "longtext", "mediumtext", "tinytext", "enum", "set":
			m[item.Field] = "string"
		case "blob", "mediumblob", "longblob", "varbinary", "binary":
			m[item.Field] = "bytes"
		case "date", "time", "datetime", "timestamp":
			m[item.Field] = "int64"
		case "bool":
			m[item.Field] = "bool"
		case "tinyint", "smallint":
			if strings.Contains(item.Field, "unsigned") {
				m[item.Field] = "uint32"
			} else {
				m[item.Field] = "int32"
			}
		case "int", "mediumint", "bigint":
			if strings.Contains(item.Field, "unsigned") {
				m[item.Field] = "uint64"
			} else {
				m[item.Field] = "int64"
			}
		case "float", "decimal", "double":
			m[item.Field] = "double"
		default:
			m[item.Field] = "string"
		}
	}

	log.Println("m:", m)
	return m
}
