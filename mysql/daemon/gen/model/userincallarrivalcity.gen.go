// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserIncallArrivalCity = "userIncallArrivalCity"

// UserIncallArrivalCity mapped from table <userIncallArrivalCity>
type UserIncallArrivalCity struct {
	ID              int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`           // 自动编号
	CreatedAt       time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                         // 创建时间
	UpdatedAt       time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                         // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                         // 删除时间
	UID             int64          `gorm:"column:uid;not null;comment:用户id" json:"uid"`                              // 用户id
	ArrivalCityCode string         `gorm:"column:arrival_city_code;not null;comment:目的地编码" json:"arrival_city_code"` // 目的地编码
	ArrivalCity     string         `gorm:"column:arrival_city;not null;comment:目的地名称" json:"arrival_city"`           // 目的地名称
}

func (m *UserIncallArrivalCity) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName UserIncallArrivalCity's table name
func (*UserIncallArrivalCity) TableName() string {
	return TableNameUserIncallArrivalCity
}
