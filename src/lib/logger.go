package lib

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func LoggerInit(logger *logrus.Logger) {
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
}
