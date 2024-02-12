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

	"github.com/liushuojia/go-sdk/example/model"
)

func newAdmin(db *gorm.DB, opts ...gen.DOOption) admin {
	_admin := admin{}

	_admin.adminDo.UseDB(db, opts...)
	_admin.adminDo.UseModel(&model.Admin{})

	tableName := _admin.adminDo.TableName()
	_admin.ALL = field.NewAsterisk(tableName)
	_admin.ID = field.NewInt64(tableName, "id")
	_admin.CreatedAt = field.NewTime(tableName, "created_at")
	_admin.UpdatedAt = field.NewTime(tableName, "updated_at")
	_admin.DeletedAt = field.NewField(tableName, "deleted_at")
	_admin.Username = field.NewString(tableName, "username")
	_admin.Realname = field.NewString(tableName, "realname")
	_admin.Email = field.NewString(tableName, "email")
	_admin.Mobile = field.NewString(tableName, "mobile")
	_admin.Status = field.NewInt64(tableName, "status")
	_admin.Verify = field.NewString(tableName, "verify")
	_admin.WeekVerify = field.NewString(tableName, "week_verify")
	_admin.EntryDate = field.NewInt64(tableName, "entry_date")
	_admin.OpRole = field.NewInt64(tableName, "op_role")
	_admin.AdminRole = field.NewInt64(tableName, "admin_role")
	_admin.EmailFlag = field.NewInt64(tableName, "email_flag")
	_admin.MobileFlag = field.NewInt64(tableName, "mobile_flag")
	_admin.DepartmentCode = field.NewString(tableName, "department_code")
	_admin.DepartmentName = field.NewString(tableName, "department_name")
	_admin.SelectStopTime = field.NewInt64(tableName, "select_stop_time")
	_admin.BelongOp = field.NewInt64(tableName, "belong_op")
	_admin.BelongOpName = field.NewString(tableName, "belong_op_name")
	_admin.BelongDepartmentCode = field.NewString(tableName, "belong_department_code")
	_admin.BelongDepartmentName = field.NewString(tableName, "belong_department_name")
	_admin.BelongOpSec = field.NewInt64(tableName, "belong_op_sec")
	_admin.BelongOpNameSec = field.NewString(tableName, "belong_op_name_sec")
	_admin.BelongDepartmentCodeSec = field.NewString(tableName, "belong_department_code_sec")
	_admin.BelongDepartmentNameSec = field.NewString(tableName, "belong_department_name_sec")
	_admin.SourceName = field.NewString(tableName, "source_name")
	_admin.SourceCode = field.NewString(tableName, "source_code")
	_admin.RoleName = field.NewString(tableName, "role_name")
	_admin.RoleCode = field.NewString(tableName, "role_code")
	_admin.PlatName = field.NewString(tableName, "plat_name")
	_admin.PlatCode = field.NewString(tableName, "plat_code")
	_admin.Sex = field.NewInt64(tableName, "sex")
	_admin.Qq = field.NewString(tableName, "qq")
	_admin.WorkQq = field.NewString(tableName, "work_qq")
	_admin.WorkMobile = field.NewString(tableName, "work_mobile")
	_admin.WorkWeixin = field.NewString(tableName, "work_weixin")
	_admin.Wangwang = field.NewString(tableName, "wangwang")
	_admin.Weixin = field.NewString(tableName, "weixin")
	_admin.LastFollowTime = field.NewInt64(tableName, "last_follow_time")
	_admin.LastConsultationTime = field.NewInt64(tableName, "last_consultation_time")
	_admin.LastIncallTime = field.NewInt64(tableName, "last_incall_time")
	_admin.IncallDistributionFlag = field.NewInt64(tableName, "incall_distribution_flag")
	_admin.IncallDistributionNum = field.NewInt64(tableName, "incall_distribution_num")
	_admin.CreateName = field.NewString(tableName, "create_name")
	_admin.OpChangeNum = field.NewInt64(tableName, "op_change_num")
	_admin.BelongUserID = field.NewInt64(tableName, "belong_user_id")

	_admin.fillFieldMap()

	return _admin
}

type admin struct {
	adminDo

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
	LastFollowTime          field.Int64  // 最后跟进时间
	LastConsultationTime    field.Int64  // 最后分配咨询时间
	LastIncallTime          field.Int64  // 最后咨询时间
	IncallDistributionFlag  field.Int64  // 咨询分配开关
	IncallDistributionNum   field.Int64  // 咨询分配总数
	CreateName              field.String // 创建人
	OpChangeNum             field.Int64  // 转交OP次数
	BelongUserID            field.Int64  // 所属联系人id，如果不为0，则被合并

	fieldMap map[string]field.Expr
}

func (a admin) Table(newTableName string) *admin {
	a.adminDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a admin) As(alias string) *admin {
	a.adminDo.DO = *(a.adminDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *admin) updateTableName(table string) *admin {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Username = field.NewString(table, "username")
	a.Realname = field.NewString(table, "realname")
	a.Email = field.NewString(table, "email")
	a.Mobile = field.NewString(table, "mobile")
	a.Status = field.NewInt64(table, "status")
	a.Verify = field.NewString(table, "verify")
	a.WeekVerify = field.NewString(table, "week_verify")
	a.EntryDate = field.NewInt64(table, "entry_date")
	a.OpRole = field.NewInt64(table, "op_role")
	a.AdminRole = field.NewInt64(table, "admin_role")
	a.EmailFlag = field.NewInt64(table, "email_flag")
	a.MobileFlag = field.NewInt64(table, "mobile_flag")
	a.DepartmentCode = field.NewString(table, "department_code")
	a.DepartmentName = field.NewString(table, "department_name")
	a.SelectStopTime = field.NewInt64(table, "select_stop_time")
	a.BelongOp = field.NewInt64(table, "belong_op")
	a.BelongOpName = field.NewString(table, "belong_op_name")
	a.BelongDepartmentCode = field.NewString(table, "belong_department_code")
	a.BelongDepartmentName = field.NewString(table, "belong_department_name")
	a.BelongOpSec = field.NewInt64(table, "belong_op_sec")
	a.BelongOpNameSec = field.NewString(table, "belong_op_name_sec")
	a.BelongDepartmentCodeSec = field.NewString(table, "belong_department_code_sec")
	a.BelongDepartmentNameSec = field.NewString(table, "belong_department_name_sec")
	a.SourceName = field.NewString(table, "source_name")
	a.SourceCode = field.NewString(table, "source_code")
	a.RoleName = field.NewString(table, "role_name")
	a.RoleCode = field.NewString(table, "role_code")
	a.PlatName = field.NewString(table, "plat_name")
	a.PlatCode = field.NewString(table, "plat_code")
	a.Sex = field.NewInt64(table, "sex")
	a.Qq = field.NewString(table, "qq")
	a.WorkQq = field.NewString(table, "work_qq")
	a.WorkMobile = field.NewString(table, "work_mobile")
	a.WorkWeixin = field.NewString(table, "work_weixin")
	a.Wangwang = field.NewString(table, "wangwang")
	a.Weixin = field.NewString(table, "weixin")
	a.LastFollowTime = field.NewInt64(table, "last_follow_time")
	a.LastConsultationTime = field.NewInt64(table, "last_consultation_time")
	a.LastIncallTime = field.NewInt64(table, "last_incall_time")
	a.IncallDistributionFlag = field.NewInt64(table, "incall_distribution_flag")
	a.IncallDistributionNum = field.NewInt64(table, "incall_distribution_num")
	a.CreateName = field.NewString(table, "create_name")
	a.OpChangeNum = field.NewInt64(table, "op_change_num")
	a.BelongUserID = field.NewInt64(table, "belong_user_id")

	a.fillFieldMap()

	return a
}

func (a *admin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *admin) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 48)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["username"] = a.Username
	a.fieldMap["realname"] = a.Realname
	a.fieldMap["email"] = a.Email
	a.fieldMap["mobile"] = a.Mobile
	a.fieldMap["status"] = a.Status
	a.fieldMap["verify"] = a.Verify
	a.fieldMap["week_verify"] = a.WeekVerify
	a.fieldMap["entry_date"] = a.EntryDate
	a.fieldMap["op_role"] = a.OpRole
	a.fieldMap["admin_role"] = a.AdminRole
	a.fieldMap["email_flag"] = a.EmailFlag
	a.fieldMap["mobile_flag"] = a.MobileFlag
	a.fieldMap["department_code"] = a.DepartmentCode
	a.fieldMap["department_name"] = a.DepartmentName
	a.fieldMap["select_stop_time"] = a.SelectStopTime
	a.fieldMap["belong_op"] = a.BelongOp
	a.fieldMap["belong_op_name"] = a.BelongOpName
	a.fieldMap["belong_department_code"] = a.BelongDepartmentCode
	a.fieldMap["belong_department_name"] = a.BelongDepartmentName
	a.fieldMap["belong_op_sec"] = a.BelongOpSec
	a.fieldMap["belong_op_name_sec"] = a.BelongOpNameSec
	a.fieldMap["belong_department_code_sec"] = a.BelongDepartmentCodeSec
	a.fieldMap["belong_department_name_sec"] = a.BelongDepartmentNameSec
	a.fieldMap["source_name"] = a.SourceName
	a.fieldMap["source_code"] = a.SourceCode
	a.fieldMap["role_name"] = a.RoleName
	a.fieldMap["role_code"] = a.RoleCode
	a.fieldMap["plat_name"] = a.PlatName
	a.fieldMap["plat_code"] = a.PlatCode
	a.fieldMap["sex"] = a.Sex
	a.fieldMap["qq"] = a.Qq
	a.fieldMap["work_qq"] = a.WorkQq
	a.fieldMap["work_mobile"] = a.WorkMobile
	a.fieldMap["work_weixin"] = a.WorkWeixin
	a.fieldMap["wangwang"] = a.Wangwang
	a.fieldMap["weixin"] = a.Weixin
	a.fieldMap["last_follow_time"] = a.LastFollowTime
	a.fieldMap["last_consultation_time"] = a.LastConsultationTime
	a.fieldMap["last_incall_time"] = a.LastIncallTime
	a.fieldMap["incall_distribution_flag"] = a.IncallDistributionFlag
	a.fieldMap["incall_distribution_num"] = a.IncallDistributionNum
	a.fieldMap["create_name"] = a.CreateName
	a.fieldMap["op_change_num"] = a.OpChangeNum
	a.fieldMap["belong_user_id"] = a.BelongUserID
}

func (a admin) clone(db *gorm.DB) admin {
	a.adminDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a admin) replaceDB(db *gorm.DB) admin {
	a.adminDo.ReplaceDB(db)
	return a
}

type adminDo struct{ gen.DO }

func (a adminDo) Debug() *adminDo {
	return a.withDO(a.DO.Debug())
}

func (a adminDo) WithContext(ctx context.Context) *adminDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminDo) ReadDB() *adminDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminDo) WriteDB() *adminDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminDo) Session(config *gorm.Session) *adminDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminDo) Clauses(conds ...clause.Expression) *adminDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminDo) Returning(value interface{}, columns ...string) *adminDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminDo) Not(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminDo) Or(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminDo) Select(conds ...field.Expr) *adminDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminDo) Where(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *adminDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a adminDo) Order(conds ...field.Expr) *adminDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminDo) Distinct(cols ...field.Expr) *adminDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminDo) Omit(cols ...field.Expr) *adminDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminDo) Join(table schema.Tabler, on ...field.Expr) *adminDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminDo) LeftJoin(table schema.Tabler, on ...field.Expr) *adminDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminDo) RightJoin(table schema.Tabler, on ...field.Expr) *adminDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminDo) Group(cols ...field.Expr) *adminDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminDo) Having(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminDo) Limit(limit int) *adminDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminDo) Offset(offset int) *adminDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *adminDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminDo) Unscoped() *adminDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminDo) Create(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminDo) CreateInBatches(values []*model.Admin, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminDo) Save(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminDo) First() (*model.Admin, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Take() (*model.Admin, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Last() (*model.Admin, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Find() ([]*model.Admin, error) {
	result, err := a.DO.Find()
	return result.([]*model.Admin), err
}

func (a adminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Admin, err error) {
	buf := make([]*model.Admin, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminDo) FindInBatches(result *[]*model.Admin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminDo) Attrs(attrs ...field.AssignExpr) *adminDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminDo) Assign(attrs ...field.AssignExpr) *adminDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminDo) Joins(fields ...field.RelationField) *adminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminDo) Preload(fields ...field.RelationField) *adminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminDo) FirstOrInit() (*model.Admin, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FirstOrCreate() (*model.Admin, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FindByPage(offset int, limit int) (result []*model.Admin, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminDo) Delete(models ...*model.Admin) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminDo) withDO(do gen.Dao) *adminDo {
	a.DO = *do.(*gen.DO)
	return a
}
