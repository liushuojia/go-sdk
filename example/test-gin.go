package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"math"
	"net/http"
	"sdk/example/model"
	"sdk/example/query"
	GIN "sdk/gin"
	mysqlConn "sdk/mysql"
)

func main() {

	conn, err := mysqlConn.GormDB("root", "liushuojia", "127.0.0.1", 3306, "my_test")
	if err != nil {
		log.Fatalln(err)
	}
	query.SetDefault(conn)

	router := gin.New()

	// 全局中间件
	router.Use(gin.Recovery())

	router.Use(Auth())

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
	})

	// 服务器状态
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	//u := router.Group("user",Auth())
	u := router.Group("user")
	{
		u.GET("", func(c *gin.Context) {
			user := query.User

			q := GIN.New().Gin(c).
				Eq(user.Name, "name").
				EqInt64(user.ID, "id").
				LikeArray(user.Addr, "addr")

			userList, totalSize, err := user.Where(q.Build()...).FindByPage((q.Page-1)*q.PageSize, q.PageSize)

			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"page": gin.H{
					"page":      q.Page,
					"pageSize":  q.PageSize,
					"totalSize": totalSize,
					"totalPage": int64(math.Ceil(float64(totalSize) / float64(q.PageSize))),
				},
				"data": userList,
			})
		})

		u.POST("", func(c *gin.Context) {
			var userCreate model.User
			if _, err := GIN.Body(c, &userCreate); err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}
			user := query.User

			if err := user.Create(&userCreate); err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			c.JSON(http.StatusOK, userCreate.ID)
		})

		u.GET(":id", func(c *gin.Context) {
			id, err := GIN.ParamInt64(c, "id")
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}

			userSelect, err := query.User.Where(query.User.ID.Eq(id)).First()
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(http.StatusNotFound, err)
					return
				}
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			c.JSON(http.StatusOK, userSelect)
		})

		u.PUT(":id", func(c *gin.Context) {
			id, err := GIN.ParamInt64(c, "id")
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}

			var updateMap map[string]interface{}
			if _, err := GIN.Body(c, &updateMap); err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}

			if v, ok := updateMap["name"]; ok {
				value, ok := v.(string)
				if !ok {
					c.JSON(http.StatusBadRequest, "姓名非字符串")
					return
				}
				if value == "" {
					c.JSON(http.StatusBadRequest, "姓名内容不为空")
					return
				}
			}

			userSelect, err := query.User.Where(query.User.ID.Eq(id)).First()
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(http.StatusNotFound, err)
					return
				}
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			if err := mysqlConn.InitUpdateMap(userSelect, updateMap); err != nil || len(updateMap) <= 0 {
				c.JSON(http.StatusInternalServerError, "无更新内容")
				return
			}

			r, err := query.User.Where(query.User.ID.Eq(id)).Updates(updateMap)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			c.JSON(http.StatusOK, r)
		})

		u.DELETE(":id", func(c *gin.Context) {
			id, err := GIN.ParamInt64(c, "id")
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}

			if _, err := query.User.Where(query.User.ID.Eq(id)).First(); err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(http.StatusNotFound, err)
					return
				}
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			r, err := query.User.Where(query.User.ID.Eq(id)).Delete()
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			c.JSON(http.StatusOK, r)
		})
	}

	router.Run(":8080")
}

// Auth 后台中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		accessToken := c.GetHeader("token")
		//check

		if err != nil {
			c.JSON(http.StatusBadRequest, "need token")
			c.Abort()
			return
		}

		c.Set("token", accessToken)
		c.Next()
	}
}
