package log

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	Logger = log
}

func Sync() {
	Logger.Sync()
}
