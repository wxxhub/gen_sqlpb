package main

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	gen "github.com/wxxhub/gen_sqlpb/internal"
	idb "github.com/wxxhub/gen_sqlpb/internal/db"
)

type SqlConfig struct {
	sqlDsn    string
	tableName string
}

type GenConfig struct {
	sqlConfigs []*SqlConfig
	srvName    string
	savePath   string
}

type Option struct {
	SrvName  string   `short:"s" long:"srvName" description:"service name"`
	SavePath string   `long:"savePath" description:"service name"`
	DSN      []string `short:"d" long:"dsn" description:"dsn"`
}

func parseTableConfig(dsn string) *SqlConfig {
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

	return &SqlConfig{
		sqlDsn:    sqlDsn,
		tableName: tableName,
	}
}

func parseFlag() *GenConfig {
	var opt Option
	_, err := flags.Parse(&opt)
	if err != nil {
		fmt.Println("err:", err)
	}

	genConfig := new(GenConfig)
	genConfig.srvName = opt.SrvName
	genConfig.savePath = opt.SavePath

	genConfig.sqlConfigs = make([]*SqlConfig, len(opt.DSN))

	for i, item := range opt.DSN {
		genConfig.sqlConfigs[i] = parseTableConfig(item)
	}

	return genConfig
}

func main() {
	genConfig := parseFlag()

	colsMap := make(map[string][]*idb.Columns)
	for _, item := range genConfig.sqlConfigs {
		db, err := sql.Open("mysql", item.sqlDsn)
		if err != nil {
			panic(err)
		}

		defer db.Close()

		cols, err := idb.GenerateSchema(db, item.tableName)
		if err != nil {
			log.Panic(err)
			//panic(err)
		}

		colsMap[item.tableName] = cols
	}

	path := genConfig.srvName + ".proto"
	if len(genConfig.savePath) > 0 {
		path = filepath.Join(genConfig.savePath, path)
	}

	gen.GenProto(colsMap, genConfig.srvName, path)
}
