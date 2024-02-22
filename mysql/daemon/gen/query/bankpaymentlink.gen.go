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

func newBankPaymentLink(db *gorm.DB, opts ...gen.DOOption) bankPaymentLink {
	_bankPaymentLink := bankPaymentLink{}

	_bankPaymentLink.bankPaymentLinkDo.UseDB(db, opts...)
	_bankPaymentLink.bankPaymentLinkDo.UseModel(&model.BankPaymentLink{})

	tableName := _bankPaymentLink.bankPaymentLinkDo.TableName()
	_bankPaymentLink.ALL = field.NewAsterisk(tableName)
	_bankPaymentLink.ID = field.NewInt64(tableName, "id")
	_bankPaymentLink.CreatedAt = field.NewTime(tableName, "created_at")
	_bankPaymentLink.UpdatedAt = field.NewTime(tableName, "updated_at")
	_bankPaymentLink.DeletedAt = field.NewField(tableName, "deleted_at")
	_bankPaymentLink.BankPaymentID = field.NewInt64(tableName, "bank_payment_id")
	_bankPaymentLink.BankID = field.NewInt64(tableName, "bank_id")

	_bankPaymentLink.fillFieldMap()

	return _bankPaymentLink
}

type bankPaymentLink struct {
	bankPaymentLinkDo

	ALL           field.Asterisk
	ID            field.Int64 // 自动编号
	CreatedAt     field.Time  // 创建时间
	UpdatedAt     field.Time  // 更新时间
	DeletedAt     field.Field // 删除时间
	BankPaymentID field.Int64 // 手续费id
	BankID        field.Int64 // 银行id

	fieldMap map[string]field.Expr
}

func (b bankPaymentLink) Table(newTableName string) *bankPaymentLink {
	b.bankPaymentLinkDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bankPaymentLink) As(alias string) *bankPaymentLink {
	b.bankPaymentLinkDo.DO = *(b.bankPaymentLinkDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bankPaymentLink) updateTableName(table string) *bankPaymentLink {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewField(table, "deleted_at")
	b.BankPaymentID = field.NewInt64(table, "bank_payment_id")
	b.BankID = field.NewInt64(table, "bank_id")

	b.fillFieldMap()

	return b
}

func (b *bankPaymentLink) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bankPaymentLink) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 6)
	b.fieldMap["id"] = b.ID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
	b.fieldMap["bank_payment_id"] = b.BankPaymentID
	b.fieldMap["bank_id"] = b.BankID
}

func (b bankPaymentLink) clone(db *gorm.DB) bankPaymentLink {
	b.bankPaymentLinkDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bankPaymentLink) replaceDB(db *gorm.DB) bankPaymentLink {
	b.bankPaymentLinkDo.ReplaceDB(db)
	return b
}

type bankPaymentLinkDo struct{ gen.DO }

func (b bankPaymentLinkDo) Debug() *bankPaymentLinkDo {
	return b.withDO(b.DO.Debug())
}

func (b bankPaymentLinkDo) WithContext(ctx context.Context) *bankPaymentLinkDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bankPaymentLinkDo) ReadDB() *bankPaymentLinkDo {
	return b.Clauses(dbresolver.Read)
}

func (b bankPaymentLinkDo) WriteDB() *bankPaymentLinkDo {
	return b.Clauses(dbresolver.Write)
}

func (b bankPaymentLinkDo) Session(config *gorm.Session) *bankPaymentLinkDo {
	return b.withDO(b.DO.Session(config))
}

func (b bankPaymentLinkDo) Clauses(conds ...clause.Expression) *bankPaymentLinkDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bankPaymentLinkDo) Returning(value interface{}, columns ...string) *bankPaymentLinkDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bankPaymentLinkDo) Not(conds ...gen.Condition) *bankPaymentLinkDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bankPaymentLinkDo) Or(conds ...gen.Condition) *bankPaymentLinkDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bankPaymentLinkDo) Select(conds ...field.Expr) *bankPaymentLinkDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bankPaymentLinkDo) Where(conds ...gen.Condition) *bankPaymentLinkDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bankPaymentLinkDo) Order(conds ...field.Expr) *bankPaymentLinkDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bankPaymentLinkDo) Distinct(cols ...field.Expr) *bankPaymentLinkDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bankPaymentLinkDo) Omit(cols ...field.Expr) *bankPaymentLinkDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bankPaymentLinkDo) Join(table schema.Tabler, on ...field.Expr) *bankPaymentLinkDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bankPaymentLinkDo) LeftJoin(table schema.Tabler, on ...field.Expr) *bankPaymentLinkDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bankPaymentLinkDo) RightJoin(table schema.Tabler, on ...field.Expr) *bankPaymentLinkDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bankPaymentLinkDo) Group(cols ...field.Expr) *bankPaymentLinkDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bankPaymentLinkDo) Having(conds ...gen.Condition) *bankPaymentLinkDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bankPaymentLinkDo) Limit(limit int) *bankPaymentLinkDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bankPaymentLinkDo) Offset(offset int) *bankPaymentLinkDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bankPaymentLinkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *bankPaymentLinkDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bankPaymentLinkDo) Unscoped() *bankPaymentLinkDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bankPaymentLinkDo) Create(values ...*model.BankPaymentLink) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bankPaymentLinkDo) CreateInBatches(values []*model.BankPaymentLink, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bankPaymentLinkDo) Save(values ...*model.BankPaymentLink) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bankPaymentLinkDo) First() (*model.BankPaymentLink, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BankPaymentLink), nil
	}
}

func (b bankPaymentLinkDo) Take() (*model.BankPaymentLink, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BankPaymentLink), nil
	}
}

func (b bankPaymentLinkDo) Last() (*model.BankPaymentLink, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BankPaymentLink), nil
	}
}

func (b bankPaymentLinkDo) Find() ([]*model.BankPaymentLink, error) {
	result, err := b.DO.Find()
	return result.([]*model.BankPaymentLink), err
}

func (b bankPaymentLinkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BankPaymentLink, err error) {
	buf := make([]*model.BankPaymentLink, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bankPaymentLinkDo) FindInBatches(result *[]*model.BankPaymentLink, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bankPaymentLinkDo) Attrs(attrs ...field.AssignExpr) *bankPaymentLinkDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bankPaymentLinkDo) Assign(attrs ...field.AssignExpr) *bankPaymentLinkDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bankPaymentLinkDo) Joins(fields ...field.RelationField) *bankPaymentLinkDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bankPaymentLinkDo) Preload(fields ...field.RelationField) *bankPaymentLinkDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bankPaymentLinkDo) FirstOrInit() (*model.BankPaymentLink, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BankPaymentLink), nil
	}
}

func (b bankPaymentLinkDo) FirstOrCreate() (*model.BankPaymentLink, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BankPaymentLink), nil
	}
}

func (b bankPaymentLinkDo) FindByPage(offset int, limit int) (result []*model.BankPaymentLink, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bankPaymentLinkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bankPaymentLinkDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bankPaymentLinkDo) Delete(models ...*model.BankPaymentLink) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bankPaymentLinkDo) withDO(do gen.Dao) *bankPaymentLinkDo {
	b.DO = *do.(*gen.DO)
	return b
}