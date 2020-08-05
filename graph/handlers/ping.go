package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler is simple keep-alive/ping handler
func PingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
