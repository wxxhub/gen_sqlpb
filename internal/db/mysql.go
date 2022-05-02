package db

import (
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
)

type Columns struct {
	Field      string
	Type       string
	Collation  string
	Null       string
	Key        string
	Default    string
	Extra      string
	Privileges string
	Comment    string
}

func GenerateSchema(db *sql.DB, table string) ([]*Columns, error) {
	cols := make([]*Columns, 0)

	if len(table) == 0 {
		return cols, nil
	}

	rows, err := db.Query("SHOW FULL COLUMNS FROM new_table")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := new(Columns)
		rows.Scan(&c.Field, &c.Type, &c.Collation, &c.Null, &c.Key, &c.Default, &c.Extra, &c.Privileges, &c.Comment)
		cols = append(cols, c)
	}

	return cols, nil
}
