package flag

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/config"
	"strings"
)

type Option struct {
	SrvName   string   `long:"srvName" description:"service name"`
	SavePath  string   `long:"savePath" description:"protobuf save path"`
	DSN       []string `long:"dsn" description:"data source name"`
	Debug     bool     `long:"debug" description:"print debug info"`
	Package   string   `long:"package" description:"protobuf package"`
	GoPackage string   `long:"goPackage" description:"golang package"`
	FileName  string   `long:"fileName" description:"protobuf file name"`
}

func parseTableConfig(dsn string) *config.SqlConfig {
	a := strings.Split(dsn, "?")
	sqlDsn := a[0]
	paramMap := make(map[string]string)
	if len(a) > 1 {
		paramsStr := a[1]
		params := strings.Split(paramsStr, "&")
		for _, item := range params {
			t := strings.Split(item, "=")
			paramMap[t[0]] = t[1]
		}
	}

	//params := ""
	//for key, value := range paramMap {
	//	switch key {
	//	case "tableName":
	//		params = fmt.Sprintf("%s?%s=%s", params, key, value)
	//	}
	//}
	//if len(params) > 0 {
	//	sqlDsn = fmt.Sprintf("%s?%s", sqlDsn, params)
	//}
	tableName := paramMap["tableName"]

	c := &config.SqlConfig{
		SqlDsn:    sqlDsn,
		TableName: tableName,
	}

	if srvName, ok := paramMap["srvName"]; ok {
		c.SrvName = srvName
	}

	return c
}

func ParseFlag() (globalConfig *config.GlobalConfig) {
	globalConfig = new(config.GlobalConfig)
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorln("ParseFlag err:", r)
		}
	}()

	var opt Option
	_, err := flags.Parse(&opt)
	if err != nil {
		fmt.Println("err:", err)
	}

	globalConfig.Debug = opt.Debug
	globalConfig.Services = make(map[string]*config.ServiceConfig)

	for _, item := range opt.DSN {
		c := parseTableConfig(item)
		srvName := opt.SrvName

		if "" != c.SrvName {
			srvName = c.SrvName
		}

		if _, ok := globalConfig.Services[srvName]; ok {
			globalConfig.Services[srvName].SqlConfigs[c.TableName] = c
		} else {
			globalConfig.Services[srvName] = new(config.ServiceConfig)
			globalConfig.Services[srvName].SrvName = srvName
			globalConfig.Services[srvName].SavePath = opt.SavePath
			//globalConfig.Services[srvName].GoPackage = opt.GoPackage
			//globalConfig.Services[srvName].Package = opt.Package
			globalConfig.Services[srvName].SqlConfigs = map[string]*config.SqlConfig{c.TableName: c}
		}
	}

	return globalConfig
}
