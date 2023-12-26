package GIN

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// 全局中间件
	//router.Use(gin.Recovery())

	// 日志中间件
	//router.Use(gin.Logger())
	//router.Use(loggerToFile())

	// 404
	router.NoRoute(func(c *gin.Context) {
		//返回404状态码
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "404, page not exists!",
			"data": map[string]string{
				"host":   c.Request.Host,
				"method": c.Request.Method,
				"proto":  c.Request.Proto,
				"path":   c.Request.URL.Path,
				"ip":     c.ClientIP(),
			},
		})
		return
	})

	// 服务器状态
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	return router
}
