// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserRolePower = "userRolePower"

// UserRolePower mapped from table <userRolePower>
type UserRolePower struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`             // 自动编号
	CreatedAt  time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                           // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                           // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
	UID        int64          `gorm:"column:uid;not null;comment:用户id" json:"uid"`                                // 用户id
	RoleCode   string         `gorm:"column:role_code;not null;comment:角色编码" json:"role_code"`                    // 角色编码
	RoleName   string         `gorm:"column:role_name;not null;comment:角色名称" json:"role_name"`                    // 角色名称
	ExpiryTime int64          `gorm:"column:expiry_time;not null;comment:有效期 0 不限 其他 时间戳为有效期" json:"expiry_time"` // 有效期 0 不限 其他 时间戳为有效期
}

func (m *UserRolePower) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName UserRolePower's table name
func (*UserRolePower) TableName() string {
	return TableNameUserRolePower
}