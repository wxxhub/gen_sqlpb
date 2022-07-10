package gen

import (
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/common"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	_ "embed"
)

//go:embed template/proto.tpl
var protoTpl string

//go:embed template/struct.tpl
var structTpl string

func GenTemples(serviceConfig *common.ServiceConfig, tableInfo *common.TableInfo, tplFiles []string) {
	for _, tplFile := range tplFiles {
		switch tplFile {
		case common.DefaultProtoFileName:
			fullPath := filepath.Join(serviceConfig.SavePath, serviceConfig.FileName)
			GenTemple(serviceConfig, tableInfo, protoTpl, fullPath)
		case common.DefaultStructFileName:
			fullPath := filepath.Join(serviceConfig.SavePath, serviceConfig.StructFileName)
			GenTemple(serviceConfig, tableInfo, structTpl, fullPath)
		default:
			GenTempleFromFile(serviceConfig, tableInfo, tplFile, "")
		}
	}

	// sql
	sqlFullPath := filepath.Join(serviceConfig.SqlSavePath, serviceConfig.SqlFileName)
	ioutil.WriteFile(sqlFullPath, []byte(tableInfo.CreateTable), os.ModePerm)
}

func GenTempleFromFile(serviceConfig *common.ServiceConfig, tableInfo *common.TableInfo, tplFile string, saveFile string) {
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorln("GenTemple err:", r)
		}
	}()

	tpl, err := ioutil.ReadFile(tplFile)

	content := &common.Content{
		Srv:       serviceConfig.SrvName,
		TableInfo: tableInfo,
		Package:   serviceConfig.Package,
		GoPackage: serviceConfig.GoPackage,
	}

	content.ProtoItems = genTableProtoContent(tableInfo)
	content.GoStructItems = genGoSturctContent(tableInfo)

	// proto
	tmpl, err := template.New("gen_temple").Parse(string(tpl))
	if err != nil {
		logrus.Panicf("Parse proto template faile: %s", err.Error())
	}

	f, err := os.OpenFile(saveFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		logrus.Panicf("OpenFile %s faile: %s", saveFile, err.Error())
	}
	defer f.Close()

	err = tmpl.Execute(f, content)
	if err != nil {
		logrus.Panicf("Execute template faile:%s", err.Error())
	}
}

func GenTemple(serviceConfig *common.ServiceConfig, tableInfo *common.TableInfo, tpl string, saveFile string) {
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorln("GenTemple err:", r)
		}
	}()

	content := &common.Content{
		Srv:       serviceConfig.SrvName,
		TableInfo: tableInfo,
		Package:   serviceConfig.Package,
		GoPackage: serviceConfig.GoPackage,
	}

	content.ProtoItems = genTableProtoContent(tableInfo)
	content.GoStructItems = genGoSturctContent(tableInfo)

	tmpl, err := template.New("gen_temple").Parse(tpl)
	if err != nil {
		logrus.Panicf("Parse proto template faile: %s", err.Error())
	}

	f, err := os.OpenFile(saveFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		logrus.Panicf("OpenFile %s faile: %s", saveFile, err.Error())
	}
	defer f.Close()

	err = tmpl.Execute(f, content)
	if err != nil {
		logrus.Panicf("Execute template faile:%s", err.Error())
	}
}

func genTableProtoContent(tableInfo *common.TableInfo) []*common.ProtoItem {
	m := make([]*common.ProtoItem, len(tableInfo.Columns))
	for index, item := range tableInfo.Columns {
		m[index] = &common.ProtoItem{
			Index:     index + 1,
			Name:      item.Field,
			NameUpper: strings.ToUpper(item.Field),
		}

		itemType := strings.Split(item.Type, "(")[0]
		switch itemType {
		case "char", "varchar", "text", "longtext", "mediumtext", "tinytext", "enum", "set":
			m[index].Type = "string"
		case "blob", "mediumblob", "longblob", "varbinary", "binary":
			m[index].Type = "bytes"
		case "date", "time", "datetime", "timestamp":
			m[index].Type = "int64"
		case "bool":
			m[index].Type = "bool"
		case "tinyint", "smallint":
			if strings.Contains(item.Field, "unsigned") {
				m[index].Type = "uint32"
			} else {
				m[index].Type = "int32"
			}
		case "int", "mediumint", "bigint":
			if strings.Contains(item.Field, "unsigned") {
				m[index].Type = "uint64"
			} else {
				m[index].Type = "int64"
			}
		case "float", "decimal", "double":
			m[index].Type = "double"
		default:
			m[index].Type = "string"
			logrus.Warnf("%s use default type xstring", itemType)
		}
	}

	logrus.Debugf("genTableContent: %+v", m)
	return m
}

func genGoSturctContent(tableInfo *common.TableInfo) []*common.GoStructItem {
	m := make([]*common.GoStructItem, len(tableInfo.Columns))
	for index, item := range tableInfo.Columns {
		m[index] = &common.GoStructItem{
			Name:      item.Field,
			NameUpper: strings.ToUpper(item.Field),
			Column:    item,
		}

		itemType := strings.Split(item.Type, "(")[0]
		switch itemType {
		case "char", "varchar", "text", "longtext", "mediumtext", "tinytext", "enum", "set":
			m[index].Type = "string"
		case "blob", "mediumblob", "longblob", "varbinary", "binary":
			m[index].Type = "bytes"
		case "date", "time", "datetime", "timestamp":
			m[index].Type = "int64"
		case "bool":
			m[index].Type = "bool"
		case "tinyint", "smallint":
			if strings.Contains(item.Field, "unsigned") {
				m[index].Type = "uint32"
			} else {
				m[index].Type = "int32"
			}
		case "int", "mediumint", "bigint":
			if strings.Contains(item.Field, "unsigned") {
				m[index].Type = "uint64"
			} else {
				m[index].Type = "int64"
			}
		case "float", "decimal", "double":
			m[index].Type = "double"
		default:
			m[index].Type = "string"
			logrus.Warnf("%s use default type xstring", itemType)
		}
	}

	logrus.Debugf("genTableContent: %+v", m)
	return m
}
