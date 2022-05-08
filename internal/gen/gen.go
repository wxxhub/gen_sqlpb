package gen

import (
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/config"
	"github.com/wxxhub/gen_sqlpb/internal/db"
	"github.com/wxxhub/gen_sqlpb/internal/xstring"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	_ "embed"
)

//go:embed template/proto.tpl
var protoTpl string

type Table struct {
	Name      string
	UpperName string
	Columns   map[string]string
}

type Content struct {
	Srv       string
	Tables    []*Table
	Package   string
	GoPackage string
}

func GenProto(genConfig *config.GenConfig, colsMap map[string][]*db.Columns) {

	tables := make([]*Table, 0)
	for tableName, item := range colsMap {
		tables = append(tables, &Table{
			Name:      tableName,
			UpperName: xstring.ToCamelWithStartUpper(tableName),
			Columns:   genTableContent(item),
		})
	}

	content := &Content{
		Srv:       genConfig.SrvName,
		Tables:    tables,
		Package:   genConfig.Package,
		GoPackage: genConfig.GoPackage,
	}

	tmpl, err := template.New("gen_proto").Parse(protoTpl)
	if err != nil {
		logrus.Panicf("Parse proto template faile: %s", err.Error())
	}

	fullPath := filepath.Join(genConfig.SavePath, genConfig.FileName)
	f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		logrus.Panicf("OpenFile %s faile: %s", fullPath, err.Error())
	}
	defer f.Close()

	err = tmpl.Execute(f, content)
	if err != nil {
		logrus.Panicf("Execute template faile:%s", err.Error())
	}
}

func genTableContent(cols []*db.Columns) map[string]string {
	m := make(map[string]string)
	for _, item := range cols {
		itemType := strings.Split(item.Type, "(")[0]
		switch itemType {
		case "char", "varchar", "text", "longtext", "mediumtext", "tinytext", "enum", "set":
			m[item.Field] = "xstring"
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
			m[item.Field] = "xstring"
			logrus.Warnf("%s use default type xstring", itemType)
		}
	}

	logrus.Debugf("genTableContent: %+v", m)
	return m
}
