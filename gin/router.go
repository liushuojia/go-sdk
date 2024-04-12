package GIN

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Router(debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	//router := gin.New()

	// 全局中间件
	//router.Use(gin.Recovery())

	// 日志中间件
	//router.Use(gin.Logger())
	//router.Use(loggerToFile())

	router := gin.Default()

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

func loggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		log.WithFields(log.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime.String(),
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
