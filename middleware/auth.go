package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Auth:", ctx.Request.Host)
		ctx.Next()
	}
}
