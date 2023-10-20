package middleware

import (
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func DomainMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url, err := url.Parse(ctx.Request.RequestURI)
		if err != nil {
			panic(err)
		}

		host := url.Hostname()

		parts := strings.Split(host, ".")
		if len(parts) >= 2 {
			ctx.Next()
		} else {
			panic("Invalid")
		}
	}
}
