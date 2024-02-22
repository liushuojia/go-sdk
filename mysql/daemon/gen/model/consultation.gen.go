// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameConsultation = "consultation"

// Consultation mapped from table <consultation>
type Consultation struct {
	ID                      int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`                            // 自动编号
	CreatedAt               time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                                          // 创建时间
	UpdatedAt               time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                                          // 更新时间
	DeletedAt               gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                          // 删除时间
	UID                     int64          `gorm:"column:uid;not null;comment:用户编码" json:"uid"`                                               // 用户编码
	UIDName                 string         `gorm:"column:uid_name;not null;comment:用户名字" json:"uid_name"`                                     // 用户名字
	UIDSex                  int64          `gorm:"column:uid_sex;not null;comment:用户性别" json:"uid_sex"`                                       // 用户性别
	Mobile                  string         `gorm:"column:mobile;not null;comment:手机号码" json:"mobile"`                                         // 手机号码
	Email                   string         `gorm:"column:email;not null;comment:邮箱" json:"email"`                                             // 邮箱
	Weixin                  string         `gorm:"column:weixin;not null;comment:微信号" json:"weixin"`                                          // 微信号
	Aliwangwang             string         `gorm:"column:aliwangwang;not null;comment:阿里旺旺" json:"aliwangwang"`                               // 阿里旺旺
	MafengwoUID             string         `gorm:"column:mafengwo_uid;not null;comment:马蜂窝UID" json:"mafengwo_uid"`                           // 马蜂窝UID
	MafengwoIm              string         `gorm:"column:mafengwo_im;not null;comment:马蜂窝IM" json:"mafengwo_im"`                              // 马蜂窝IM
	CtripAccount            string         `gorm:"column:ctrip_account;not null;comment:携程子账号" json:"ctrip_account"`                          // 携程子账号
	CtripGrab               string         `gorm:"column:ctrip_grab;not null;comment:携程抢单号" json:"ctrip_grab"`                                // 携程抢单号
	QunaerPhone             string         `gorm:"column:qunaer_phone;not null;comment:去哪儿电话" json:"qunaer_phone"`                            // 去哪儿电话
	DepartureCityCode       string         `gorm:"column:departure_city_code;not null;comment:出发城市编码" json:"departure_city_code"`             // 出发城市编码
	DepartureCity           string         `gorm:"column:departure_city;not null;comment:出发城市名称" json:"departure_city"`                       // 出发城市名称
	ArrivalCityCode         string         `gorm:"column:arrival_city_code;not null;comment:目的地编码" json:"arrival_city_code"`                  // 目的地编码
	ArrivalCity             string         `gorm:"column:arrival_city;not null;comment:目的地名称" json:"arrival_city"`                            // 目的地名称
	AdultNum                int64          `gorm:"column:adult_num;not null;comment:成人数" json:"adult_num"`                                    // 成人数
	ChildrenNum             int64          `gorm:"column:children_num;not null;comment:儿童数" json:"children_num"`                              // 儿童数
	TravelDate              int64          `gorm:"column:travel_date;not null;comment:出行日期" json:"travel_date"`                               // 出行日期
	TravelDateEnd           int64          `gorm:"column:travel_date_end;not null;comment:出行日期截止" json:"travel_date_end"`                     // 出行日期截止
	TravelDayNum            int64          `gorm:"column:travel_day_num;not null;comment:出行天数" json:"travel_day_num"`                         // 出行天数
	ReturnDate              int64          `gorm:"column:return_date;not null;comment:返程日期" json:"return_date"`                               // 返程日期
	ReturnDateEnd           int64          `gorm:"column:return_date_end;not null;comment:返程日期截止" json:"return_date_end"`                     // 返程日期截止
	Budget                  string         `gorm:"column:budget;not null;comment:预算" json:"budget"`                                           // 预算
	BudgetCode              string         `gorm:"column:budget_code;not null;comment:预算编码" json:"budget_code"`                               // 预算编码
	DemandNumber            string         `gorm:"column:demand_number;not null;comment:需求单号" json:"demand_number"`                           // 需求单号
	TransferNumber          string         `gorm:"column:transfer_number;not null;comment:转接号" json:"transfer_number"`                        // 转接号
	SourceCode              string         `gorm:"column:source_code;not null;comment:系统来源编码" json:"source_code"`                             // 系统来源编码
	SourceName              string         `gorm:"column:source_name;not null;comment:系统来源名称" json:"source_name"`                             // 系统来源名称
	SourceAbbreviation      string         `gorm:"column:source_abbreviation;not null;comment:系统来源简称" json:"source_abbreviation"`             // 系统来源简称
	MethodCode              string         `gorm:"column:method_code;not null;comment:咨询方式编码" json:"method_code"`                             // 咨询方式编码
	MethodName              string         `gorm:"column:method_name;not null;comment:咨询方式名称" json:"method_name"`                             // 咨询方式名称
	FollowStatusCode        string         `gorm:"column:follow_status_code;not null;comment:咨询跟进标识编码" json:"follow_status_code"`             // 咨询跟进标识编码
	FollowStatusName        string         `gorm:"column:follow_status_name;not null;comment:咨询跟进标识名称" json:"follow_status_name"`             // 咨询跟进标识名称
	ProductTypeCode         string         `gorm:"column:product_type_code;not null;comment:产品类型编码" json:"product_type_code"`                 // 产品类型编码
	ProductType             string         `gorm:"column:product_type;not null;comment:产品类型" json:"product_type"`                             // 产品类型
	ProductTypeAbbreviation string         `gorm:"column:product_type_abbreviation;not null;comment:产品类型简称" json:"product_type_abbreviation"` // 产品类型简称
	CreateOp                int64          `gorm:"column:create_op;not null;comment:创建opID" json:"create_op"`                                 // 创建opID
	CreateOpName            string         `gorm:"column:create_op_name;not null;comment:创建op姓名" json:"create_op_name"`                       // 创建op姓名
	BelongOp                int64          `gorm:"column:belong_op;not null;comment:所属op" json:"belong_op"`                                   // 所属op
	BelongOpName            string         `gorm:"column:belong_op_name;not null;comment:所属op姓名" json:"belong_op_name"`                       // 所属op姓名
	BelongDepartmentCode    string         `gorm:"column:belong_department_code;not null;comment:所属部门编码" json:"belong_department_code"`       // 所属部门编码
	BelongDepartmentName    string         `gorm:"column:belong_department_name;not null;comment:所属部门名称" json:"belong_department_name"`       // 所属部门名称
	FollowPlanFlag          int64          `gorm:"column:follow_plan_flag;not null;comment:是否远期计划" json:"follow_plan_flag"`                   // 是否远期计划
	FollowPlanDay           int64          `gorm:"column:follow_plan_day;not null;comment:远期计划日期" json:"follow_plan_day"`                     // 远期计划日期
	Status                  int64          `gorm:"column:status;not null;comment:状态 0 跟进中 1 历史咨询 2 远期计划" json:"status"`                       // 状态 0 跟进中 1 历史咨询 2 远期计划
	StatusOrder             int64          `gorm:"column:status_order;not null;comment:成单状态 0 未成单 1 已成单" json:"status_order"`                 // 成单状态 0 未成单 1 已成单
	BecomeOrderAt           int64          `gorm:"column:become_order_at;not null;comment:成单时间" json:"become_order_at"`                       // 成单时间
	BecomeOrderSecond       int64          `gorm:"column:become_order_second;not null;comment:成单时长 单位秒" json:"become_order_second"`           // 成单时长 单位秒
	StatusClose             int64          `gorm:"column:status_close;not null;comment:关闭状态	0 未关闭	1 已关闭	2 关闭中" json:"status_close"`           // 关闭状态	0 未关闭	1 已关闭	2 关闭中
	CloseReason             string         `gorm:"column:close_reason;not null;comment:关闭原因" json:"close_reason"`                             // 关闭原因
	CloseReasonCode         string         `gorm:"column:close_reason_code;not null;comment:关闭原因编码" json:"close_reason_code"`                 // 关闭原因编码
	LastFollowTime          int64          `gorm:"column:last_follow_time;not null;comment:最后处理时间" json:"last_follow_time"`                   // 最后处理时间
	SystemAutoFlag          int64          `gorm:"column:system_auto_flag;not null;comment:系统自动分配" json:"system_auto_flag"`                   // 系统自动分配
	SystemAutoTime          int64          `gorm:"column:system_auto_time;not null;comment:咨询分配op时间" json:"system_auto_time"`                 // 咨询分配op时间
}

func (m *Consultation) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName Consultation's table name
func (*Consultation) TableName() string {
	return TableNameConsultation
}