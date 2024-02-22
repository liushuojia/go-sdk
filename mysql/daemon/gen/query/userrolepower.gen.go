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

func newUserRolePower(db *gorm.DB, opts ...gen.DOOption) userRolePower {
	_userRolePower := userRolePower{}

	_userRolePower.userRolePowerDo.UseDB(db, opts...)
	_userRolePower.userRolePowerDo.UseModel(&model.UserRolePower{})

	tableName := _userRolePower.userRolePowerDo.TableName()
	_userRolePower.ALL = field.NewAsterisk(tableName)
	_userRolePower.ID = field.NewInt64(tableName, "id")
	_userRolePower.CreatedAt = field.NewTime(tableName, "created_at")
	_userRolePower.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userRolePower.DeletedAt = field.NewField(tableName, "deleted_at")
	_userRolePower.UID = field.NewInt64(tableName, "uid")
	_userRolePower.RoleCode = field.NewString(tableName, "role_code")
	_userRolePower.RoleName = field.NewString(tableName, "role_name")
	_userRolePower.ExpiryTime = field.NewInt64(tableName, "expiry_time")

	_userRolePower.fillFieldMap()

	return _userRolePower
}

type userRolePower struct {
	userRolePowerDo

	ALL        field.Asterisk
	ID         field.Int64  // 自动编号
	CreatedAt  field.Time   // 创建时间
	UpdatedAt  field.Time   // 更新时间
	DeletedAt  field.Field  // 删除时间
	UID        field.Int64  // 用户id
	RoleCode   field.String // 角色编码
	RoleName   field.String // 角色名称
	ExpiryTime field.Int64  // 有效期 0 不限 其他 时间戳为有效期

	fieldMap map[string]field.Expr
}

func (u userRolePower) Table(newTableName string) *userRolePower {
	u.userRolePowerDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userRolePower) As(alias string) *userRolePower {
	u.userRolePowerDo.DO = *(u.userRolePowerDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userRolePower) updateTableName(table string) *userRolePower {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.UID = field.NewInt64(table, "uid")
	u.RoleCode = field.NewString(table, "role_code")
	u.RoleName = field.NewString(table, "role_name")
	u.ExpiryTime = field.NewInt64(table, "expiry_time")

	u.fillFieldMap()

	return u
}

func (u *userRolePower) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userRolePower) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["uid"] = u.UID
	u.fieldMap["role_code"] = u.RoleCode
	u.fieldMap["role_name"] = u.RoleName
	u.fieldMap["expiry_time"] = u.ExpiryTime
}

func (u userRolePower) clone(db *gorm.DB) userRolePower {
	u.userRolePowerDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userRolePower) replaceDB(db *gorm.DB) userRolePower {
	u.userRolePowerDo.ReplaceDB(db)
	return u
}

type userRolePowerDo struct{ gen.DO }

func (u userRolePowerDo) Debug() *userRolePowerDo {
	return u.withDO(u.DO.Debug())
}

func (u userRolePowerDo) WithContext(ctx context.Context) *userRolePowerDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userRolePowerDo) ReadDB() *userRolePowerDo {
	return u.Clauses(dbresolver.Read)
}

func (u userRolePowerDo) WriteDB() *userRolePowerDo {
	return u.Clauses(dbresolver.Write)
}

func (u userRolePowerDo) Session(config *gorm.Session) *userRolePowerDo {
	return u.withDO(u.DO.Session(config))
}

func (u userRolePowerDo) Clauses(conds ...clause.Expression) *userRolePowerDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userRolePowerDo) Returning(value interface{}, columns ...string) *userRolePowerDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userRolePowerDo) Not(conds ...gen.Condition) *userRolePowerDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userRolePowerDo) Or(conds ...gen.Condition) *userRolePowerDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userRolePowerDo) Select(conds ...field.Expr) *userRolePowerDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userRolePowerDo) Where(conds ...gen.Condition) *userRolePowerDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userRolePowerDo) Order(conds ...field.Expr) *userRolePowerDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userRolePowerDo) Distinct(cols ...field.Expr) *userRolePowerDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userRolePowerDo) Omit(cols ...field.Expr) *userRolePowerDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userRolePowerDo) Join(table schema.Tabler, on ...field.Expr) *userRolePowerDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userRolePowerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userRolePowerDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userRolePowerDo) RightJoin(table schema.Tabler, on ...field.Expr) *userRolePowerDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userRolePowerDo) Group(cols ...field.Expr) *userRolePowerDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userRolePowerDo) Having(conds ...gen.Condition) *userRolePowerDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userRolePowerDo) Limit(limit int) *userRolePowerDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userRolePowerDo) Offset(offset int) *userRolePowerDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userRolePowerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userRolePowerDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userRolePowerDo) Unscoped() *userRolePowerDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userRolePowerDo) Create(values ...*model.UserRolePower) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userRolePowerDo) CreateInBatches(values []*model.UserRolePower, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userRolePowerDo) Save(values ...*model.UserRolePower) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userRolePowerDo) First() (*model.UserRolePower, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRolePower), nil
	}
}

func (u userRolePowerDo) Take() (*model.UserRolePower, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRolePower), nil
	}
}

func (u userRolePowerDo) Last() (*model.UserRolePower, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRolePower), nil
	}
}

func (u userRolePowerDo) Find() ([]*model.UserRolePower, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserRolePower), err
}

func (u userRolePowerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserRolePower, err error) {
	buf := make([]*model.UserRolePower, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userRolePowerDo) FindInBatches(result *[]*model.UserRolePower, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userRolePowerDo) Attrs(attrs ...field.AssignExpr) *userRolePowerDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userRolePowerDo) Assign(attrs ...field.AssignExpr) *userRolePowerDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userRolePowerDo) Joins(fields ...field.RelationField) *userRolePowerDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userRolePowerDo) Preload(fields ...field.RelationField) *userRolePowerDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userRolePowerDo) FirstOrInit() (*model.UserRolePower, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRolePower), nil
	}
}

func (u userRolePowerDo) FirstOrCreate() (*model.UserRolePower, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRolePower), nil
	}
}

func (u userRolePowerDo) FindByPage(offset int, limit int) (result []*model.UserRolePower, count int64, err error) {
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

func (u userRolePowerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userRolePowerDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userRolePowerDo) Delete(models ...*model.UserRolePower) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userRolePowerDo) withDO(do gen.Dao) *userRolePowerDo {
	u.DO = *do.(*gen.DO)
	return u
}
