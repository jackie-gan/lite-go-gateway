package main

import (
	"lite-gateway/config"
	"lite-gateway/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if config.InitConfig() != nil {
		os.Exit(1)
	}

	gw := gin.Default()

	paths := config.GetConfig().ProxyPaths

	for _, path := range paths {
		gw.Any(path.Prefix, middleware.ProxyMiddleware(path.Url))
	}

	if err := gw.Run(":8080"); err != nil {
		panic(err)
	}
}
