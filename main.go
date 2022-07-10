package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/common"
	"github.com/wxxhub/gen_sqlpb/internal/db"
	"github.com/wxxhub/gen_sqlpb/internal/flag"
	"github.com/wxxhub/gen_sqlpb/internal/gen"
	"github.com/wxxhub/gen_sqlpb/internal/xstring"
	"os"

	"strings"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorln(r)
		}
	}()
	gloabalConfig := flag.ParseFlag()
	// set log level
	if gloabalConfig.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Debugln("gloabalConfig: ", gloabalConfig)

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

		tableInfo.UpperName = strings.ToUpper(tableInfo.Name)
		tableInfo.CamelName = xstring.ToCamelWithStartUpper(tableInfo.Name)
		tableInfo.FName = strings.ToLower(string(tableInfo.Name[0]))

		gen.GenTemples(srvConfig, tableInfo, gloabalConfig.Option.Temples)
		//gen.GenProto(srvConfig, tableInfo)
	}
}
