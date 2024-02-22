// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameBankPaymentLink = "bankPaymentLink"

// BankPaymentLink mapped from table <bankPaymentLink>
type BankPaymentLink struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`       // 自动编号
	CreatedAt     time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                     // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                     // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                     // 删除时间
	BankPaymentID int64          `gorm:"column:bank_payment_id;not null;comment:手续费id" json:"bank_payment_id"` // 手续费id
	BankID        int64          `gorm:"column:bank_id;not null;comment:银行id" json:"bank_id"`                  // 银行id
}

func (m *BankPaymentLink) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName BankPaymentLink's table name
func (*BankPaymentLink) TableName() string {
	return TableNameBankPaymentLink
}
