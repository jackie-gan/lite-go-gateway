package main

import (
	"lite-gateway/config"
	"lite-gateway/log"
	"lite-gateway/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

const serverPort = "8080"

func main() {
	log.InitLogger()
	defer log.Sync()

	if config.InitConfig() != nil {
		os.Exit(1)
	}

	gw := gin.Default()

	gw.Use(middleware.AuthMiddleware())

	paths := config.GetConfig().ProxyPaths
	for _, path := range paths {
		gw.Any(path.Prefix, middleware.ProxyMiddleware(path.Url))
	}

	log.Logger.Sugar().Infof("Server start at %s", serverPort)
	if err := gw.Run(":" + serverPort); err != nil {
		panic(err)
	}
}
