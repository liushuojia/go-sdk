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

func newCustomInput(db *gorm.DB, opts ...gen.DOOption) customInput {
	_customInput := customInput{}

	_customInput.customInputDo.UseDB(db, opts...)
	_customInput.customInputDo.UseModel(&model.CustomInput{})

	tableName := _customInput.customInputDo.TableName()
	_customInput.ALL = field.NewAsterisk(tableName)
	_customInput.ID = field.NewInt64(tableName, "id")
	_customInput.CreatedAt = field.NewTime(tableName, "created_at")
	_customInput.UpdatedAt = field.NewTime(tableName, "updated_at")
	_customInput.DeletedAt = field.NewField(tableName, "deleted_at")
	_customInput.DataCode = field.NewString(tableName, "data_code")
	_customInput.DataType = field.NewInt64(tableName, "data_type")
	_customInput.DataID = field.NewInt64(tableName, "data_id")
	_customInput.InputKey = field.NewString(tableName, "input_key")
	_customInput.InputValue = field.NewString(tableName, "input_value")

	_customInput.fillFieldMap()

	return _customInput
}

type customInput struct {
	customInputDo

	ALL        field.Asterisk
	ID         field.Int64  // 自动编号
	CreatedAt  field.Time   // 创建时间
	UpdatedAt  field.Time   // 更新时间
	DeletedAt  field.Field  // 删除时间
	DataCode   field.String // 业务所属编码
	DataType   field.Int64  // 业务所属类型id
	DataID     field.Int64  // 业务所属自增id
	InputKey   field.String // 输入框KEY
	InputValue field.String // 输入框VALUE

	fieldMap map[string]field.Expr
}

func (c customInput) Table(newTableName string) *customInput {
	c.customInputDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c customInput) As(alias string) *customInput {
	c.customInputDo.DO = *(c.customInputDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *customInput) updateTableName(table string) *customInput {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.DataCode = field.NewString(table, "data_code")
	c.DataType = field.NewInt64(table, "data_type")
	c.DataID = field.NewInt64(table, "data_id")
	c.InputKey = field.NewString(table, "input_key")
	c.InputValue = field.NewString(table, "input_value")

	c.fillFieldMap()

	return c
}

func (c *customInput) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *customInput) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["data_code"] = c.DataCode
	c.fieldMap["data_type"] = c.DataType
	c.fieldMap["data_id"] = c.DataID
	c.fieldMap["input_key"] = c.InputKey
	c.fieldMap["input_value"] = c.InputValue
}

func (c customInput) clone(db *gorm.DB) customInput {
	c.customInputDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c customInput) replaceDB(db *gorm.DB) customInput {
	c.customInputDo.ReplaceDB(db)
	return c
}

type customInputDo struct{ gen.DO }

func (c customInputDo) Debug() *customInputDo {
	return c.withDO(c.DO.Debug())
}

func (c customInputDo) WithContext(ctx context.Context) *customInputDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c customInputDo) ReadDB() *customInputDo {
	return c.Clauses(dbresolver.Read)
}

func (c customInputDo) WriteDB() *customInputDo {
	return c.Clauses(dbresolver.Write)
}

func (c customInputDo) Session(config *gorm.Session) *customInputDo {
	return c.withDO(c.DO.Session(config))
}

func (c customInputDo) Clauses(conds ...clause.Expression) *customInputDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c customInputDo) Returning(value interface{}, columns ...string) *customInputDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c customInputDo) Not(conds ...gen.Condition) *customInputDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c customInputDo) Or(conds ...gen.Condition) *customInputDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c customInputDo) Select(conds ...field.Expr) *customInputDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c customInputDo) Where(conds ...gen.Condition) *customInputDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c customInputDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *customInputDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c customInputDo) Order(conds ...field.Expr) *customInputDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c customInputDo) Distinct(cols ...field.Expr) *customInputDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c customInputDo) Omit(cols ...field.Expr) *customInputDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c customInputDo) Join(table schema.Tabler, on ...field.Expr) *customInputDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c customInputDo) LeftJoin(table schema.Tabler, on ...field.Expr) *customInputDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c customInputDo) RightJoin(table schema.Tabler, on ...field.Expr) *customInputDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c customInputDo) Group(cols ...field.Expr) *customInputDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c customInputDo) Having(conds ...gen.Condition) *customInputDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c customInputDo) Limit(limit int) *customInputDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c customInputDo) Offset(offset int) *customInputDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c customInputDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *customInputDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c customInputDo) Unscoped() *customInputDo {
	return c.withDO(c.DO.Unscoped())
}

func (c customInputDo) Create(values ...*model.CustomInput) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c customInputDo) CreateInBatches(values []*model.CustomInput, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c customInputDo) Save(values ...*model.CustomInput) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c customInputDo) First() (*model.CustomInput, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomInput), nil
	}
}

func (c customInputDo) Take() (*model.CustomInput, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomInput), nil
	}
}

func (c customInputDo) Last() (*model.CustomInput, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomInput), nil
	}
}

func (c customInputDo) Find() ([]*model.CustomInput, error) {
	result, err := c.DO.Find()
	return result.([]*model.CustomInput), err
}

func (c customInputDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CustomInput, err error) {
	buf := make([]*model.CustomInput, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c customInputDo) FindInBatches(result *[]*model.CustomInput, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c customInputDo) Attrs(attrs ...field.AssignExpr) *customInputDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c customInputDo) Assign(attrs ...field.AssignExpr) *customInputDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c customInputDo) Joins(fields ...field.RelationField) *customInputDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c customInputDo) Preload(fields ...field.RelationField) *customInputDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c customInputDo) FirstOrInit() (*model.CustomInput, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomInput), nil
	}
}

func (c customInputDo) FirstOrCreate() (*model.CustomInput, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomInput), nil
	}
}

func (c customInputDo) FindByPage(offset int, limit int) (result []*model.CustomInput, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c customInputDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c customInputDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c customInputDo) Delete(models ...*model.CustomInput) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *customInputDo) withDO(do gen.Dao) *customInputDo {
	c.DO = *do.(*gen.DO)
	return c
}
