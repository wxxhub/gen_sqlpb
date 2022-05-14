package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
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

func GenerateSchema(driverName, dsn, tableName string) ([]*Columns, error) {
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorln("GenerateSchema err:", r)
		}
	}()

	switch driverName {
	case "mysql":
		return GenerateMysqlSchema(dsn, tableName)
	}

	return nil, fmt.Errorf("implement %s", driverName)
}
