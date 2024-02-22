// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/liushuojia/go-sdk/mysql/daemon/gen/model"
)

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(&model.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewAsterisk(tableName)
	_user.ID = field.NewInt64(tableName, "id")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")
	_user.Username = field.NewString(tableName, "username")
	_user.Realname = field.NewString(tableName, "realname")
	_user.Email = field.NewString(tableName, "email")
	_user.Mobile = field.NewString(tableName, "mobile")
	_user.Status = field.NewInt64(tableName, "status")
	_user.Verify = field.NewString(tableName, "verify")
	_user.WeekVerify = field.NewString(tableName, "week_verify")
	_user.EntryDate = field.NewInt64(tableName, "entry_date")
	_user.OpRole = field.NewInt64(tableName, "op_role")
	_user.AdminRole = field.NewInt64(tableName, "admin_role")
	_user.EmailFlag = field.NewInt64(tableName, "email_flag")
	_user.MobileFlag = field.NewInt64(tableName, "mobile_flag")
	_user.DepartmentCode = field.NewString(tableName, "department_code")
	_user.DepartmentName = field.NewString(tableName, "department_name")
	_user.SelectStopTime = field.NewInt64(tableName, "select_stop_time")
	_user.BelongOp = field.NewInt64(tableName, "belong_op")
	_user.BelongOpName = field.NewString(tableName, "belong_op_name")
	_user.BelongDepartmentCode = field.NewString(tableName, "belong_department_code")
	_user.BelongDepartmentName = field.NewString(tableName, "belong_department_name")
	_user.BelongOpSec = field.NewInt64(tableName, "belong_op_sec")
	_user.BelongOpNameSec = field.NewString(tableName, "belong_op_name_sec")
	_user.BelongDepartmentCodeSec = field.NewString(tableName, "belong_department_code_sec")
	_user.BelongDepartmentNameSec = field.NewString(tableName, "belong_department_name_sec")
	_user.SourceName = field.NewString(tableName, "source_name")
	_user.SourceCode = field.NewString(tableName, "source_code")
	_user.RoleName = field.NewString(tableName, "role_name")
	_user.RoleCode = field.NewString(tableName, "role_code")
	_user.PlatName = field.NewString(tableName, "plat_name")
	_user.PlatCode = field.NewString(tableName, "plat_code")
	_user.Sex = field.NewInt64(tableName, "sex")
	_user.Qq = field.NewString(tableName, "qq")
	_user.WorkQq = field.NewString(tableName, "work_qq")
	_user.WorkMobile = field.NewString(tableName, "work_mobile")
	_user.WorkWeixin = field.NewString(tableName, "work_weixin")
	_user.Wangwang = field.NewString(tableName, "wangwang")
	_user.Weixin = field.NewString(tableName, "weixin")
	_user.UserType = field.NewString(tableName, "user_type")
	_user.UserTypeCode = field.NewString(tableName, "user_type_code")
	_user.LastFollowTime = field.NewInt64(tableName, "last_follow_time")
	_user.LastConsultationTime = field.NewInt64(tableName, "last_consultation_time")
	_user.LastIncallTime = field.NewInt64(tableName, "last_incall_time")
	_user.IncallDistributionFlag = field.NewInt64(tableName, "incall_distribution_flag")
	_user.IncallDistributionNum = field.NewInt64(tableName, "incall_distribution_num")
	_user.CreateName = field.NewString(tableName, "create_name")
	_user.OpChangeNum = field.NewInt64(tableName, "op_change_num")
	_user.BelongUserID = field.NewInt64(tableName, "belong_user_id")
	_user.WjtToken = field.NewString(tableName, "wjt_token")

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo

	ALL                     field.Asterisk
	ID                      field.Int64  // 自动编号
	CreatedAt               field.Time   // 创建时间
	UpdatedAt               field.Time   // 更新时间
	DeletedAt               field.Field  // 删除时间
	Username                field.String // 英文名
	Realname                field.String // 姓名
	Email                   field.String // 邮箱
	Mobile                  field.String // 手机
	Status                  field.Int64  // 状态
	Verify                  field.String // 密钥
	WeekVerify              field.String // 每周密钥
	EntryDate               field.Int64  // 入职时间
	OpRole                  field.Int64  // 后台用户
	AdminRole               field.Int64  // 管理员
	EmailFlag               field.Int64  // 邮件验证状态
	MobileFlag              field.Int64  // 手机验证状态
	DepartmentCode          field.String // 部门编码
	DepartmentName          field.String // 部门名称
	SelectStopTime          field.Int64  // 停用账号时间
	BelongOp                field.Int64  // 所属客服id
	BelongOpName            field.String // 所属客服名称
	BelongDepartmentCode    field.String // 客服所属部门编码
	BelongDepartmentName    field.String // 客服所属部门
	BelongOpSec             field.Int64  // 所属客服二id
	BelongOpNameSec         field.String // 所属客服二名称
	BelongDepartmentCodeSec field.String // 所属客服二部门编码
	BelongDepartmentNameSec field.String // 所属客服二部门
	SourceName              field.String // 来源
	SourceCode              field.String // 来源编码
	RoleName                field.String // 新角色名称
	RoleCode                field.String // 新角色编码
	PlatName                field.String // 角色菜单名称
	PlatCode                field.String // 角色菜单编码
	Sex                     field.Int64  // 性别
	Qq                      field.String // QQ号
	WorkQq                  field.String // 工作QQ
	WorkMobile              field.String // 工作手机
	WorkWeixin              field.String // 工作微信
	Wangwang                field.String // 阿里旺旺
	Weixin                  field.String // 微信号
	UserType                field.String // 客户类型
	UserTypeCode            field.String // 客户类型编码
	LastFollowTime          field.Int64  // 最后跟进时间
	LastConsultationTime    field.Int64  // 最后分配咨询时间
	LastIncallTime          field.Int64  // 最后咨询时间
	IncallDistributionFlag  field.Int64  // 咨询分配开关
	IncallDistributionNum   field.Int64  // 咨询分配总数
	CreateName              field.String // 创建人
	OpChangeNum             field.Int64  // 转交OP次数
	BelongUserID            field.Int64  // 所属联系人id，如果不为0，则被合并
	WjtToken                field.String // WJT加密密钥

	fieldMap map[string]field.Expr
}

func (u user) Table(newTableName string) *user {
	u.userDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user) updateTableName(table string) *user {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.Username = field.NewString(table, "username")
	u.Realname = field.NewString(table, "realname")
	u.Email = field.NewString(table, "email")
	u.Mobile = field.NewString(table, "mobile")
	u.Status = field.NewInt64(table, "status")
	u.Verify = field.NewString(table, "verify")
	u.WeekVerify = field.NewString(table, "week_verify")
	u.EntryDate = field.NewInt64(table, "entry_date")
	u.OpRole = field.NewInt64(table, "op_role")
	u.AdminRole = field.NewInt64(table, "admin_role")
	u.EmailFlag = field.NewInt64(table, "email_flag")
	u.MobileFlag = field.NewInt64(table, "mobile_flag")
	u.DepartmentCode = field.NewString(table, "department_code")
	u.DepartmentName = field.NewString(table, "department_name")
	u.SelectStopTime = field.NewInt64(table, "select_stop_time")
	u.BelongOp = field.NewInt64(table, "belong_op")
	u.BelongOpName = field.NewString(table, "belong_op_name")
	u.BelongDepartmentCode = field.NewString(table, "belong_department_code")
	u.BelongDepartmentName = field.NewString(table, "belong_department_name")
	u.BelongOpSec = field.NewInt64(table, "belong_op_sec")
	u.BelongOpNameSec = field.NewString(table, "belong_op_name_sec")
	u.BelongDepartmentCodeSec = field.NewString(table, "belong_department_code_sec")
	u.BelongDepartmentNameSec = field.NewString(table, "belong_department_name_sec")
	u.SourceName = field.NewString(table, "source_name")
	u.SourceCode = field.NewString(table, "source_code")
	u.RoleName = field.NewString(table, "role_name")
	u.RoleCode = field.NewString(table, "role_code")
	u.PlatName = field.NewString(table, "plat_name")
	u.PlatCode = field.NewString(table, "plat_code")
	u.Sex = field.NewInt64(table, "sex")
	u.Qq = field.NewString(table, "qq")
	u.WorkQq = field.NewString(table, "work_qq")
	u.WorkMobile = field.NewString(table, "work_mobile")
	u.WorkWeixin = field.NewString(table, "work_weixin")
	u.Wangwang = field.NewString(table, "wangwang")
	u.Weixin = field.NewString(table, "weixin")
	u.UserType = field.NewString(table, "user_type")
	u.UserTypeCode = field.NewString(table, "user_type_code")
	u.LastFollowTime = field.NewInt64(table, "last_follow_time")
	u.LastConsultationTime = field.NewInt64(table, "last_consultation_time")
	u.LastIncallTime = field.NewInt64(table, "last_incall_time")
	u.IncallDistributionFlag = field.NewInt64(table, "incall_distribution_flag")
	u.IncallDistributionNum = field.NewInt64(table, "incall_distribution_num")
	u.CreateName = field.NewString(table, "create_name")
	u.OpChangeNum = field.NewInt64(table, "op_change_num")
	u.BelongUserID = field.NewInt64(table, "belong_user_id")
	u.WjtToken = field.NewString(table, "wjt_token")

	u.fillFieldMap()

	return u
}

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 51)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["username"] = u.Username
	u.fieldMap["realname"] = u.Realname
	u.fieldMap["email"] = u.Email
	u.fieldMap["mobile"] = u.Mobile
	u.fieldMap["status"] = u.Status
	u.fieldMap["verify"] = u.Verify
	u.fieldMap["week_verify"] = u.WeekVerify
	u.fieldMap["entry_date"] = u.EntryDate
	u.fieldMap["op_role"] = u.OpRole
	u.fieldMap["admin_role"] = u.AdminRole
	u.fieldMap["email_flag"] = u.EmailFlag
	u.fieldMap["mobile_flag"] = u.MobileFlag
	u.fieldMap["department_code"] = u.DepartmentCode
	u.fieldMap["department_name"] = u.DepartmentName
	u.fieldMap["select_stop_time"] = u.SelectStopTime
	u.fieldMap["belong_op"] = u.BelongOp
	u.fieldMap["belong_op_name"] = u.BelongOpName
	u.fieldMap["belong_department_code"] = u.BelongDepartmentCode
	u.fieldMap["belong_department_name"] = u.BelongDepartmentName
	u.fieldMap["belong_op_sec"] = u.BelongOpSec
	u.fieldMap["belong_op_name_sec"] = u.BelongOpNameSec
	u.fieldMap["belong_department_code_sec"] = u.BelongDepartmentCodeSec
	u.fieldMap["belong_department_name_sec"] = u.BelongDepartmentNameSec
	u.fieldMap["source_name"] = u.SourceName
	u.fieldMap["source_code"] = u.SourceCode
	u.fieldMap["role_name"] = u.RoleName
	u.fieldMap["role_code"] = u.RoleCode
	u.fieldMap["plat_name"] = u.PlatName
	u.fieldMap["plat_code"] = u.PlatCode
	u.fieldMap["sex"] = u.Sex
	u.fieldMap["qq"] = u.Qq
	u.fieldMap["work_qq"] = u.WorkQq
	u.fieldMap["work_mobile"] = u.WorkMobile
	u.fieldMap["work_weixin"] = u.WorkWeixin
	u.fieldMap["wangwang"] = u.Wangwang
	u.fieldMap["weixin"] = u.Weixin
	u.fieldMap["user_type"] = u.UserType
	u.fieldMap["user_type_code"] = u.UserTypeCode
	u.fieldMap["last_follow_time"] = u.LastFollowTime
	u.fieldMap["last_consultation_time"] = u.LastConsultationTime
	u.fieldMap["last_incall_time"] = u.LastIncallTime
	u.fieldMap["incall_distribution_flag"] = u.IncallDistributionFlag
	u.fieldMap["incall_distribution_num"] = u.IncallDistributionNum
	u.fieldMap["create_name"] = u.CreateName
	u.fieldMap["op_change_num"] = u.OpChangeNum
	u.fieldMap["belong_user_id"] = u.BelongUserID
	u.fieldMap["wjt_token"] = u.WjtToken
}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user) replaceDB(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userDo struct{ gen.DO }

func (u userDo) Debug() *userDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) *userDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) ReadDB() *userDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDo) WriteDB() *userDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDo) Session(config *gorm.Session) *userDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDo) Clauses(conds ...clause.Expression) *userDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Returning(value interface{}, columns ...string) *userDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDo) Not(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) *userDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) *userDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() *userDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*model.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*model.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Take() (*model.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Last() (*model.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Find() ([]*model.User, error) {
	result, err := u.DO.Find()
	return result.([]*model.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User, err error) {
	buf := make([]*model.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(fields ...field.RelationField) *userDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDo) Preload(fields ...field.RelationField) *userDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDo) FirstOrInit() (*model.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FirstOrCreate() (*model.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*model.User, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDo) Delete(models ...*model.User) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}