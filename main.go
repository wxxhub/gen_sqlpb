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
	gloabalConfig := flag.ParseFlag()
	// set log level
	if gloabalConfig.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Debugln("gloabalConfig: ", gloabalConfig)

	gen.Gen(gloabalConfig)
}
