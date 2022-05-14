package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/db"
	"github.com/wxxhub/gen_sqlpb/internal/flag"
	"github.com/wxxhub/gen_sqlpb/internal/gen"
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
	serviceConfig := flag.ParseFlag()
	// set log level
	if serviceConfig.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Debugln("serviceConfig: ", serviceConfig)

	// get data struct
	for _, srvConfig := range serviceConfig.Services {
		colsMap := make(map[string][]*db.Columns)
		for _, item := range srvConfig.SqlConfigs {
			cols, err := db.GenerateSchema("mysql", item.SqlDsn, item.TableName)
			if err != nil {
				logrus.Panicf("GenerateSchema faile: %s", err.Error())
			}
			colsMap[item.TableName] = cols
		}

		// config
		if len(srvConfig.Package) == 0 {
			srvConfig.Package = strings.ToLower(srvConfig.SrvName)
		}
		if len(srvConfig.GoPackage) == 0 {
			srvConfig.GoPackage = strings.ToLower(srvConfig.SrvName)
		}

		if len(srvConfig.SavePath) > 0 {
			err := os.MkdirAll(srvConfig.SavePath, os.ModePerm)
			if err != nil {
				logrus.Panicf("mkdir %s faile:%s", srvConfig.SavePath, err.Error())
			}
		}
		if len(srvConfig.FileName) == 0 {
			srvConfig.FileName = srvConfig.SrvName + ".proto"
		}

		logrus.Debugln("srvConfig: ", srvConfig)
		gen.GenProto(srvConfig, colsMap)
	}
}
