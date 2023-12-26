package GIN

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"time"
)

/*
	conditionList []gen.Condition

	ginQuery := GIN.New().Gin(c).
		InInt64(admin.ID, "id").
		InInt64(admin.Status, "status").
		LikeArray(admin.DepartmentCode, "department_code").
		LikeArray(admin.AreaCode, "area_code").
		LikeArray(admin.RoleCode, "role_code").
		RangeTime(admin.CreatedAt, "created_at").
		RangeTime(admin.UpdatedAt, "updated_at")

	if len(ginQuery.Where) > 0 {
		conditionList = append(conditionList, ginQuery.Where...)
	}
	if len(ginQuery.Or) > 0 {
		for _, v := range ginQuery.Or {
			tmpDo := admin.WithContext(ctx)
			if len(v) > 0 {
				for _, vq := range v {
					tmpDo = tmpDo.Or(vq)
				}
				conditionList = append(conditionList, tmpDo)
			}
		}
	}

	if valueList := extend.GinStringSlice(c, "order_by"); len(valueList) > 0 {
		for _, v := range valueList {
			if v == "" {
				continue
			}
			tt := strings.Split(v, " ")
			f, ok := admin.GetFieldByName(tt[0])
			if ok {
				if len(tt) == 2 && tt[1] == "desc" {
					orderExpr = append(orderExpr, f.Desc())
				} else {
					orderExpr = append(orderExpr, f)
				}
			}
		}
	} else {
		orderExpr = append(orderExpr, admin.UpdatedAt.Desc())
	}
*/

func New() *Query {
	return &Query{}
}

const (
	defaultPageField     = "page"
	defaultPageSizeField = "pageSize"
	defaultPageSize      = 30
)

type Query struct {
	c     *gin.Context
	Where []gen.Condition
	//Or    [][]gen.Condition

	Page     int
	PageSize int
}

func (q *Query) Gin(c *gin.Context) *Query {
	q.c = c
	return q.InitPage()
}
func (q *Query) InitPage() *Query {
	var errPage error
	var errPageSize error

	q.Page, errPage = Int(q.c, defaultPageField)
	q.PageSize, errPageSize = Int(q.c, defaultPageSizeField)
	if errPage != nil || q.Page <= 0 {
		q.Page = 1
	}
	if errPageSize != nil || q.PageSize <= 0 {
		q.PageSize = defaultPageSize
	}
	return q
}
func (q *Query) EqInt64(f field.Int64, n string) *Query {
	if value, err := Int64(q.c, n); err == nil {
		q.Where = append(q.Where, f.Eq(value))
	}
	return q
}
func (q *Query) Eq(f field.String, n string) *Query {
	if value := String(q.c, n); value != "" {
		q.Where = append(q.Where, f.Eq(value))
	}
	return q
}
func (q *Query) Like(f field.String, n string) *Query {
	if value := String(q.c, n+"_like"); value != "" {
		q.Where = append(q.Where, f.Like("%"+value+"%"))
	}
	return q.Eq(f, n)
}
func (q *Query) LikeArray(f field.String, n string) *Query {
	if value := String(q.c, n+"_right_like"); value != "" {
		q.Where = append(q.Where, f.Like(value+"%"))
	}
	//if value := StringSlice(q.c, n+"_right_like_in"); len(value) > 0 {
	//	var conditions []gen.Condition
	//	for _, v := range value {
	//		if v != "" {
	//			conditions = append(conditions, f.Like(v+"%"))
	//		}
	//	}
	//	if len(conditions) > 0 {
	//		q.Or = append(q.Or, conditions)
	//		//q.Where = append(q.Where)
	//	}
	//}
	return q.Eq(f, n)
}
func (q *Query) InInt64(f field.Int64, n string) *Query {
	if value := Int64Slice(q.c, n+"_in"); len(value) > 0 {
		q.Where = append(q.Where, f.In(value...))
	}
	return q.EqInt64(f, n)
}
func (q *Query) In(f field.String, n string) *Query {
	if value := StringSlice(q.c, n+"_in"); len(value) > 0 {
		q.Where = append(q.Where, f.In(value...))
	}
	return q.Eq(f, n)
}
func (q *Query) RangeInt64(f field.Int64, n string) *Query {
	if value, err := Int64(q.c, n+"_min"); err == nil {
		q.Where = append(q.Where, f.Gte(value))
	}
	if value, err := Int64(q.c, n+"_max"); err == nil {
		q.Where = append(q.Where, f.Lte(value))
	}
	return q.EqInt64(f, n)
}
func (q *Query) RangeString(f field.String, n string) *Query {
	if value := String(q.c, n+"_min"); value != "" {
		q.Where = append(q.Where, f.Gte(value))
	}
	if value := String(q.c, n+"_max"); value != "" {
		q.Where = append(q.Where, f.Lte(value))
	}
	return q.Eq(f, n)
}
func (q *Query) RangeTime(f field.Time, n string) *Query {
	if value, err := Int64(q.c, n+"_min"); err == nil && value > 0 {
		q.Where = append(q.Where, f.Gte(time.Unix(value, 0)))
	} else {
		if value := String(q.c, n+"_min"); value != "" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
			if err == nil {
				q.Where = append(q.Where, f.Gte(t))
			}
		}
	}
	if value, err := Int64(q.c, n+"_max"); err == nil && value > 0 {
		q.Where = append(q.Where, f.Lte(time.Unix(value, 0)))
	} else {
		if value := String(q.c, n+"_max"); value != "" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
			if err == nil {
				q.Where = append(q.Where, f.Lte(t))
			}
		}
	}
	return q
}

/*
Build
or的方式比较特殊需要实体化的query，暂放弃
*/
func (q *Query) Build() []gen.Condition {
	var conditionList []gen.Condition
	if len(q.Where) > 0 {
		conditionList = append(conditionList, q.Where...)
	}
	//if len(q.Or) > 0 {
	//	for _, v := range q.Or {
	//		tmpDo := gen.Cond()
	//		if len(v) > 0 {
	//			for _, vq := range v {
	//				tmpDo = tmpDo.Or(vq)
	//			}
	//			conditionList = append(conditionList, tmpDo)
	//		}
	//	}
	//}
	//fmt.Println(q.Where)
	//fmt.Println(q.Or)
	//if len(q.Or) > 0 {
	//	for _, v := range q.Or {
	//		tmpDo := admin.WithContext(ctx)
	//		if len(v) > 0 {
	//			for _, vq := range v {
	//				tmpDo = tmpDo.Or(vq)
	//			}
	//			conditionList = append(conditionList, tmpDo)
	//		}
	//	}
	//}
	return conditionList
}
