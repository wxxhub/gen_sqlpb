package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/db"
	"github.com/wxxhub/gen_sqlpb/internal/flag"
	"github.com/wxxhub/gen_sqlpb/internal/gen"
	"os"
	"strings"
)

func main() {
	genConfig := flag.ParseFlag()

	// set log level
	if genConfig.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// get data struct
	colsMap := make(map[string][]*db.Columns)
	for _, item := range genConfig.SqlConfigs {
		cols, err := db.GenerateSchema("mysql", item.SqlDsn, item.TableName)
		if err != nil {
			logrus.Panicf("GenerateSchema faile: %s", err.Error())
		}
		colsMap[item.TableName] = cols

		logrus.Infof("gen table %s", item.TableName)
	}

	// config
	if len(genConfig.Package) == 0 {
		genConfig.Package = strings.ToLower(genConfig.SrvName)
	}
	if len(genConfig.GoPackage) == 0 {
		genConfig.GoPackage = strings.ToLower(genConfig.SrvName)
	}
	if len(genConfig.SavePath) > 0 {
		err := os.MkdirAll(genConfig.SavePath, os.ModePerm)
		if err != nil {
			logrus.Panicf("mkdir %s faile:%s", genConfig.SavePath, err.Error())
		}
	}
	if len(genConfig.FileName) == 0 {
		genConfig.FileName = genConfig.SrvName + ".proto"
	}

	// gen proto
	gen.GenProto(genConfig, colsMap)
}
