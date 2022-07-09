package gen

import (
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/common"
	"github.com/wxxhub/gen_sqlpb/internal/config"
	"github.com/wxxhub/gen_sqlpb/internal/xstring"
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

func GenProto(serviceConfig *config.ServiceConfig, tableInfo *common.TableInfo) {
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorln("GenProto err:", r)
		}
	}()

	tableInfo.UpperName = xstring.ToCamelWithStartUpper(tableInfo.Name)

	content := &common.Content{
		Srv:       serviceConfig.SrvName,
		TableInfo: tableInfo,
		Package:   serviceConfig.Package,
		GoPackage: serviceConfig.GoPackage,
	}

	content.ProtoItems = genTableProtoContent(tableInfo)
	content.GoStructItems = genGoSturctContent(tableInfo)

	// proto
	tmpl, err := template.New("gen_proto").Parse(protoTpl)
	if err != nil {
		logrus.Panicf("Parse proto template faile: %s", err.Error())
	}

	fullPath := filepath.Join(serviceConfig.SavePath, serviceConfig.FileName)
	f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		logrus.Panicf("OpenFile %s faile: %s", fullPath, err.Error())
	}
	defer f.Close()

	err = tmpl.Execute(f, content)
	if err != nil {
		logrus.Panicf("Execute template faile:%s", err.Error())
	}

	// go struct
	tmplStruct, err := template.New("gen_struct").Parse(structTpl)
	if err != nil {
		logrus.Panicf("Parse proto template faile: %s", err.Error())
	}

	fullPath = filepath.Join(serviceConfig.StructSavePath, serviceConfig.StructFileName)
	structF, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		logrus.Panicf("Open Struct File %s faile: %s", fullPath, err.Error())
	}
	defer structF.Close()

	err = tmplStruct.Execute(structF, content)
	if err != nil {
		logrus.Panicf("Execute template faile:%s", err.Error())
	}

	// sql
	fullPath = filepath.Join(serviceConfig.SqlSavePath, serviceConfig.SqlFileName)
	ioutil.WriteFile(fullPath, []byte(tableInfo.CreateTable), os.ModePerm)
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
