package system

import (
	"github.com/sirupsen/logrus"
)

func init() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	logrus.SetLevel(logrus.DebugLevel)
}

func Startup() (err error) {
	//todo
	return
}
