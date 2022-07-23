package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wxxhub/gen_sqlpb/internal/common"
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
		c := new(common.Column)
		fmt.Println(rows.Columns())
		rows.Scan(&c.Field, &c.Type, &c.Collation, &c.Null, &c.Key, &c.Default, &c.Extra, &c.Privileges, &c.Comment)
		fmt.Println("c:", c)
		fmt.Println("err:", err)
		fmt.Println("desc: ", *c)
	}
}

func TestParsePrimaryKey(t *testing.T) {
	line := "PRIMARY KEY (`id`)"
	index := parseCreateTableIndexField(line)
	fmt.Printf("%+v\n", index)
	fmt.Println("==============")

	line = "PRIMARY KEY (`id`,`keyword`)"
	index = parseCreateTableIndexField(line)
	fmt.Printf("%+v\n", index)
	fmt.Println("==============")

	line = "UNIQUE KEY `catename` (`catid`)  "
	index = parseCreateTableIndexField(line)
	fmt.Printf("%+v\n", index)
	fmt.Println("==============")

	line = "KEY `idx_id_name` (`id`, `name`)  "
	index = parseCreateTableIndexField(line)
	fmt.Printf("%+v\n", index)
	fmt.Println("==============")
}

func TestParseField(t *testing.T) {
	line := "`id` bigint(20) NOT NULL DEFAULT '123' AUTO_INCREMENT comment 'test comment',"
	c := parseCreateTableField(line)
	fmt.Printf("%+v\n", c)
}
