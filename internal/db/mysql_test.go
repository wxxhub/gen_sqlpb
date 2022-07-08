package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDbInfo(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(www.wxxhome.com:3306)/test")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var schema string

	err = db.QueryRow("SELECT SCHEMA()").Scan(&schema)
	if err != nil {
		panic(err)
	}
	fmt.Println("schema:", schema)
	tableName := "new_table"
	//var desc xstring
	rows, err := db.Query(fmt.Sprintf("SHOW FULL COLUMNS FROM %s", tableName))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//desc := make(map[xstring]xstring)
	for rows.Next() {
		c := new(Columns)
		fmt.Println(rows.Columns())
		rows.Scan(&c.Field, &c.Type, &c.Collation, &c.Null, &c.Key, &c.Default, &c.Extra, &c.Privileges, &c.Comment)
		fmt.Println("c:", c)
		fmt.Println("err:", err)
		fmt.Println("desc: ", *c)
	}
}
