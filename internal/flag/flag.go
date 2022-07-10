package flag

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/common"
	"strings"
)

func parseTableConfig(dsn string) *common.DbConfig {
	a := strings.Split(dsn, "?")
	dsnN := a[0]
	paramMap := make(map[string]string)
	if len(a) > 1 {
		paramsStr := a[1]
		params := strings.Split(paramsStr, "&")
		for _, item := range params {
			t := strings.Split(item, "=")
			paramMap[t[0]] = t[1]
		}
	}

	tableName := paramMap["tableName"]

	c := &common.DbConfig{
		Dsn:       dsnN,
		TableName: tableName,
	}

	if srvName, ok := paramMap["srvName"]; ok {
		c.SrvName = srvName
	}

	// database
	dsnNSplit := strings.Split(dsnN, "/")
	if len(dsnNSplit) == 2 {
		c.DataBase = dsnNSplit[1]
	} else {
		logrus.Panicln("dsn need database")
	}

	return c
}

func ParseFlag() (globalConfig *common.GlobalConfig) {
	globalConfig = new(common.GlobalConfig)
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorln("ParseFlag err:", r)
		}
	}()

	var opt common.Option
	_, err := flags.Parse(&opt)
	if err != nil {
		fmt.Println("err:", err)
	}

	globalConfig.Debug = opt.Debug
	globalConfig.Services = make(map[string]*common.ServiceConfig)

	for _, item := range opt.DSN {
		dbConfig := parseTableConfig(item)
		if "" == dbConfig.SrvName {
			dbConfig.SrvName = dbConfig.TableName
		}

		srvName := dbConfig.SrvName

		globalConfig.Services[srvName] = new(common.ServiceConfig)
		globalConfig.Services[srvName].SrvName = srvName
		globalConfig.Services[srvName].SavePath = opt.SavePath
		globalConfig.Services[srvName].StructSavePath = opt.SavePath
		globalConfig.Services[srvName].SqlSavePath = opt.SavePath
		//globalConfig.Services[srvName].GoPackage = opt.GoPackage
		//globalConfig.Services[srvName].Package = opt.Package
		globalConfig.Services[srvName].DbConfig = dbConfig
		globalConfig.Services[srvName].SrvName = dbConfig.SrvName

		if len(globalConfig.Services[srvName].SavePath) == 0 {
			globalConfig.Services[srvName].SavePath = "./proto"
		}

		if len(globalConfig.Services[srvName].StructSavePath) == 0 {
			globalConfig.Services[srvName].StructSavePath = "./struct"
		}

		if len(globalConfig.Services[srvName].SqlSavePath) == 0 {
			globalConfig.Services[srvName].SqlSavePath = "./sql"
		}
	}

	return globalConfig
}
