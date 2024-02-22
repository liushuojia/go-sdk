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

func newOrderArchiveItem(db *gorm.DB, opts ...gen.DOOption) orderArchiveItem {
	_orderArchiveItem := orderArchiveItem{}

	_orderArchiveItem.orderArchiveItemDo.UseDB(db, opts...)
	_orderArchiveItem.orderArchiveItemDo.UseModel(&model.OrderArchiveItem{})

	tableName := _orderArchiveItem.orderArchiveItemDo.TableName()
	_orderArchiveItem.ALL = field.NewAsterisk(tableName)
	_orderArchiveItem.ID = field.NewInt64(tableName, "id")
	_orderArchiveItem.CreatedAt = field.NewTime(tableName, "created_at")
	_orderArchiveItem.UpdatedAt = field.NewTime(tableName, "updated_at")
	_orderArchiveItem.DeletedAt = field.NewField(tableName, "deleted_at")
	_orderArchiveItem.ArchiveID = field.NewInt64(tableName, "archive_id")
	_orderArchiveItem.ArchiveType = field.NewInt64(tableName, "archive_type")
	_orderArchiveItem.PaymentType = field.NewInt64(tableName, "payment_type")
	_orderArchiveItem.ArchiveName = field.NewInt64(tableName, "archive_name")
	_orderArchiveItem.ProfitFlag = field.NewInt64(tableName, "profit_flag")
	_orderArchiveItem.CurrencyAbbreviation = field.NewString(tableName, "currency_abbreviation")
	_orderArchiveItem.CurrencyPrice = field.NewInt64(tableName, "currency_price")
	_orderArchiveItem.TotalPrice = field.NewInt64(tableName, "total_price")
	_orderArchiveItem.BankPaymentTitle = field.NewString(tableName, "bank_payment_title")
	_orderArchiveItem.PaymentProjectName = field.NewString(tableName, "payment_project_name")
	_orderArchiveItem.PayBankName = field.NewString(tableName, "pay_bank_name")
	_orderArchiveItem.ReceiveBankName = field.NewString(tableName, "receive_bank_name")
	_orderArchiveItem.ProviderName = field.NewString(tableName, "provider_name")
	_orderArchiveItem.PaymentTime = field.NewInt64(tableName, "payment_time")

	_orderArchiveItem.fillFieldMap()

	return _orderArchiveItem
}

type orderArchiveItem struct {
	orderArchiveItemDo

	ALL                  field.Asterisk
	ID                   field.Int64  // 自动编号
	CreatedAt            field.Time   // 创建时间
	UpdatedAt            field.Time   // 更新时间
	DeletedAt            field.Field  // 删除时间
	ArchiveID            field.Int64  // 归档id
	ArchiveType          field.Int64  // 类型 0 客户 1 供应商 2 回收发票
	PaymentType          field.Int64  // 类型 0 收款  1 付款  2 收款手续费 3 付款手续费
	ArchiveName          field.Int64  // 客户名称/供应商名称
	ProfitFlag           field.Int64  // 利润 0 影响利润 1 不影响利润
	CurrencyAbbreviation field.String // 币种简写
	CurrencyPrice        field.Int64  // 币种价格 单位分
	TotalPrice           field.Int64  // 折合人民币 单位分
	BankPaymentTitle     field.String // 收付款 名称
	PaymentProjectName   field.String // 收付款项目 名称
	PayBankName          field.String // 付款账户 开户行
	ReceiveBankName      field.String // 收款账户 开户行
	ProviderName         field.String // 供应商名称
	PaymentTime          field.Int64  // 付款/收款 时间

	fieldMap map[string]field.Expr
}

func (o orderArchiveItem) Table(newTableName string) *orderArchiveItem {
	o.orderArchiveItemDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o orderArchiveItem) As(alias string) *orderArchiveItem {
	o.orderArchiveItemDo.DO = *(o.orderArchiveItemDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *orderArchiveItem) updateTableName(table string) *orderArchiveItem {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")
	o.ArchiveID = field.NewInt64(table, "archive_id")
	o.ArchiveType = field.NewInt64(table, "archive_type")
	o.PaymentType = field.NewInt64(table, "payment_type")
	o.ArchiveName = field.NewInt64(table, "archive_name")
	o.ProfitFlag = field.NewInt64(table, "profit_flag")
	o.CurrencyAbbreviation = field.NewString(table, "currency_abbreviation")
	o.CurrencyPrice = field.NewInt64(table, "currency_price")
	o.TotalPrice = field.NewInt64(table, "total_price")
	o.BankPaymentTitle = field.NewString(table, "bank_payment_title")
	o.PaymentProjectName = field.NewString(table, "payment_project_name")
	o.PayBankName = field.NewString(table, "pay_bank_name")
	o.ReceiveBankName = field.NewString(table, "receive_bank_name")
	o.ProviderName = field.NewString(table, "provider_name")
	o.PaymentTime = field.NewInt64(table, "payment_time")

	o.fillFieldMap()

	return o
}

func (o *orderArchiveItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *orderArchiveItem) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 18)
	o.fieldMap["id"] = o.ID
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
	o.fieldMap["archive_id"] = o.ArchiveID
	o.fieldMap["archive_type"] = o.ArchiveType
	o.fieldMap["payment_type"] = o.PaymentType
	o.fieldMap["archive_name"] = o.ArchiveName
	o.fieldMap["profit_flag"] = o.ProfitFlag
	o.fieldMap["currency_abbreviation"] = o.CurrencyAbbreviation
	o.fieldMap["currency_price"] = o.CurrencyPrice
	o.fieldMap["total_price"] = o.TotalPrice
	o.fieldMap["bank_payment_title"] = o.BankPaymentTitle
	o.fieldMap["payment_project_name"] = o.PaymentProjectName
	o.fieldMap["pay_bank_name"] = o.PayBankName
	o.fieldMap["receive_bank_name"] = o.ReceiveBankName
	o.fieldMap["provider_name"] = o.ProviderName
	o.fieldMap["payment_time"] = o.PaymentTime
}

func (o orderArchiveItem) clone(db *gorm.DB) orderArchiveItem {
	o.orderArchiveItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o orderArchiveItem) replaceDB(db *gorm.DB) orderArchiveItem {
	o.orderArchiveItemDo.ReplaceDB(db)
	return o
}

type orderArchiveItemDo struct{ gen.DO }

func (o orderArchiveItemDo) Debug() *orderArchiveItemDo {
	return o.withDO(o.DO.Debug())
}

func (o orderArchiveItemDo) WithContext(ctx context.Context) *orderArchiveItemDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderArchiveItemDo) ReadDB() *orderArchiveItemDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderArchiveItemDo) WriteDB() *orderArchiveItemDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderArchiveItemDo) Session(config *gorm.Session) *orderArchiveItemDo {
	return o.withDO(o.DO.Session(config))
}

func (o orderArchiveItemDo) Clauses(conds ...clause.Expression) *orderArchiveItemDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderArchiveItemDo) Returning(value interface{}, columns ...string) *orderArchiveItemDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderArchiveItemDo) Not(conds ...gen.Condition) *orderArchiveItemDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderArchiveItemDo) Or(conds ...gen.Condition) *orderArchiveItemDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderArchiveItemDo) Select(conds ...field.Expr) *orderArchiveItemDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderArchiveItemDo) Where(conds ...gen.Condition) *orderArchiveItemDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderArchiveItemDo) Order(conds ...field.Expr) *orderArchiveItemDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderArchiveItemDo) Distinct(cols ...field.Expr) *orderArchiveItemDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderArchiveItemDo) Omit(cols ...field.Expr) *orderArchiveItemDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderArchiveItemDo) Join(table schema.Tabler, on ...field.Expr) *orderArchiveItemDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderArchiveItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) *orderArchiveItemDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderArchiveItemDo) RightJoin(table schema.Tabler, on ...field.Expr) *orderArchiveItemDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderArchiveItemDo) Group(cols ...field.Expr) *orderArchiveItemDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderArchiveItemDo) Having(conds ...gen.Condition) *orderArchiveItemDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderArchiveItemDo) Limit(limit int) *orderArchiveItemDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderArchiveItemDo) Offset(offset int) *orderArchiveItemDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderArchiveItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *orderArchiveItemDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderArchiveItemDo) Unscoped() *orderArchiveItemDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderArchiveItemDo) Create(values ...*model.OrderArchiveItem) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderArchiveItemDo) CreateInBatches(values []*model.OrderArchiveItem, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderArchiveItemDo) Save(values ...*model.OrderArchiveItem) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderArchiveItemDo) First() (*model.OrderArchiveItem, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderArchiveItem), nil
	}
}

func (o orderArchiveItemDo) Take() (*model.OrderArchiveItem, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderArchiveItem), nil
	}
}

func (o orderArchiveItemDo) Last() (*model.OrderArchiveItem, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderArchiveItem), nil
	}
}

func (o orderArchiveItemDo) Find() ([]*model.OrderArchiveItem, error) {
	result, err := o.DO.Find()
	return result.([]*model.OrderArchiveItem), err
}

func (o orderArchiveItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrderArchiveItem, err error) {
	buf := make([]*model.OrderArchiveItem, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderArchiveItemDo) FindInBatches(result *[]*model.OrderArchiveItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderArchiveItemDo) Attrs(attrs ...field.AssignExpr) *orderArchiveItemDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderArchiveItemDo) Assign(attrs ...field.AssignExpr) *orderArchiveItemDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderArchiveItemDo) Joins(fields ...field.RelationField) *orderArchiveItemDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderArchiveItemDo) Preload(fields ...field.RelationField) *orderArchiveItemDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderArchiveItemDo) FirstOrInit() (*model.OrderArchiveItem, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderArchiveItem), nil
	}
}

func (o orderArchiveItemDo) FirstOrCreate() (*model.OrderArchiveItem, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderArchiveItem), nil
	}
}

func (o orderArchiveItemDo) FindByPage(offset int, limit int) (result []*model.OrderArchiveItem, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o orderArchiveItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderArchiveItemDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderArchiveItemDo) Delete(models ...*model.OrderArchiveItem) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderArchiveItemDo) withDO(do gen.Dao) *orderArchiveItemDo {
	o.DO = *do.(*gen.DO)
	return o
}
