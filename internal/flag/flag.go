package flag

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/wxxhub/gen_sqlpb/internal/config"
	"strings"
)

type Option struct {
	SrvName  string   `short:"s" long:"srvName" description:"service name"`
	SavePath string   `long:"savePath" description:"service name"`
	DSN      []string `short:"d" long:"dsn" description:"dsn"`
	Debug    bool     `long:"debug" description:"debug"`
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

	return &config.SqlConfig{
		SqlDsn:    sqlDsn,
		TableName: tableName,
	}
}

func ParseFlag() *config.GenConfig {
	var opt Option
	_, err := flags.Parse(&opt)
	if err != nil {
		fmt.Println("err:", err)
	}

	genConfig := new(config.GenConfig)
	genConfig.SrvName = opt.SrvName
	genConfig.SavePath = opt.SavePath
	genConfig.Debug = opt.Debug

	genConfig.SqlConfigs = make([]*config.SqlConfig, len(opt.DSN))

	for i, item := range opt.DSN {
		genConfig.SqlConfigs[i] = parseTableConfig(item)
	}

	return genConfig
}