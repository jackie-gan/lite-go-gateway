package middleware

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse("")
		if err != nil {
			panic(err)
		}

		reverseProxy := httputil.NewSingleHostReverseProxy(remote)

		reverseProxy.ServeHTTP(c.Writer, c.Request)
	}
}
