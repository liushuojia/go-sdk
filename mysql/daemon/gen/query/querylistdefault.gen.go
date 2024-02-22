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

func newQueryListDefault(db *gorm.DB, opts ...gen.DOOption) queryListDefault {
	_queryListDefault := queryListDefault{}

	_queryListDefault.queryListDefaultDo.UseDB(db, opts...)
	_queryListDefault.queryListDefaultDo.UseModel(&model.QueryListDefault{})

	tableName := _queryListDefault.queryListDefaultDo.TableName()
	_queryListDefault.ALL = field.NewAsterisk(tableName)
	_queryListDefault.ID = field.NewInt64(tableName, "id")
	_queryListDefault.CreatedAt = field.NewTime(tableName, "created_at")
	_queryListDefault.UpdatedAt = field.NewTime(tableName, "updated_at")
	_queryListDefault.DeletedAt = field.NewField(tableName, "deleted_at")
	_queryListDefault.QueryKey = field.NewString(tableName, "query_key")
	_queryListDefault.QueryTitle = field.NewString(tableName, "query_title")
	_queryListDefault.QueryColumns = field.NewString(tableName, "query_columns")
	_queryListDefault.QueryMap = field.NewString(tableName, "query_map")
	_queryListDefault.QueryMapMore = field.NewString(tableName, "query_map_more")
	_queryListDefault.QueryKeys = field.NewString(tableName, "query_keys")
	_queryListDefault.IndexOpen = field.NewInt64(tableName, "index_open")
	_queryListDefault.Weight = field.NewInt64(tableName, "weight")
	_queryListDefault.Status = field.NewInt64(tableName, "status")

	_queryListDefault.fillFieldMap()

	return _queryListDefault
}

type queryListDefault struct {
	queryListDefaultDo

	ALL          field.Asterisk
	ID           field.Int64  // 自动编号
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 更新时间
	DeletedAt    field.Field  // 删除时间
	QueryKey     field.String // 关键key 跟业务前端挂钩
	QueryTitle   field.String // 显示标题
	QueryColumns field.String // 字段数组 为空时跟默认一致 json
	QueryMap     field.String // 字段数组 查询 json
	QueryMapMore field.String // 字段数组 更多查询 json
	QueryKeys    field.String // 初始化查询结果 json
	IndexOpen    field.Int64  // 1 新窗打开详情  0 页面跳转
	Weight       field.Int64  // 权重排序 越大越靠前
	Status       field.Int64  // 状态 1有效 0 停用

	fieldMap map[string]field.Expr
}

func (q queryListDefault) Table(newTableName string) *queryListDefault {
	q.queryListDefaultDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q queryListDefault) As(alias string) *queryListDefault {
	q.queryListDefaultDo.DO = *(q.queryListDefaultDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *queryListDefault) updateTableName(table string) *queryListDefault {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt64(table, "id")
	q.CreatedAt = field.NewTime(table, "created_at")
	q.UpdatedAt = field.NewTime(table, "updated_at")
	q.DeletedAt = field.NewField(table, "deleted_at")
	q.QueryKey = field.NewString(table, "query_key")
	q.QueryTitle = field.NewString(table, "query_title")
	q.QueryColumns = field.NewString(table, "query_columns")
	q.QueryMap = field.NewString(table, "query_map")
	q.QueryMapMore = field.NewString(table, "query_map_more")
	q.QueryKeys = field.NewString(table, "query_keys")
	q.IndexOpen = field.NewInt64(table, "index_open")
	q.Weight = field.NewInt64(table, "weight")
	q.Status = field.NewInt64(table, "status")

	q.fillFieldMap()

	return q
}

func (q *queryListDefault) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *queryListDefault) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 13)
	q.fieldMap["id"] = q.ID
	q.fieldMap["created_at"] = q.CreatedAt
	q.fieldMap["updated_at"] = q.UpdatedAt
	q.fieldMap["deleted_at"] = q.DeletedAt
	q.fieldMap["query_key"] = q.QueryKey
	q.fieldMap["query_title"] = q.QueryTitle
	q.fieldMap["query_columns"] = q.QueryColumns
	q.fieldMap["query_map"] = q.QueryMap
	q.fieldMap["query_map_more"] = q.QueryMapMore
	q.fieldMap["query_keys"] = q.QueryKeys
	q.fieldMap["index_open"] = q.IndexOpen
	q.fieldMap["weight"] = q.Weight
	q.fieldMap["status"] = q.Status
}

func (q queryListDefault) clone(db *gorm.DB) queryListDefault {
	q.queryListDefaultDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q queryListDefault) replaceDB(db *gorm.DB) queryListDefault {
	q.queryListDefaultDo.ReplaceDB(db)
	return q
}

type queryListDefaultDo struct{ gen.DO }

func (q queryListDefaultDo) Debug() *queryListDefaultDo {
	return q.withDO(q.DO.Debug())
}

func (q queryListDefaultDo) WithContext(ctx context.Context) *queryListDefaultDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q queryListDefaultDo) ReadDB() *queryListDefaultDo {
	return q.Clauses(dbresolver.Read)
}

func (q queryListDefaultDo) WriteDB() *queryListDefaultDo {
	return q.Clauses(dbresolver.Write)
}

func (q queryListDefaultDo) Session(config *gorm.Session) *queryListDefaultDo {
	return q.withDO(q.DO.Session(config))
}

func (q queryListDefaultDo) Clauses(conds ...clause.Expression) *queryListDefaultDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q queryListDefaultDo) Returning(value interface{}, columns ...string) *queryListDefaultDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q queryListDefaultDo) Not(conds ...gen.Condition) *queryListDefaultDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q queryListDefaultDo) Or(conds ...gen.Condition) *queryListDefaultDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q queryListDefaultDo) Select(conds ...field.Expr) *queryListDefaultDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q queryListDefaultDo) Where(conds ...gen.Condition) *queryListDefaultDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q queryListDefaultDo) Order(conds ...field.Expr) *queryListDefaultDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q queryListDefaultDo) Distinct(cols ...field.Expr) *queryListDefaultDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q queryListDefaultDo) Omit(cols ...field.Expr) *queryListDefaultDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q queryListDefaultDo) Join(table schema.Tabler, on ...field.Expr) *queryListDefaultDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q queryListDefaultDo) LeftJoin(table schema.Tabler, on ...field.Expr) *queryListDefaultDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q queryListDefaultDo) RightJoin(table schema.Tabler, on ...field.Expr) *queryListDefaultDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q queryListDefaultDo) Group(cols ...field.Expr) *queryListDefaultDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q queryListDefaultDo) Having(conds ...gen.Condition) *queryListDefaultDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q queryListDefaultDo) Limit(limit int) *queryListDefaultDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q queryListDefaultDo) Offset(offset int) *queryListDefaultDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q queryListDefaultDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *queryListDefaultDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q queryListDefaultDo) Unscoped() *queryListDefaultDo {
	return q.withDO(q.DO.Unscoped())
}

func (q queryListDefaultDo) Create(values ...*model.QueryListDefault) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q queryListDefaultDo) CreateInBatches(values []*model.QueryListDefault, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q queryListDefaultDo) Save(values ...*model.QueryListDefault) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q queryListDefaultDo) First() (*model.QueryListDefault, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryListDefault), nil
	}
}

func (q queryListDefaultDo) Take() (*model.QueryListDefault, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryListDefault), nil
	}
}

func (q queryListDefaultDo) Last() (*model.QueryListDefault, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryListDefault), nil
	}
}

func (q queryListDefaultDo) Find() ([]*model.QueryListDefault, error) {
	result, err := q.DO.Find()
	return result.([]*model.QueryListDefault), err
}

func (q queryListDefaultDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QueryListDefault, err error) {
	buf := make([]*model.QueryListDefault, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q queryListDefaultDo) FindInBatches(result *[]*model.QueryListDefault, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q queryListDefaultDo) Attrs(attrs ...field.AssignExpr) *queryListDefaultDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q queryListDefaultDo) Assign(attrs ...field.AssignExpr) *queryListDefaultDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q queryListDefaultDo) Joins(fields ...field.RelationField) *queryListDefaultDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q queryListDefaultDo) Preload(fields ...field.RelationField) *queryListDefaultDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q queryListDefaultDo) FirstOrInit() (*model.QueryListDefault, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryListDefault), nil
	}
}

func (q queryListDefaultDo) FirstOrCreate() (*model.QueryListDefault, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryListDefault), nil
	}
}

func (q queryListDefaultDo) FindByPage(offset int, limit int) (result []*model.QueryListDefault, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q queryListDefaultDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q queryListDefaultDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q queryListDefaultDo) Delete(models ...*model.QueryListDefault) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *queryListDefaultDo) withDO(do gen.Dao) *queryListDefaultDo {
	q.DO = *do.(*gen.DO)
	return q
}