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

func newInvoiceReceivedItem(db *gorm.DB, opts ...gen.DOOption) invoiceReceivedItem {
	_invoiceReceivedItem := invoiceReceivedItem{}

	_invoiceReceivedItem.invoiceReceivedItemDo.UseDB(db, opts...)
	_invoiceReceivedItem.invoiceReceivedItemDo.UseModel(&model.InvoiceReceivedItem{})

	tableName := _invoiceReceivedItem.invoiceReceivedItemDo.TableName()
	_invoiceReceivedItem.ALL = field.NewAsterisk(tableName)
	_invoiceReceivedItem.ID = field.NewInt64(tableName, "id")
	_invoiceReceivedItem.CreatedAt = field.NewTime(tableName, "created_at")
	_invoiceReceivedItem.UpdatedAt = field.NewTime(tableName, "updated_at")
	_invoiceReceivedItem.DeletedAt = field.NewField(tableName, "deleted_at")
	_invoiceReceivedItem.ReceivedID = field.NewInt64(tableName, "received_id")
	_invoiceReceivedItem.InvoiceTypeCode = field.NewString(tableName, "invoice_type_code")
	_invoiceReceivedItem.InvoiceTypeName = field.NewString(tableName, "invoice_type_name")
	_invoiceReceivedItem.InvoiceCode = field.NewString(tableName, "invoice_code")
	_invoiceReceivedItem.InvoiceNumber = field.NewString(tableName, "invoice_number")
	_invoiceReceivedItem.InvoiceDay = field.NewInt64(tableName, "invoice_day")
	_invoiceReceivedItem.InvoiceTotalPrice = field.NewInt64(tableName, "invoice_total_price")
	_invoiceReceivedItem.InvoiceProductPrice = field.NewInt64(tableName, "invoice_product_price")

	_invoiceReceivedItem.fillFieldMap()

	return _invoiceReceivedItem
}

type invoiceReceivedItem struct {
	invoiceReceivedItemDo

	ALL                 field.Asterisk
	ID                  field.Int64  // 自动编号
	CreatedAt           field.Time   // 创建时间
	UpdatedAt           field.Time   // 更新时间
	DeletedAt           field.Field  // 删除时间
	ReceivedID          field.Int64  // 回收发票id
	InvoiceTypeCode     field.String // 票据类型编码
	InvoiceTypeName     field.String // 票据类型名称
	InvoiceCode         field.String // 票据代码
	InvoiceNumber       field.String // 票据号码
	InvoiceDay          field.Int64  // 票据日期
	InvoiceTotalPrice   field.Int64  // 票据含税金额
	InvoiceProductPrice field.Int64  // 票据不含税金额

	fieldMap map[string]field.Expr
}

func (i invoiceReceivedItem) Table(newTableName string) *invoiceReceivedItem {
	i.invoiceReceivedItemDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i invoiceReceivedItem) As(alias string) *invoiceReceivedItem {
	i.invoiceReceivedItemDo.DO = *(i.invoiceReceivedItemDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *invoiceReceivedItem) updateTableName(table string) *invoiceReceivedItem {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.UpdatedAt = field.NewTime(table, "updated_at")
	i.DeletedAt = field.NewField(table, "deleted_at")
	i.ReceivedID = field.NewInt64(table, "received_id")
	i.InvoiceTypeCode = field.NewString(table, "invoice_type_code")
	i.InvoiceTypeName = field.NewString(table, "invoice_type_name")
	i.InvoiceCode = field.NewString(table, "invoice_code")
	i.InvoiceNumber = field.NewString(table, "invoice_number")
	i.InvoiceDay = field.NewInt64(table, "invoice_day")
	i.InvoiceTotalPrice = field.NewInt64(table, "invoice_total_price")
	i.InvoiceProductPrice = field.NewInt64(table, "invoice_product_price")

	i.fillFieldMap()

	return i
}

func (i *invoiceReceivedItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *invoiceReceivedItem) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 12)
	i.fieldMap["id"] = i.ID
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["updated_at"] = i.UpdatedAt
	i.fieldMap["deleted_at"] = i.DeletedAt
	i.fieldMap["received_id"] = i.ReceivedID
	i.fieldMap["invoice_type_code"] = i.InvoiceTypeCode
	i.fieldMap["invoice_type_name"] = i.InvoiceTypeName
	i.fieldMap["invoice_code"] = i.InvoiceCode
	i.fieldMap["invoice_number"] = i.InvoiceNumber
	i.fieldMap["invoice_day"] = i.InvoiceDay
	i.fieldMap["invoice_total_price"] = i.InvoiceTotalPrice
	i.fieldMap["invoice_product_price"] = i.InvoiceProductPrice
}

func (i invoiceReceivedItem) clone(db *gorm.DB) invoiceReceivedItem {
	i.invoiceReceivedItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i invoiceReceivedItem) replaceDB(db *gorm.DB) invoiceReceivedItem {
	i.invoiceReceivedItemDo.ReplaceDB(db)
	return i
}

type invoiceReceivedItemDo struct{ gen.DO }

func (i invoiceReceivedItemDo) Debug() *invoiceReceivedItemDo {
	return i.withDO(i.DO.Debug())
}

func (i invoiceReceivedItemDo) WithContext(ctx context.Context) *invoiceReceivedItemDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i invoiceReceivedItemDo) ReadDB() *invoiceReceivedItemDo {
	return i.Clauses(dbresolver.Read)
}

func (i invoiceReceivedItemDo) WriteDB() *invoiceReceivedItemDo {
	return i.Clauses(dbresolver.Write)
}

func (i invoiceReceivedItemDo) Session(config *gorm.Session) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Session(config))
}

func (i invoiceReceivedItemDo) Clauses(conds ...clause.Expression) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i invoiceReceivedItemDo) Returning(value interface{}, columns ...string) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i invoiceReceivedItemDo) Not(conds ...gen.Condition) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i invoiceReceivedItemDo) Or(conds ...gen.Condition) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i invoiceReceivedItemDo) Select(conds ...field.Expr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i invoiceReceivedItemDo) Where(conds ...gen.Condition) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i invoiceReceivedItemDo) Order(conds ...field.Expr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i invoiceReceivedItemDo) Distinct(cols ...field.Expr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i invoiceReceivedItemDo) Omit(cols ...field.Expr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i invoiceReceivedItemDo) Join(table schema.Tabler, on ...field.Expr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i invoiceReceivedItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i invoiceReceivedItemDo) RightJoin(table schema.Tabler, on ...field.Expr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i invoiceReceivedItemDo) Group(cols ...field.Expr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i invoiceReceivedItemDo) Having(conds ...gen.Condition) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i invoiceReceivedItemDo) Limit(limit int) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i invoiceReceivedItemDo) Offset(offset int) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i invoiceReceivedItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i invoiceReceivedItemDo) Unscoped() *invoiceReceivedItemDo {
	return i.withDO(i.DO.Unscoped())
}

func (i invoiceReceivedItemDo) Create(values ...*model.InvoiceReceivedItem) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i invoiceReceivedItemDo) CreateInBatches(values []*model.InvoiceReceivedItem, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i invoiceReceivedItemDo) Save(values ...*model.InvoiceReceivedItem) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i invoiceReceivedItemDo) First() (*model.InvoiceReceivedItem, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvoiceReceivedItem), nil
	}
}

func (i invoiceReceivedItemDo) Take() (*model.InvoiceReceivedItem, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvoiceReceivedItem), nil
	}
}

func (i invoiceReceivedItemDo) Last() (*model.InvoiceReceivedItem, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvoiceReceivedItem), nil
	}
}

func (i invoiceReceivedItemDo) Find() ([]*model.InvoiceReceivedItem, error) {
	result, err := i.DO.Find()
	return result.([]*model.InvoiceReceivedItem), err
}

func (i invoiceReceivedItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InvoiceReceivedItem, err error) {
	buf := make([]*model.InvoiceReceivedItem, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i invoiceReceivedItemDo) FindInBatches(result *[]*model.InvoiceReceivedItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i invoiceReceivedItemDo) Attrs(attrs ...field.AssignExpr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i invoiceReceivedItemDo) Assign(attrs ...field.AssignExpr) *invoiceReceivedItemDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i invoiceReceivedItemDo) Joins(fields ...field.RelationField) *invoiceReceivedItemDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i invoiceReceivedItemDo) Preload(fields ...field.RelationField) *invoiceReceivedItemDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i invoiceReceivedItemDo) FirstOrInit() (*model.InvoiceReceivedItem, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvoiceReceivedItem), nil
	}
}

func (i invoiceReceivedItemDo) FirstOrCreate() (*model.InvoiceReceivedItem, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvoiceReceivedItem), nil
	}
}

func (i invoiceReceivedItemDo) FindByPage(offset int, limit int) (result []*model.InvoiceReceivedItem, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i invoiceReceivedItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i invoiceReceivedItemDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i invoiceReceivedItemDo) Delete(models ...*model.InvoiceReceivedItem) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *invoiceReceivedItemDo) withDO(do gen.Dao) *invoiceReceivedItemDo {
	i.DO = *do.(*gen.DO)
	return i
}
