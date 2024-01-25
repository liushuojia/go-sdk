package query

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func TableOfCode(u m, code string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := u.TableName()
		if code != "" {
			tableName += "_" + code
		}
		if !DefaultField().GetAutoMigrate(tableName) {
			db.Table(tableName).AutoMigrate(u)
		}
		return db.Table(tableName)
	}
}
func TableOfDB(u m, dbName string, code string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := u.TableName()
		if dbName != "" {
			tableName = dbName + "." + tableName
		}
		if code != "" {
			tableName += "_" + code
		}
		return db.Table(tableName)
	}
}
func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
