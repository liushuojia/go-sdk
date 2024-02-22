// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserCheckPower = "userCheckPower"

// UserCheckPower mapped from table <userCheckPower>
type UserCheckPower struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`             // 自动编号
	CreatedAt  time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                           // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                           // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
	UID        int64          `gorm:"column:uid;not null;comment:用户id" json:"uid"`                                // 用户id
	UseCode    string         `gorm:"column:use_code;not null;comment:权限编码" json:"use_code"`                      // 权限编码
	UseName    string         `gorm:"column:use_name;not null;comment:权限名称" json:"use_name"`                      // 权限名称
	ExpiryTime int64          `gorm:"column:expiry_time;not null;comment:有效期 0 不限 其他 时间戳为有效期" json:"expiry_time"` // 有效期 0 不限 其他 时间戳为有效期
}

func (m *UserCheckPower) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName UserCheckPower's table name
func (*UserCheckPower) TableName() string {
	return TableNameUserCheckPower
}