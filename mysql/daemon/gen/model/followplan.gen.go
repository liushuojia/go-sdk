// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFollowPlan = "followPlan"

// FollowPlan mapped from table <followPlan>
type FollowPlan struct {
	ID                   int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`                      // 自动编号
	CreatedAt            time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                                    // 创建时间
	UpdatedAt            time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                                    // 更新时间
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                    // 删除时间
	DataCode             string         `gorm:"column:data_code;not null;comment:业务所属编码" json:"data_code"`                           // 业务所属编码
	DataType             int64          `gorm:"column:data_type;not null;comment:业务所属类型id" json:"data_type"`                         // 业务所属类型id
	DataID               int64          `gorm:"column:data_id;not null;comment:业务所属自增id" json:"data_id"`                             // 业务所属自增id
	BelongOp             int64          `gorm:"column:belong_op;not null;comment:所属OP" json:"belong_op"`                             // 所属OP
	BelongOpName         string         `gorm:"column:belong_op_name;not null;comment:所属op姓名" json:"belong_op_name"`                 // 所属op姓名
	BelongDepartmentCode string         `gorm:"column:belong_department_code;not null;comment:所属部门编码" json:"belong_department_code"` // 所属部门编码
	BelongDepartmentName string         `gorm:"column:belong_department_name;not null;comment:所属部门名称" json:"belong_department_name"` // 所属部门名称
	CreateOp             int64          `gorm:"column:create_op;not null;comment:创建OP" json:"create_op"`                             // 创建OP
	CreateOpName         string         `gorm:"column:create_op_name;not null;comment:创建OP姓名" json:"create_op_name"`                 // 创建OP姓名
	FollowTitle          string         `gorm:"column:follow_title;not null;comment:跟进标题" json:"follow_title"`                       // 跟进标题
	FollowDesc           string         `gorm:"column:follow_desc;comment:跟进内容" json:"follow_desc"`                                  // 跟进内容
	Status               int64          `gorm:"column:status;not null;comment:状态 0 计划跟进中 1 跟进完成" json:"status"`                      // 状态 0 计划跟进中 1 跟进完成
	FollowResult         string         `gorm:"column:follow_result;comment:跟进结果描述" json:"follow_result"`                            // 跟进结果描述
	FollowPlanDay        int64          `gorm:"column:follow_plan_day;not null;comment:计划跟进日期时间戳" json:"follow_plan_day"`            // 计划跟进日期时间戳
	AutoFollowID         int64          `gorm:"column:auto_follow_id;not null;comment:系统计划跟进id 索引id" json:"auto_follow_id"`          // 系统计划跟进id 索引id
}

func (m *FollowPlan) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName FollowPlan's table name
func (*FollowPlan) TableName() string {
	return TableNameFollowPlan
}
