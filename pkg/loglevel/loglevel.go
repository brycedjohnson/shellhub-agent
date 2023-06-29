package loglevel

import (
	"github.com/sirupsen/logrus"
)

func SetLogLevel() {
	level := logrus.TraceLevel

	logrus.WithField("log_level", level.String()).Info("Setting log level")
	logrus.SetLevel(level)
}
