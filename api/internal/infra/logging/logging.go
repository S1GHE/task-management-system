package logging

import "github.com/sirupsen/logrus"

func MustLog(mode string) *logrus.Logger {
	log := logrus.New()

	switch mode {
	case "dev":
		log.SetLevel(logrus.DebugLevel)
	case "prod":
		log.SetLevel(logrus.InfoLevel)
	}

	return log
}
