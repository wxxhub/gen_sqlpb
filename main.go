package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/db"
	"github.com/wxxhub/gen_sqlpb/internal/flag"
	"github.com/wxxhub/gen_sqlpb/internal/gen"
	"os"
	"path/filepath"
)

func main() {
	genConfig := flag.ParseFlag()

	if genConfig.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	colsMap := make(map[string][]*db.Columns)
	for _, item := range genConfig.SqlConfigs {
		cols, err := db.GenerateSchema("mysql", item.SqlDsn, item.TableName)
		if err != nil {
			logrus.Panicf("GenerateSchema sql faile: %s", err.Error())
		}
		colsMap[item.TableName] = cols

		logrus.Infof("gen table %s", item.TableName)
	}

	path := genConfig.SrvName + ".proto"
	if len(genConfig.SavePath) > 0 {
		err := os.MkdirAll(genConfig.SavePath, os.ModePerm)
		if err != nil {
			logrus.Panicf("mkdir %s faile:%s", genConfig.SavePath, err.Error())
		}

		path = filepath.Join(genConfig.SavePath, path)
	}

	gen.GenProto(colsMap, genConfig.SrvName, path)
}
