package middleware

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyMiddleware(proxyTarget string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		target, err := url.Parse(proxyTarget)
		if err != nil {
			panic(err)
		}

		reverseProxy := httputil.NewSingleHostReverseProxy(target)

		reverseProxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
