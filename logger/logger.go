package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	logFileName := "logs/schedule_api_" + time.Now().Format("2006-01-02_15") + ".log"

	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatal("Erro in creating log file: ", err)
	}

	Log.SetOutput(logFile)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetLevel(logrus.InfoLevel)
}
