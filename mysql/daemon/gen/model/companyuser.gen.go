// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCompanyUser = "companyUser"

// CompanyUser mapped from table <companyUser>
type CompanyUser struct {
	ID              int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`             // 自动编号
	CreatedAt       time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                           // 创建时间
	UpdatedAt       time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                           // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
	UID             int64          `gorm:"column:uid;not null;comment:用户编号" json:"uid"`                                // 用户编号
	Cid             int64          `gorm:"column:cid;not null;comment:公司编码" json:"cid"`                                // 公司编码
	CompanyUserCode string         `gorm:"column:company_user_code;not null;comment:联系人类型编码" json:"company_user_code"` // 联系人类型编码
	CompanyUserName string         `gorm:"column:company_user_name;not null;comment:联系人类型名称" json:"company_user_name"` // 联系人类型名称
	ExpiryTime      int64          `gorm:"column:expiry_time;not null;comment:有效期 0 长期有效" json:"expiry_time"`          // 有效期 0 长期有效
	Status          int64          `gorm:"column:status;not null;comment:状态 0 禁用 1正常" json:"status"`                   // 状态 0 禁用 1正常
}

func (m *CompanyUser) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName CompanyUser's table name
func (*CompanyUser) TableName() string {
	return TableNameCompanyUser
}
