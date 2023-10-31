package log

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"lite-gateway/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func InitLogrus(log *logrus.Logger, loggerName string) {
	log.SetReportCaller(true)
	log.SetLevel(logrus.InfoLevel)
	if config.C != nil && config.C.Log.Path != "" {
		logPath := config.C.Log.Path
		if _, err := os.Stat(logPath); os.IsNotExist(err) {
			os.Mkdir(logPath, os.ModePerm)
		}
		loggerFile := path.Join(logPath, loggerName)
		if !strings.HasSuffix(loggerFile, ".log") {
			loggerFile = fmt.Sprintf("%s.log", loggerFile)
		}
		age := config.C.Log.Age
		if age <= 0 {
			age = 1
		}
		writer, _ := rotatelogs.New(
			loggerFile+".%Y%m%d%H",
			rotatelogs.WithLinkName(loggerFile),
			rotatelogs.WithMaxAge(time.Duration(24*age)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		)
		log.SetOutput(writer)
	} else {
		log.SetOutput(os.Stdout)
	}

}

func NewLogrus(loggerName string) *logrus.Logger {
	log := logrus.New()
	InitLogrus(log, loggerName)
	return log
}

var (
	Access *logrus.Logger
	Error  *logrus.Logger
	Main   *logrus.Logger
)

func InitLog() {
	Access = NewLogrus("access")
	Error = NewLogrus("error")
	Main = NewLogrus("main")
}
