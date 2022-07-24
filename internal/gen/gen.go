package gen

import (
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/common"
	"github.com/wxxhub/gen_sqlpb/internal/db"
	"github.com/wxxhub/gen_sqlpb/internal/xstring"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	_ "embed"
)

//go:embed template/proto.tpl
var protoTpl string

//go:embed template/struct.tpl
var structTpl string

func Gen(gloabalConfig *common.GlobalConfig) {
	if !gloabalConfig.Option.NotUseDefaultTemple {
		gloabalConfig.Option.Temples = append(gloabalConfig.Option.Temples, common.DefaultProtoFileName)
		gloabalConfig.Option.Temples = append(gloabalConfig.Option.Temples, common.DefaultStructFileName)
	}

	// get data struct
	for _, srvConfig := range gloabalConfig.Services {
		// config
		if len(srvConfig.Package) == 0 {
			srvConfig.Package = strings.ToLower(srvConfig.SrvName)
		}
		if len(srvConfig.GoPackage) == 0 {
			srvConfig.GoPackage = strings.ToLower(srvConfig.SrvName)
		}

		//mkdir
		if len(srvConfig.SavePath) > 0 {
			err := os.MkdirAll(srvConfig.SavePath, os.ModePerm)
			if err != nil {
				logrus.Panicf("mkdir %s faile:%s", srvConfig.SavePath, err.Error())
			}
		}
		if len(srvConfig.StructSavePath) > 0 {
			err := os.MkdirAll(srvConfig.StructSavePath, os.ModePerm)
			if err != nil {
				logrus.Panicf("mkdir %s faile:%s", srvConfig.StructSavePath, err.Error())
			}
		}
		if len(srvConfig.SqlSavePath) > 0 {
			err := os.MkdirAll(srvConfig.SqlSavePath, os.ModePerm)
			if err != nil {
				logrus.Panicf("mkdir %s faile:%s", srvConfig.StructSavePath, err.Error())
			}
		}

		//check fileName
		if len(srvConfig.FileName) == 0 {
			srvConfig.FileName = srvConfig.SrvName + ".proto"
		}
		if len(srvConfig.StructFileName) == 0 {
			srvConfig.StructFileName = srvConfig.SrvName + ".go"
		}
		if len(srvConfig.SqlFileName) == 0 {
			srvConfig.SqlFileName = srvConfig.SrvName + ".sql"
		}

		logrus.Debugln("srvConfig: ", srvConfig)
		tableInfo, err := db.GenerateSchema("mysql", srvConfig.DbConfig.Dsn, srvConfig.DbConfig.DataBase, srvConfig.DbConfig.TableName)
		if err != nil {
			logrus.Panicf("GenerateSchema faile: %s", err.Error())
		}

		genTemples(srvConfig, tableInfo, gloabalConfig.Option.Temples)
		//gen.GenProto(srvConfig, tableInfo)
	}
}

func genTemples(serviceConfig *common.ServiceConfig, tableInfo *common.TableInfo, tplFiles []string) {
	for _, tplFile := range tplFiles {
		switch tplFile {
		case common.DefaultProtoFileName:
			fullPath := filepath.Join(serviceConfig.SavePath, serviceConfig.FileName)
			genTemple(serviceConfig, tableInfo, protoTpl, fullPath)
		case common.DefaultStructFileName:
			fullPath := filepath.Join(serviceConfig.SavePath, serviceConfig.StructFileName)
			genTemple(serviceConfig, tableInfo, structTpl, fullPath)
		default:
			genTempleFromFile(serviceConfig, tableInfo, tplFile, "")
		}
	}

	// sql
	sqlFullPath := filepath.Join(serviceConfig.SqlSavePath, serviceConfig.SqlFileName)
	ioutil.WriteFile(sqlFullPath, []byte(tableInfo.CreateTable), os.ModePerm)
}

func genTempleFromFile(serviceConfig *common.ServiceConfig, tableInfo *common.TableInfo, tplFile string, saveFile string) {
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorf("genTemple err: %+v\n%s", r, string(debug.Stack()))
		}
	}()

	tpl, err := ioutil.ReadFile(tplFile)
	if err != nil {
		logrus.Panicf("ReadFile err: %+v", err)
	}

	genTemple(serviceConfig, tableInfo, string(tpl), saveFile)
}

func genTemple(serviceConfig *common.ServiceConfig, tableInfo *common.TableInfo, tpl string, saveFile string) {
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorf("genTemple err: %+v\n%s", r, string(debug.Stack()))
		}
	}()

	content := &common.Content{
		Srv:       serviceConfig.SrvName,
		TableInfo: tableInfo,
		Package:   serviceConfig.Package,
		GoPackage: serviceConfig.GoPackage,
	}

	content.ProtoContent = genTableProtoContent(tableInfo)
	content.GoStructContent = genGoStructContent(tableInfo)

	tmpl, err := template.New("gen_temple").Funcs(GetTemplateFuncList()).Parse(tpl)

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

func genTableProtoContent(tableInfo *common.TableInfo) *common.ProtoContent {
	content := new(common.ProtoContent)

	content.ProtoItems = make([]*common.ProtoItem, len(tableInfo.Columns))
	protoItemMap := make(map[string]*common.ProtoItem)
	for index, item := range tableInfo.Columns {
		content.ProtoItems[index] = &common.ProtoItem{
			GenItem: common.GenItem{
				Name:    item.Field,
				Comment: item.Comment,
			},
		}

		content.ProtoItems[index].Type = getProtoType(item)
		protoItemMap[content.ProtoItems[index].GenItem.Name] = content.ProtoItems[index]
	}

	// index
	content.PrimaryIndexItem = new(common.ProtoIndexItem)
	content.PrimaryIndexItem.Name = xstring.ToCamelWithStartUpper(tableInfo.PrimaryIndex.ColumnName[0])
	if len(tableInfo.PrimaryIndex.ColumnName) > 0 {
		for i := 1; i < len(tableInfo.PrimaryIndex.ColumnName); i++ {
			content.PrimaryIndexItem.Name = content.PrimaryIndexItem.Name + "_" + tableInfo.PrimaryIndex.ColumnName[i]
		}
	}

	content.PrimaryIndexItem.Fields = tableInfo.PrimaryIndex.ColumnName
	content.PrimaryIndexItem.IndexItems = make([]*common.ProtoItem, len(content.PrimaryIndexItem.Fields))
	for index, item := range content.PrimaryIndexItem.Fields {
		content.PrimaryIndexItem.IndexItems[index] = protoItemMap[item]
	}
	content.PrimaryIndexItem.Type = ""
	content.PrimaryIndexItem.Comment = tableInfo.PrimaryIndex.Comment

	logrus.Debugf("ProtoContent: %+v", content)
	return content
}

func getProtoType(col *common.Column) string {
	itemType := strings.Split(col.Type, "(")[0]

	switch itemType {
	case "char", "varchar", "text", "longtext", "mediumtext", "tinytext", "enum", "set":
		return "string"
	case "blob", "mediumblob", "longblob", "varbinary", "binary":
		return "bytes"
	case "date", "time", "datetime", "timestamp":
		return "int64"
	case "bool":
		return "bool"
	case "tinyint", "smallint":
		if strings.Contains(col.Type, "unsigned") {
			return "uint32"
		} else {
			return "int32"
		}
	case "int", "mediumint", "bigint":
		if strings.Contains(col.Type, "unsigned") {
			return "uint64"
		} else {
			return "int64"
		}
	case "float", "decimal", "double":
		return "double"
	default:
		return "string"
		logrus.Warnf("%s use default type xstring", col)
	}

	return ""
}

func genGoStructContent(tableInfo *common.TableInfo) *common.GoStructContent {
	content := new(common.GoStructContent)
	content.GoStructItems = make([]*common.GoStructItem, len(tableInfo.Columns))
	for index, item := range tableInfo.Columns {
		content.GoStructItems[index] = &common.GoStructItem{
			Column: item,
			GenItem: common.GenItem{
				Name:    item.Field,
				Comment: item.Comment,
			},
		}

		content.GoStructItems[index].Type = getGoStructType(item)
	}

	// index
	//content.PrimaryIndexItem = new(common.ProtoItem)
	//content.PrimaryIndexItem.Name = tableInfo.PrimaryIndex.Collation
	//content.PrimaryIndexItem.CamelName = xstring.ToCamelWithStartUpper(content.PrimaryIndexItem.Name)
	//content.PrimaryIndexItem.NameUpper = strings.ToUpper(content.PrimaryIndexItem.Name)
	//content.PrimaryIndexItem.Type = itemTypeMap[content.PrimaryIndexItem.Name]

	logrus.Debugf("genGoStructContent: %+v", content)
	return content
}

func getGoStructType(col *common.Column) string {
	itemType := strings.Split(col.Type, "(")[0]
	switch itemType {
	case "char", "varchar", "text", "longtext", "mediumtext", "tinytext", "enum", "set":
		return "string"
	case "blob", "mediumblob", "longblob", "varbinary", "binary":
		return "[]byte"
	case "date", "time", "datetime", "timestamp":
		return "int64"
	case "bool":
		return "bool"
	case "tinyint", "smallint":
		if strings.Contains(col.Type, "unsigned") {
			return "uint32"
		} else {
			return "int32"
		}
	case "int", "mediumint", "bigint":
		if strings.Contains(col.Type, "unsigned") {
			return "uint64"
		} else {
			return "int64"
		}
	case "float", "decimal", "double":
		return "float64"
	default:
		return "string"
		logrus.Warnf("%s use default type xstring", itemType)
	}

	return ""
}
