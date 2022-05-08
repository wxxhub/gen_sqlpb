package db

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func GenerateMysqlSchema(dsn, table string) ([]*Columns, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	cols := make([]*Columns, 0)

	if len(table) == 0 {
		logrus.Warnf("table is empty, dsn: %s", dsn)
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
