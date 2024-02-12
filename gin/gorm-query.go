package GIN

import (
	"github.com/gin-gonic/gin"
	"time"
)

func NewGorm() *Gorm {
	return &Gorm{
		SearchMap: make(map[string]any),
	}
}

type Gorm struct {
	c         *gin.Context
	SearchMap map[string]any
	page      int
	pageSize  int
}

func (q *Gorm) Gin(c *gin.Context) *Gorm {
	q.c = c
	return q.InitPage().Eq("order_by")
}

func (q *Gorm) SetPage(page int) *Gorm {
	q.page = page
	return q
}
func (q *Gorm) SetPageSize(pageSize int) *Gorm {
	q.pageSize = pageSize
	return q
}

func (q *Gorm) GetSearchMap() map[string]any {
	return q.SearchMap
}
func (q *Gorm) GetPage() int {
	return q.page
}
func (q *Gorm) GetPageSize() int {
	return q.pageSize
}
func (q *Gorm) AddSearch(key string, value any) *Gorm {
	q.SearchMap[key] = value
	return q
}

func (q *Gorm) InitPage() *Gorm {
	var errPage error
	var errPageSize error

	q.page, errPage = Int(q.c, defaultPageField)
	q.pageSize, errPageSize = Int(q.c, defaultPageSizeField)
	if errPage != nil || q.page <= 0 {
		q.page = 1
	}
	if errPageSize != nil || q.pageSize <= 0 {
		q.pageSize = defaultPageSize
	}
	return q
}

func (q *Gorm) EqInt64(n string) *Gorm {
	if value, err := Int64(q.c, n); err == nil {
		q.SearchMap[n] = value
	}
	return q
}
func (q *Gorm) Eq(n string) *Gorm {
	if value := String(q.c, n); value != "" {
		q.SearchMap[n] = value
	}
	return q
}
func (q *Gorm) Like(n string) *Gorm {
	if value := String(q.c, n+"_like"); value != "" {
		q.SearchMap[n+"_like"] = "%" + value + "%"
	}
	return q.Eq(n)
}
func (q *Gorm) LikeArray(n string) *Gorm {
	if value := String(q.c, n+"_right_like"); value != "" {
		q.SearchMap[n+"_right_like"] = value + "%"
	}
	if value := StringSlice(q.c, n+"_right_like_in"); len(value) > 0 {
		q.SearchMap[n+"_right_like_in"] = value
	}
	return q.Eq(n)
}
func (q *Gorm) InInt64(n string) *Gorm {
	if value := Int64Slice(q.c, n+"_in"); len(value) > 0 {
		q.SearchMap[n+"_in"] = value
	}
	return q.EqInt64(n)
}
func (q *Gorm) In(n string) *Gorm {
	if value := StringSlice(q.c, n+"_in"); len(value) > 0 {
		q.SearchMap[n+"_in"] = value
	}
	return q.Eq(n)
}
func (q *Gorm) RangeInt64(n string) *Gorm {
	if value, err := Int64(q.c, n+"_min"); err == nil {
		q.SearchMap[n+"_min"] = value

	}
	if value, err := Int64(q.c, n+"_max"); err == nil {
		q.SearchMap[n+"_max"] = value
	}
	return q.EqInt64(n)
}
func (q *Gorm) RangeString(n string) *Gorm {
	if value := String(q.c, n+"_min"); value != "" {
		q.SearchMap[n+"_min"] = value
	}
	if value := String(q.c, n+"_max"); value != "" {
		q.SearchMap[n+"_max"] = value
	}
	return q.Eq(n)
}
func (q *Gorm) RangeTime(n string) *Gorm {
	if value, err := Int64(q.c, n+"_min"); err == nil && value > 0 {
		q.SearchMap[n+"_min"] = time.Unix(value, 0)
	} else {
		if value := String(q.c, n+"_min"); value != "" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
			if err == nil {
				q.SearchMap[n+"_min"] = t
			}
		}
	}
	if value, err := Int64(q.c, n+"_max"); err == nil && value > 0 {
		q.SearchMap[n+"_max"] = time.Unix(value, 0)
	} else {
		if value := String(q.c, n+"_max"); value != "" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
			if err == nil {
				q.SearchMap[n+"_max"] = t
			}
		}
	}
	return q
}
