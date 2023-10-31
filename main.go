package main

import (
	"lite-gateway/config"
	"lite-gateway/log"
	"lite-gateway/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

const SERVER_PORT = "8080"

func before() {
	if config.InitConfig() != nil {
		os.Exit(1)
	}

	log.InitLog()
}

func main() {
	before()

	gw := gin.Default()

	gw.Use(middleware.AuthMiddleware())

	paths := config.C.ProxyPaths
	for _, path := range paths {
		gw.Any(path.Prefix, middleware.ProxyMiddleware(path.Url))
	}

	log.Main.Infof("Server start at %s", SERVER_PORT)
	if err := gw.Run(":" + SERVER_PORT); err != nil {
		panic(err)
	}
}
