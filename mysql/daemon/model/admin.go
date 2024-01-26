package model

import (
	"errors"
	"github.com/liushuojia/go-sdk/mysql"
	"time"

	"gorm.io/gorm"
)

const TableNameAdmin = "admin"

type Admin struct {
	db              *gorm.DB
	tableNameSuffix string

	ID                      int64          `json:"id" gorm:"column:id; primaryKey; type:bigint; comment:自动编号;"`
	CreatedAt               time.Time      `json:"created_at" gorm:"column:created_at; type:timestamp; autoCreateTime; not null; comment:创建时间;"`
	UpdatedAt               time.Time      `json:"updated_at" gorm:"column:updated_at; type:timestamp; autoUpdateTime; not null; comment:更新时间;"`
	DeletedAt               gorm.DeletedAt `json:"-" gorm:"column:deleted_at; type:timestamp; index; comment:删除时间;"`
	Username                string         `json:"username" gorm:"column:username; type:char(32); not null; comment:英文名;"`
	Realname                string         `json:"realname" gorm:"column:realname; type:char(32); not null; comment:姓名;"`
	Email                   string         `json:"email" gorm:"column:email; type:char(64); not null; comment:邮箱;"`
	Mobile                  string         `json:"mobile" gorm:"column:mobile; type:char(11); not null; index; comment:手机;"`
	Status                  int64          `json:"status" gorm:"column:status; type:tinyint; not null; index; comment:状态;"`
	Verify                  string         `json:"verify" gorm:"column:verify; type:char(32); not null; comment:密钥;"`
	WeekVerify              string         `json:"week_verify" gorm:"column:week_verify; type:char(32); not null; comment:每周密钥;"`
	EntryDate               int64          `json:"entry_date" gorm:"column:entry_date; type:int(11); not null; index; comment:入职时间;"`
	OpRole                  int64          `json:"op_role" gorm:"column:op_role; type:tinyint; not null; index; comment:后台用户;"`
	AdminRole               int64          `json:"admin_role" gorm:"column:admin_role; type:tinyint; not null; index; comment:管理员;"`
	EmailFlag               int64          `json:"email_flag" gorm:"column:email_flag; type:tinyint; not null; index; comment:邮件验证状态;"`
	MobileFlag              int64          `json:"mobile_flag" gorm:"column:mobile_flag; type:tinyint; not null; index; comment:手机验证状态;"`
	DepartmentCode          string         `json:"department_code" gorm:"column:department_code; type:char(32); not null; index; comment:部门编码;"`
	DepartmentName          string         `json:"department_name" gorm:"column:department_name; type:char(32); not null; comment:部门名称;"`
	SelectStopTime          int64          `json:"select_stop_time" gorm:"column:select_stop_time; type:int(11); not null; index; comment:停用账号时间;"`
	BelongOp                int64          `json:"belong_op" gorm:"column:belong_op; type:bigint; not null; index; comment:所属客服id;"`
	BelongOpName            string         `json:"belong_op_name" gorm:"column:belong_op_name; type:char(32); not null; comment:所属客服名称;"`
	BelongDepartmentCode    string         `json:"belong_department_code" gorm:"column:belong_department_code; type:char(32); not null; index; comment:客服所属部门编码;"`
	BelongDepartmentName    string         `json:"belong_department_name" gorm:"column:belong_department_name; type:char(32); not null; comment:客服所属部门;"`
	BelongOpSec             int64          `json:"belong_op_sec" gorm:"column:belong_op_sec; type:bigint; not null; index; comment:所属客服二id;"`
	BelongOpNameSec         string         `json:"belong_op_name_sec" gorm:"column:belong_op_name_sec; type:char(32); not null; comment:所属客服二名称;"`
	BelongDepartmentCodeSec string         `json:"belong_department_code_sec" gorm:"column:belong_department_code_sec; type:char(32); not null; index; comment:所属客服二部门编码;"`
	BelongDepartmentNameSec string         `json:"belong_department_name_sec" gorm:"column:belong_department_name_sec; type:char(32); not null; comment:所属客服二部门;"`
	SourceName              string         `json:"source_name" gorm:"column:source_name; type:char(32); not null; comment:来源;"`
	SourceCode              string         `json:"source_code" gorm:"column:source_code; type:char(32); not null; index; comment:来源编码;"`
	RoleName                string         `json:"role_name" gorm:"column:role_name; type:char(32); not null; comment:新角色名称;"`
	RoleCode                string         `json:"role_code" gorm:"column:role_code; type:char(32); not null; index; comment:新角色编码;"`
	PlatName                string         `json:"plat_name" gorm:"column:plat_name; type:char(32); not null; comment:角色菜单名称;"`
	PlatCode                string         `json:"plat_code" gorm:"column:plat_code; type:char(32); not null; index; comment:角色菜单编码;"`
	Sex                     int64          `json:"sex" gorm:"column:sex; type:tinyint; not null; comment:性别;"`
	Qq                      string         `json:"qq" gorm:"column:qq; type:char(32); not null; comment:QQ号;"`
	WorkQq                  string         `json:"work_qq" gorm:"column:work_qq; type:char(32); not null; comment:工作QQ;"`
	WorkMobile              string         `json:"work_mobile" gorm:"column:work_mobile; type:char(11); not null; comment:工作手机;"`
	WorkWeixin              string         `json:"work_weixin" gorm:"column:work_weixin; type:char(32); not null; comment:工作微信;"`
	Wangwang                string         `json:"wangwang" gorm:"column:wangwang; type:char(32); not null; comment:阿里旺旺;"`
	Weixin                  string         `json:"weixin" gorm:"column:weixin; type:char(32); not null; comment:微信号;"`
	LastFollowTime          int64          `json:"last_follow_time" gorm:"column:last_follow_time; type:int(11); not null; index; comment:最后跟进时间;"`
	LastConsultationTime    int64          `json:"last_consultation_time" gorm:"column:last_consultation_time; type:int(11); not null; index; comment:最后分配咨询时间;"`
	LastIncallTime          int64          `json:"last_incall_time" gorm:"column:last_incall_time; type:int(11); not null; index; comment:最后咨询时间;"`
	IncallDistributionFlag  int64          `json:"incall_distribution_flag" gorm:"column:incall_distribution_flag; type:tinyint; not null; index; comment:咨询分配开关;"`
	IncallDistributionNum   int64          `json:"incall_distribution_num" gorm:"column:incall_distribution_num; type:int(11); not null; comment:咨询分配总数;"`
	CreateName              string         `json:"create_name" gorm:"column:create_name; type:char(32); not null; comment:创建人;"`
	OpChangeNum             int64          `json:"op_change_num" gorm:"column:op_change_num; type:int(11); not null; comment:转交OP次数;"`
	BelongUserId            int64          `json:"belong_user_id" gorm:"column:belong_user_id; type:bigint; not null; index; comment:所属联系人id，如果不为0，则被合并;"`
}

func (m *Admin) TableName() string {
	return TableNameAdmin
}

func (m *Admin) QueryDB(searchMap map[string]any) *gorm.DB {
	db := mysqlConn.DefaultField().Query(m, searchMap)
	return db
}
func (m *Admin) SetDB(db *gorm.DB) *Admin {
	m.db = db
	return m
}
func (m *Admin) GetDB() *gorm.DB {
	return m.db.
		Session(
			&gorm.Session{
				QueryFields: true,
				PrepareStmt: true,
				NewDB:       true,
			},
		).
		Set("tableNameSuffix", m.tableNameSuffix).
		Scopes(mysqlConn.TableOfCode(m, m.tableNameSuffix)).
		Model(m).
		Table(m.TableName()).
		Select(mysqlConn.DefaultField().GetFieldList(m))
}
func (m *Admin) SetTableNameSuffix(tableNameCode string) *Admin {
	m.tableNameSuffix = tableNameCode
	return m
}
func (m *Admin) GetTableNameSuffix() string {
	return m.tableNameSuffix
}
func (m *Admin) Create() error {
	return m.GetDB().Create(m).Error
}
func (m *Admin) Select(id int64) error {
	if id <= 0 {
		return errors.New("ID 为空")
	}
	return m.GetDB().Take(m).Error
}
func (m *Admin) Change(updateMap map[string]interface{}) error {
	if m.ID <= 0 {
		return errors.New("ID 为空")
	}

	updateDataMap := mysqlConn.DefaultField().CleanUpdateMap(m, updateMap)
	if len(updateDataMap) <= 0 {
		return errors.New("无更新")
	}
	return m.GetDB().Updates(updateDataMap).Error
}
func (m *Admin) Update() error {
	if m.ID <= 0 {
		return errors.New("ID 为空")
	}
	return m.GetDB().Updates(m).Error
}
func (m *Admin) Delete() error {
	if m.ID <= 0 {
		return errors.New("ID 为空")
	}
	return m.GetDB().Delete(m).Error
}
func (m *Admin) Query(searchMap map[string]any, page int64, pageSize int64) (results []*Admin, err error) {
	db := m.QueryDB(searchMap)
	if page > 1 && pageSize > 0 {
		db = db.
			Limit(int(page)).
			Offset(int((page - 1) * pageSize))
	} else {
		if pageSize > 0 {
			db = db.Limit(int(pageSize))
		}
	}
	err = db.Find(&results).Error
	return
}
func (m *Admin) Size(searchMap map[string]any) (int64, error) {
	var num int64
	err := m.QueryDB(searchMap).
		Count(&num).
		Error
	return num, err
}
func (m *Admin) AfterFind(tx *gorm.DB) error {
	if tableNameCode, ok := tx.Get("tableNameCode"); ok {
		if value, ok := tableNameCode.(string); ok {
			m.SetTableNameSuffix(value)
		}
	}
	m.SetDB(tx)
	return nil
}
