package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	gen "github.com/wxxhub/gen_sql_pb/internal"
	idb "github.com/wxxhub/gen_sql_pb/internal/db"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(www.wxxhome.com:3306)/test")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	cols, err := idb.GenerateSchema(db, "")
	if err != nil {
		panic(err)
	}

	gen.GenProto(cols, "TestService", "Test", "cache")
}
