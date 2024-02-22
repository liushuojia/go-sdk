// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserLogin = "userLogin"

// UserLogin mapped from table <userLogin>
type UserLogin struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自动编号" json:"id"`   // 自动编号
	CreatedAt   time.Time      `gorm:"column:created_at;comment:创建时间" json:"created_at"`                 // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                 // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                 // 删除时间
	UID         int64          `gorm:"column:uid;not null;comment:用户id" json:"uid"`                      // 用户id
	LoginType   int64          `gorm:"column:login_type;not null;comment:登录类型 0微信" json:"login_type"`    // 登录类型 0微信
	WeixinID    int64          `gorm:"column:weixin_id;not null;comment:微信的账号id" json:"weixin_id"`       // 微信的账号id
	WeixinTitle string         `gorm:"column:weixin_title;not null;comment:微信的账号名称" json:"weixin_title"` // 微信的账号名称
	OpenID      string         `gorm:"column:open_id;not null;comment:微信openid" json:"open_id"`          // 微信openid
	Headimgurl  string         `gorm:"column:headimgurl;not null;comment:wx头像" json:"headimgurl"`        // wx头像
	Nickname    string         `gorm:"column:nickname;comment:wx名称" json:"nickname"`                     // wx名称
	Sex         int64          `gorm:"column:sex;not null;comment:性别" json:"sex"`                        // 性别
	Subscribe   int64          `gorm:"column:subscribe;not null;comment:是否订阅" json:"subscribe"`          // 是否订阅
}

func (m *UserLogin) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName UserLogin's table name
func (*UserLogin) TableName() string {
	return TableNameUserLogin
}
