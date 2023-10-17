package gateway

import (
	"lite-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func NewGateway() *gin.Engine {
	engine := gin.Default()

	engine.Any("/ws", middleware.ProxyMiddleware())

	return engine
}
