package middleware

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyMiddleware(proxyTarget string) gin.HandlerFunc {
	return func(c *gin.Context) {
		target, err := url.Parse(proxyTarget)
		if err != nil {
			panic(err)
		}

		reverseProxy := httputil.NewSingleHostReverseProxy(target)

		reverseProxy.ServeHTTP(c.Writer, c.Request)
	}
}
