// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOrderPayment = "orderPayment"

// OrderPayment mapped from table <orderPayment>
type OrderPayment struct {
	ID                           int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`                                      // 自动编号
	CreatedAt                    time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                                                    // 创建时间
	UpdatedAt                    time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                                                    // 更新时间
	DeletedAt                    gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                                    // 删除时间
	Oid                          int64          `gorm:"column:oid;not null;comment:订单号" json:"oid"`                                                          // 订单号
	CustomerCurrencyAbbreviation string         `gorm:"column:customer_currency_abbreviation;not null;comment:客户币种" json:"customer_currency_abbreviation"`   // 客户币种
	CustomerReceived             int64          `gorm:"column:customer_received;not null;comment:客户已收款" json:"customer_received"`                            // 客户已收款
	CustomerWaitReceived         int64          `gorm:"column:customer_wait_received;not null;comment:客户待收款" json:"customer_wait_received"`                  // 客户待收款
	CustomerCommission           int64          `gorm:"column:customer_commission;not null;comment:客户收款手续费" json:"customer_commission"`                      // 客户收款手续费
	CustomerRefund               int64          `gorm:"column:customer_refund;not null;comment:客户已退款" json:"customer_refund"`                                // 客户已退款
	CustomerWaitRefund           int64          `gorm:"column:customer_wait_refund;not null;comment:客户待退款" json:"customer_wait_refund"`                      // 客户待退款
	ProviderCurrencyAbbreviation string         `gorm:"column:provider_currency_abbreviation;not null;comment:供应商币种" json:"provider_currency_abbreviation"`  // 供应商币种
	ProviderSpend                int64          `gorm:"column:provider_spend;not null;comment:供应商已付款" json:"provider_spend"`                                 // 供应商已付款
	ProviderWaitSpend            int64          `gorm:"column:provider_wait_spend;not null;comment:供应商待付款" json:"provider_wait_spend"`                       // 供应商待付款
	ProviderReceived             int64          `gorm:"column:provider_received;not null;comment:供应商已退款" json:"provider_received"`                           // 供应商已退款
	ProviderWaitReceived         int64          `gorm:"column:provider_wait_received;not null;comment:供应商待退款" json:"provider_wait_received"`                 // 供应商待退款
	ProfitCurrencyAbbreviation   string         `gorm:"column:profit_currency_abbreviation;not null;comment:毛利币种" json:"profit_currency_abbreviation"`       // 毛利币种
	Profit                       int64          `gorm:"column:profit;not null;comment:当前盈亏" json:"profit"`                                                   // 当前盈亏
	ProfitEstimate               int64          `gorm:"column:profit_estimate;not null;comment:预估毛利" json:"profit_estimate"`                                 // 预估毛利
	OtherFlag                    int64          `gorm:"column:other_flag;not null;comment:业外开关 1 有 0 无" json:"other_flag"`                                   // 业外开关 1 有 0 无
	OtherCustomerReceived        int64          `gorm:"column:other_customer_received;not null;comment:营业外 客户已收款" json:"other_customer_received"`            // 营业外 客户已收款
	OtherCustomerWaitReceived    int64          `gorm:"column:other_customer_wait_received;not null;comment:营业外 客户待收款" json:"other_customer_wait_received"`  // 营业外 客户待收款
	OtherCustomerCommission      int64          `gorm:"column:other_customer_commission;not null;comment:营业外 客户收款手续费" json:"other_customer_commission"`      // 营业外 客户收款手续费
	OtherCustomerRefund          int64          `gorm:"column:other_customer_refund;not null;comment:营业外 客户已退款" json:"other_customer_refund"`                // 营业外 客户已退款
	OtherCustomerWaitRefund      int64          `gorm:"column:other_customer_wait_refund;not null;comment:营业外 客户待退款" json:"other_customer_wait_refund"`      // 营业外 客户待退款
	OtherProviderSpend           int64          `gorm:"column:other_provider_spend;not null;comment:营业外 供应商已付款" json:"other_provider_spend"`                 // 营业外 供应商已付款
	OtherProviderWaitSpend       int64          `gorm:"column:other_provider_wait_spend;not null;comment:营业外 供应商待付款" json:"other_provider_wait_spend"`       // 营业外 供应商待付款
	OtherProviderReceived        int64          `gorm:"column:other_provider_received;not null;comment:营业外 供应商已退款" json:"other_provider_received"`           // 营业外 供应商已退款
	OtherProviderWaitReceived    int64          `gorm:"column:other_provider_wait_received;not null;comment:营业外 供应商待退款" json:"other_provider_wait_received"` // 营业外 供应商待退款
}

func (m *OrderPayment) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName OrderPayment's table name
func (*OrderPayment) TableName() string {
	return TableNameOrderPayment
}
