// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameProductDefault = "productDefault"

// ProductDefault mapped from table <productDefault>
type ProductDefault struct {
	ID                   int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`                        // 自动编号
	CreatedAt            time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                                      // 创建时间
	UpdatedAt            time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                                      // 更新时间
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                      // 删除时间
	ProductName          string         `gorm:"column:product_name;not null;comment:产品名称" json:"product_name"`                         // 产品名称
	ProductDesc          string         `gorm:"column:product_desc;not null;comment:产品富描述" json:"product_desc"`                        // 产品富描述
	ProductManagerDesc   string         `gorm:"column:product_manager_desc;comment:产品经理推荐" json:"product_manager_desc"`                // 产品经理推荐
	ProductType          string         `gorm:"column:product_type;not null;comment:产品类型" json:"product_type"`                         // 产品类型
	ProductTypeCode      string         `gorm:"column:product_type_code;not null;comment:产品类型编码" json:"product_type_code"`             // 产品类型编码
	SaleLabel            string         `gorm:"column:sale_label;not null;comment:热卖标签" json:"sale_label"`                             // 热卖标签
	SaleLabelCode        string         `gorm:"column:sale_label_code;not null;comment:热卖标签编码" json:"sale_label_code"`                 // 热卖标签编码
	DepartureCity        string         `gorm:"column:departure_city;not null;comment:出发城市名称" json:"departure_city"`                   // 出发城市名称
	DepartureCityCode    string         `gorm:"column:departure_city_code;not null;comment:出发城市编码" json:"departure_city_code"`         // 出发城市编码
	ArrivalCity          string         `gorm:"column:arrival_city;not null;comment:目的地名称" json:"arrival_city"`                        // 目的地名称
	ArrivalCityCode      string         `gorm:"column:arrival_city_code;not null;comment:目的地编码" json:"arrival_city_code"`              // 目的地编码
	TravelDays           int64          `gorm:"column:travel_days;not null;comment:几天" json:"travel_days"`                             // 几天
	TravelNights         int64          `gorm:"column:travel_nights;not null;comment:几晚" json:"travel_nights"`                         // 几晚
	DepartureTraffic     string         `gorm:"column:departure_traffic;not null;comment:去程交通" json:"departure_traffic"`               // 去程交通
	DepartureTrafficCode string         `gorm:"column:departure_traffic_code;not null;comment:去程交通编码" json:"departure_traffic_code"`   // 去程交通编码
	DepartureTrafficDesc string         `gorm:"column:departure_traffic_desc;not null;comment:去程交通描述" json:"departure_traffic_desc"`   // 去程交通描述
	ReturnTraffic        string         `gorm:"column:return_traffic;not null;comment:返程交通" json:"return_traffic"`                     // 返程交通
	ReturnTrafficCode    string         `gorm:"column:return_traffic_code;not null;comment:返程交通编码" json:"return_traffic_code"`         // 返程交通编码
	ReturnTrafficDesc    string         `gorm:"column:return_traffic_desc;not null;comment:返程交通描述" json:"return_traffic_desc"`         // 返程交通描述
	PriceDesc            string         `gorm:"column:price_desc;not null;comment:价格描述" json:"price_desc"`                             // 价格描述
	PriceUnit            string         `gorm:"column:price_unit;not null;comment:价格单位" json:"price_unit"`                             // 价格单位
	PriceUnitCode        string         `gorm:"column:price_unit_code;not null;comment:价格单位编码" json:"price_unit_code"`                 // 价格单位编码
	BelongOp             int64          `gorm:"column:belong_op;not null;comment:所属客服id" json:"belong_op"`                             // 所属客服id
	BelongOpName         string         `gorm:"column:belong_op_name;not null;comment:所属客服名称" json:"belong_op_name"`                   // 所属客服名称
	BelongDepartmentCode string         `gorm:"column:belong_department_code;not null;comment:客服所属部门编码" json:"belong_department_code"` // 客服所属部门编码
	BelongDepartmentName string         `gorm:"column:belong_department_name;not null;comment:客服所属部门" json:"belong_department_name"`   // 客服所属部门
	Status               int64          `gorm:"column:status;not null;comment:状态 1 正常 0 停用" json:"status"`                             // 状态 1 正常 0 停用
	OpChangeNum          int64          `gorm:"column:op_change_num;not null;comment:转OP次数" json:"op_change_num"`                      // 转OP次数
}

func (m *ProductDefault) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName ProductDefault's table name
func (*ProductDefault) TableName() string {
	return TableNameProductDefault
}