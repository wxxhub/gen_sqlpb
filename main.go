package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wxxhub/gen_sqlpb/internal/flag"
	"github.com/wxxhub/gen_sqlpb/internal/gen"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			logrus.Errorln(r)
		}
	}()
	globalConfig := flag.ParseFlag()
	// set log level
	if globalConfig.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Debugln("globalConfig: ", globalConfig)

	gen.Gen(globalConfig)
}
