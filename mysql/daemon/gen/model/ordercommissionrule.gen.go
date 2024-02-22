// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOrderCommissionRule = "orderCommissionRule"

// OrderCommissionRule mapped from table <orderCommissionRule>
type OrderCommissionRule struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`          // 自动编号
	CreatedAt     time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                        // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                        // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                        // 删除时间
	Title         string         `gorm:"column:title;not null;comment:提成规则标题" json:"title"`                       // 提成规则标题
	OpGradient    int64          `gorm:"column:op_gradient;not null;comment:销售一 提成梯度" json:"op_gradient"`         // 销售一 提成梯度
	OpSecGradient int64          `gorm:"column:op_sec_gradient;not null;comment:销售二 提成梯度" json:"op_sec_gradient"` // 销售二 提成梯度
	Status        int64          `gorm:"column:status;not null;comment:0 暂停	1 有效" json:"status"`                  // 0 暂停	1 有效
}

func (m *OrderCommissionRule) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName OrderCommissionRule's table name
func (*OrderCommissionRule) TableName() string {
	return TableNameOrderCommissionRule
}
