package routes

// route 路由的配置及初始化

import (
	"02-GO_Web_CLI/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET(
		"/", func(c *gin.Context) {
			c.String(http.StatusOK, "ok!")
		},
	)
	return r
}
