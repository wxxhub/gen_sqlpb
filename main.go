package main

import (
	"database/sql"
	"flag"
	"log"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	gen "github.com/wxxhub/gen_sqlpb/internal"
	idb "github.com/wxxhub/gen_sqlpb/internal/db"
)

type SqlConfig struct {
	sqldsn    string
	tableName string
}

type GenConfig struct {
	sqlConfigs  []*SqlConfig
	serviceName string
	savePath    string
}

func parseTableConfig(dsn string) *SqlConfig {
	a := strings.Split(dsn, "?")
	p := a[0]

	if len(a) > 1 {
		paramStr := a[1]
	}
	
}

func parseFlag() *GenConfig {
	config := &GenConfig{
		sqlConfigs: make([]*SqlConfig, 0),
	}
	flag.Parse()

	for i := 0; i < flag.NArg(); i++ {
		switch flag.Arg(i) {
		case "dsn":
			config.sqlConfigs = append(config.sqlConfigs, parseTableConfig(flag.Arg(i)))
		case "servive_name":
			config.serviceName = flag.Arg(i)
		case "save_path":
			config.savePath = flag.Arg(i)
		}
	}

	return config
}

func main() {
	dsn := flag.String("dsn", "", "the database dsn")
	serviceName := flag.String("srv", "", "the protobuf service name , defaults to the database schema.")
	savePath := flag.String("save_path", "", "the protobuf service name , defaults to the database schema.")
	tableName := flag.String("table_name", "", "the protobuf service name , defaults to the database schema.")

	//packageName := flag.String("package", *schema, "the protocol buffer package. defaults to the database schema.")
	//goPackageName := flag.String("go_package", "", "the protocol buffer go_package. defaults to the database schema.")

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	cols, err := idb.GenerateSchema(db, *tableName)
	if err != nil {
		log.Panic(err)
		//panic(err)
	}

	path := *serviceName + ".proto"
	if len(*savePath) > 0 {
		path = filepath.Join(*savePath, path)
	}

	gen.GenProto(cols, *serviceName, *tableName, path)
}
