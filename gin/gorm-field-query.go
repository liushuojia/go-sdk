package GIN

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

func NewGormQuery() *GormQuery {
	return &GormQuery{}
}

type GormQuery struct {
	c          *gin.Context
	expression []clause.Expression
	orderBy    []string
	page       int64
	pageSize   int64
}

func (q *GormQuery) Gin(c *gin.Context) *GormQuery {
	q.c = c
	return q.InitPage().InitOrderBy()
}

func (q *GormQuery) GetPage() int64 {
	return q.page
}
func (q *GormQuery) GetPageSize() int64 {
	return q.pageSize
}
func (q *GormQuery) GetOrderBy() []string {
	return q.orderBy
}
func (q *GormQuery) GetQuery() []clause.Expression {
	return q.expression
}

func (q *GormQuery) SetPage(page int64) *GormQuery {
	q.page = page
	return q
}
func (q *GormQuery) SetPageSize(pageSize int64) *GormQuery {
	q.pageSize = pageSize
	return q
}
func (q *GormQuery) InitPage() *GormQuery {
	var errPage error
	var errPageSize error

	q.page, errPage = Int64(q.c, defaultPageField)
	q.pageSize, errPageSize = Int64(q.c, defaultPageSizeField)
	if errPage != nil || q.page <= 0 {
		q.page = 1
	}
	if errPageSize != nil || q.pageSize <= 0 {
		q.pageSize = defaultPageSize
	}
	return q
}
func (q *GormQuery) InitOrderBy() *GormQuery {
	var orderExpr []string
	if valueList := StringSlice(q.c, defaultOrderByField); len(valueList) > 0 {
		for _, v := range valueList {
			if v == "" {
				continue
			}
			tt := strings.Split(v, " ")
			if len(tt) == 2 && tt[1] == "desc" {
				orderExpr = append(orderExpr, fmt.Sprintf("%s desc", tt[0]))
			} else {
				orderExpr = append(orderExpr, tt[0])
			}
		}
	} else {
		orderExpr = append(orderExpr, "updated_at desc")
	}
	return q
}
func (q *GormQuery) AddSearch(sqlString string, args ...interface{}) *GormQuery {
	q.expression = append(q.expression, gorm.Expr(sqlString, args...))
	return q
}
func (q *GormQuery) EqInt64(n string) *GormQuery {
	if value, err := Int64(q.c, n); err == nil {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" = ?", value),
		)
	}
	return q
}
func (q *GormQuery) Eq(n string) *GormQuery {
	if value := String(q.c, n); value != "" {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" = ?", value),
		)
	}
	return q
}
func (q *GormQuery) Like(n string) *GormQuery {
	if value := String(q.c, n+"_like"); value != "" {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" like ?", "%"+value+"%"),
		)
	}
	return q.Eq(n)
}
func (q *GormQuery) LikeArray(n string) *GormQuery {
	if value := String(q.c, n+"_right_like"); value != "" {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" like ?", value+"%"),
		)
	}
	if value := StringSlice(q.c, n+"_right_like_in"); len(value) > 0 {
		sqlString := "1=2 "
		var conditions []interface{}
		for _, v := range value {
			if v != "" {
				sqlString += " or " + n + " like ?"
				conditions = append(conditions, v+"%")
			}
		}
		if len(conditions) > 0 {
			q.expression = append(
				q.expression,
				gorm.Expr(sqlString, conditions...),
			)
		}
	}
	return q.Eq(n)
}
func (q *GormQuery) InInt64(n string) *GormQuery {
	if value := Int64Slice(q.c, n+"_in"); len(value) > 0 {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" in ?", value),
		)
	}
	return q.EqInt64(n)
}
func (q *GormQuery) In(n string) *GormQuery {
	if value := StringSlice(q.c, n+"_in"); len(value) > 0 {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" in ?", value),
		)
	}
	return q.Eq(n)
}
func (q *GormQuery) RangeInt64(n string) *GormQuery {
	if value, err := Int64(q.c, n+"_min"); err == nil {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" >= ?", value),
		)
	}
	if value, err := Int64(q.c, n+"_max"); err == nil {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" <= ?", value),
		)
	}
	return q.EqInt64(n)
}
func (q *GormQuery) RangeString(n string) *GormQuery {
	if value := String(q.c, n+"_min"); value != "" {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" >= ?", value),
		)
	}
	if value := String(q.c, n+"_max"); value != "" {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" <= ?", value),
		)
	}
	return q.Eq(n)
}
func (q *GormQuery) RangeTime(n string) *GormQuery {
	if value, err := Int64(q.c, n+"_min"); err == nil && value > 0 {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" >= ?", time.Unix(value, 0)),
		)
	} else {
		if value := String(q.c, n+"_min"); value != "" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
			if err == nil {
				q.expression = append(
					q.expression,
					gorm.Expr(n+" >= ?", t),
				)
			}
		}
	}
	if value, err := Int64(q.c, n+"_max"); err == nil && value > 0 {
		q.expression = append(
			q.expression,
			gorm.Expr(n+" <= ?", time.Unix(value, 0)),
		)
	} else {
		if value := String(q.c, n+"_max"); value != "" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
			if err == nil {
				q.expression = append(
					q.expression,
					gorm.Expr(n+" <= ?", t),
				)
			}
		}
	}
	return q
}
