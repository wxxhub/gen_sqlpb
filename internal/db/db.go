package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/common"
)

func GenerateSchema(driverName, dsn, dataBaseName, tableName string) (*common.TableInfo, error) {
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorln("GenerateSchema err:", r)
		}
	}()

	switch driverName {
	case "mysql", "mariadb":
		return GenerateMysqlSchema(dsn, dataBaseName, tableName)
	}

	return nil, fmt.Errorf("implement %s", driverName)
}
