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

func newUserExt(db *gorm.DB, opts ...gen.DOOption) userExt {
	_userExt := userExt{}

	_userExt.userExtDo.UseDB(db, opts...)
	_userExt.userExtDo.UseModel(&model.UserExt{})

	tableName := _userExt.userExtDo.TableName()
	_userExt.ALL = field.NewAsterisk(tableName)
	_userExt.ID = field.NewInt64(tableName, "id")
	_userExt.CreatedAt = field.NewTime(tableName, "created_at")
	_userExt.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userExt.DeletedAt = field.NewField(tableName, "deleted_at")
	_userExt.UID = field.NewInt64(tableName, "uid")
	_userExt.Phone = field.NewString(tableName, "phone")
	_userExt.Headimg = field.NewString(tableName, "headimg")
	_userExt.QunaerPhone = field.NewString(tableName, "qunaer_phone")
	_userExt.MafengwoUID = field.NewString(tableName, "mafengwo_uid")
	_userExt.MafengwoIm = field.NewString(tableName, "mafengwo_im")
	_userExt.CtripAccount = field.NewString(tableName, "ctrip_account")
	_userExt.Fax = field.NewString(tableName, "fax")
	_userExt.Birthday = field.NewInt64(tableName, "birthday")
	_userExt.BirthdayLunar = field.NewInt64(tableName, "birthday_lunar")
	_userExt.Address = field.NewString(tableName, "address")
	_userExt.AreaName = field.NewString(tableName, "area_name")
	_userExt.AreaCode = field.NewString(tableName, "area_code")
	_userExt.EmergencyContact1 = field.NewString(tableName, "emergency_contact1")
	_userExt.EmergencyPhone1 = field.NewString(tableName, "emergency_phone1")
	_userExt.EmergencyRelationship1 = field.NewString(tableName, "emergency_relationship1")
	_userExt.EmergencyContact2 = field.NewString(tableName, "emergency_contact2")
	_userExt.EmergencyPhone2 = field.NewString(tableName, "emergency_phone2")
	_userExt.EmergencyRelationship2 = field.NewString(tableName, "emergency_relationship2")
	_userExt.PublicKey = field.NewString(tableName, "public_key")
	_userExt.PrivateKey = field.NewString(tableName, "private_key")

	_userExt.fillFieldMap()

	return _userExt
}

type userExt struct {
	userExtDo

	ALL                    field.Asterisk
	ID                     field.Int64  // 自动编号
	CreatedAt              field.Time   // 创建时间
	UpdatedAt              field.Time   // 更新时间
	DeletedAt              field.Field  // 删除时间
	UID                    field.Int64  // 用户id
	Phone                  field.String // 固话
	Headimg                field.String // 头像
	QunaerPhone            field.String // 去哪儿电话
	MafengwoUID            field.String // 蚂蜂窝Uid
	MafengwoIm             field.String // 蚂蜂窝IM
	CtripAccount           field.String // 携程子账号
	Fax                    field.String // 传真
	Birthday               field.Int64  // 生日
	BirthdayLunar          field.Int64  // 农历生日
	Address                field.String // 地址
	AreaName               field.String // 区域名称
	AreaCode               field.String // 区域编码
	EmergencyContact1      field.String // 紧急联系人1
	EmergencyPhone1        field.String // 联系人1电话
	EmergencyRelationship1 field.String // 联系人1关系
	EmergencyContact2      field.String // 紧急联系人2
	EmergencyPhone2        field.String // 联系人2电话
	EmergencyRelationship2 field.String // 联系人2关系
	PublicKey              field.String // 公钥
	PrivateKey             field.String // 密钥

	fieldMap map[string]field.Expr
}

func (u userExt) Table(newTableName string) *userExt {
	u.userExtDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userExt) As(alias string) *userExt {
	u.userExtDo.DO = *(u.userExtDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userExt) updateTableName(table string) *userExt {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.UID = field.NewInt64(table, "uid")
	u.Phone = field.NewString(table, "phone")
	u.Headimg = field.NewString(table, "headimg")
	u.QunaerPhone = field.NewString(table, "qunaer_phone")
	u.MafengwoUID = field.NewString(table, "mafengwo_uid")
	u.MafengwoIm = field.NewString(table, "mafengwo_im")
	u.CtripAccount = field.NewString(table, "ctrip_account")
	u.Fax = field.NewString(table, "fax")
	u.Birthday = field.NewInt64(table, "birthday")
	u.BirthdayLunar = field.NewInt64(table, "birthday_lunar")
	u.Address = field.NewString(table, "address")
	u.AreaName = field.NewString(table, "area_name")
	u.AreaCode = field.NewString(table, "area_code")
	u.EmergencyContact1 = field.NewString(table, "emergency_contact1")
	u.EmergencyPhone1 = field.NewString(table, "emergency_phone1")
	u.EmergencyRelationship1 = field.NewString(table, "emergency_relationship1")
	u.EmergencyContact2 = field.NewString(table, "emergency_contact2")
	u.EmergencyPhone2 = field.NewString(table, "emergency_phone2")
	u.EmergencyRelationship2 = field.NewString(table, "emergency_relationship2")
	u.PublicKey = field.NewString(table, "public_key")
	u.PrivateKey = field.NewString(table, "private_key")

	u.fillFieldMap()

	return u
}

func (u *userExt) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userExt) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 25)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["uid"] = u.UID
	u.fieldMap["phone"] = u.Phone
	u.fieldMap["headimg"] = u.Headimg
	u.fieldMap["qunaer_phone"] = u.QunaerPhone
	u.fieldMap["mafengwo_uid"] = u.MafengwoUID
	u.fieldMap["mafengwo_im"] = u.MafengwoIm
	u.fieldMap["ctrip_account"] = u.CtripAccount
	u.fieldMap["fax"] = u.Fax
	u.fieldMap["birthday"] = u.Birthday
	u.fieldMap["birthday_lunar"] = u.BirthdayLunar
	u.fieldMap["address"] = u.Address
	u.fieldMap["area_name"] = u.AreaName
	u.fieldMap["area_code"] = u.AreaCode
	u.fieldMap["emergency_contact1"] = u.EmergencyContact1
	u.fieldMap["emergency_phone1"] = u.EmergencyPhone1
	u.fieldMap["emergency_relationship1"] = u.EmergencyRelationship1
	u.fieldMap["emergency_contact2"] = u.EmergencyContact2
	u.fieldMap["emergency_phone2"] = u.EmergencyPhone2
	u.fieldMap["emergency_relationship2"] = u.EmergencyRelationship2
	u.fieldMap["public_key"] = u.PublicKey
	u.fieldMap["private_key"] = u.PrivateKey
}

func (u userExt) clone(db *gorm.DB) userExt {
	u.userExtDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userExt) replaceDB(db *gorm.DB) userExt {
	u.userExtDo.ReplaceDB(db)
	return u
}

type userExtDo struct{ gen.DO }

func (u userExtDo) Debug() *userExtDo {
	return u.withDO(u.DO.Debug())
}

func (u userExtDo) WithContext(ctx context.Context) *userExtDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userExtDo) ReadDB() *userExtDo {
	return u.Clauses(dbresolver.Read)
}

func (u userExtDo) WriteDB() *userExtDo {
	return u.Clauses(dbresolver.Write)
}

func (u userExtDo) Session(config *gorm.Session) *userExtDo {
	return u.withDO(u.DO.Session(config))
}

func (u userExtDo) Clauses(conds ...clause.Expression) *userExtDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userExtDo) Returning(value interface{}, columns ...string) *userExtDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userExtDo) Not(conds ...gen.Condition) *userExtDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userExtDo) Or(conds ...gen.Condition) *userExtDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userExtDo) Select(conds ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userExtDo) Where(conds ...gen.Condition) *userExtDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userExtDo) Order(conds ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userExtDo) Distinct(cols ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userExtDo) Omit(cols ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userExtDo) Join(table schema.Tabler, on ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userExtDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userExtDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userExtDo) RightJoin(table schema.Tabler, on ...field.Expr) *userExtDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userExtDo) Group(cols ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userExtDo) Having(conds ...gen.Condition) *userExtDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userExtDo) Limit(limit int) *userExtDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userExtDo) Offset(offset int) *userExtDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userExtDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userExtDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userExtDo) Unscoped() *userExtDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userExtDo) Create(values ...*model.UserExt) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userExtDo) CreateInBatches(values []*model.UserExt, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userExtDo) Save(values ...*model.UserExt) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userExtDo) First() (*model.UserExt, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) Take() (*model.UserExt, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) Last() (*model.UserExt, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) Find() ([]*model.UserExt, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserExt), err
}

func (u userExtDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserExt, err error) {
	buf := make([]*model.UserExt, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userExtDo) FindInBatches(result *[]*model.UserExt, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userExtDo) Attrs(attrs ...field.AssignExpr) *userExtDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userExtDo) Assign(attrs ...field.AssignExpr) *userExtDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userExtDo) Joins(fields ...field.RelationField) *userExtDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userExtDo) Preload(fields ...field.RelationField) *userExtDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userExtDo) FirstOrInit() (*model.UserExt, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) FirstOrCreate() (*model.UserExt, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) FindByPage(offset int, limit int) (result []*model.UserExt, count int64, err error) {
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

func (u userExtDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userExtDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userExtDo) Delete(models ...*model.UserExt) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userExtDo) withDO(do gen.Dao) *userExtDo {
	u.DO = *do.(*gen.DO)
	return u
}
