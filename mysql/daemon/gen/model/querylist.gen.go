// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQueryList = "queryList"

// QueryList mapped from table <queryList>
type QueryList struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`        // 自动编号
	CreatedAt    time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                      // 创建时间
	UpdatedAt    time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                      // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                      // 删除时间
	UserID       int64          `gorm:"column:user_id;not null;comment:用户编号" json:"user_id"`                   // 用户编号
	QueryKey     string         `gorm:"column:query_key;not null;comment:关键key 跟业务前端挂钩" json:"query_key"`      // 关键key 跟业务前端挂钩
	QueryTitle   string         `gorm:"column:query_title;not null;comment:显示标题" json:"query_title"`           // 显示标题
	QueryColumns string         `gorm:"column:query_columns;comment:字段数组 为空时跟默认一致 json" json:"query_columns"`  // 字段数组 为空时跟默认一致 json
	QueryMap     string         `gorm:"column:query_map;comment:字段数组 查询 json" json:"query_map"`                // 字段数组 查询 json
	QueryMapMore string         `gorm:"column:query_map_more;comment:字段数组 更多查询 json" json:"query_map_more"`    // 字段数组 更多查询 json
	QueryKeys    string         `gorm:"column:query_keys;comment:初始化查询结果 json" json:"query_keys"`              // 初始化查询结果 json
	IndexOpen    int64          `gorm:"column:index_open;not null;comment:1 新窗打开详情  0 页面跳转" json:"index_open"` // 1 新窗打开详情  0 页面跳转
	Weight       int64          `gorm:"column:weight;not null;comment:权重排序 越大越靠前" json:"weight"`               // 权重排序 越大越靠前
	Status       int64          `gorm:"column:status;not null;comment:状态 1有效 0 停用" json:"status"`              // 状态 1有效 0 停用
}

func (m *QueryList) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName QueryList's table name
func (*QueryList) TableName() string {
	return TableNameQueryList
}
