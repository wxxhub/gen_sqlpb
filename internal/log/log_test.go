package log

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogrus(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Debugln("This is debug")
	logrus.Infoln("This is Info")
	logrus.Warning("This is Warn")
	logrus.Errorln("This is Error")

	logrus.Infof("\n\nset WarnLevel\n\n")
	logrus.SetLevel(logrus.WarnLevel)

	logrus.Debugln("This is debug2")
	logrus.Infoln("This is Info2")
	logrus.Warning("This is Warn2")
	logrus.Errorln("This is Error2")
}
