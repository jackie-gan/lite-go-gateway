package gateway

import "github.com/gin-gonic/gin"

func NewGateway() *gin.Engine {
	engine := gin.Default()

	engine.Any("/")

	return engine
}
