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

func newQueryList(db *gorm.DB, opts ...gen.DOOption) queryList {
	_queryList := queryList{}

	_queryList.queryListDo.UseDB(db, opts...)
	_queryList.queryListDo.UseModel(&model.QueryList{})

	tableName := _queryList.queryListDo.TableName()
	_queryList.ALL = field.NewAsterisk(tableName)
	_queryList.ID = field.NewInt64(tableName, "id")
	_queryList.CreatedAt = field.NewTime(tableName, "created_at")
	_queryList.UpdatedAt = field.NewTime(tableName, "updated_at")
	_queryList.DeletedAt = field.NewField(tableName, "deleted_at")
	_queryList.UserID = field.NewInt64(tableName, "user_id")
	_queryList.QueryKey = field.NewString(tableName, "query_key")
	_queryList.QueryTitle = field.NewString(tableName, "query_title")
	_queryList.QueryColumns = field.NewString(tableName, "query_columns")
	_queryList.QueryMap = field.NewString(tableName, "query_map")
	_queryList.QueryMapMore = field.NewString(tableName, "query_map_more")
	_queryList.QueryKeys = field.NewString(tableName, "query_keys")
	_queryList.IndexOpen = field.NewInt64(tableName, "index_open")
	_queryList.Weight = field.NewInt64(tableName, "weight")
	_queryList.Status = field.NewInt64(tableName, "status")

	_queryList.fillFieldMap()

	return _queryList
}

type queryList struct {
	queryListDo

	ALL          field.Asterisk
	ID           field.Int64  // 自动编号
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 更新时间
	DeletedAt    field.Field  // 删除时间
	UserID       field.Int64  // 用户编号
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

func (q queryList) Table(newTableName string) *queryList {
	q.queryListDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q queryList) As(alias string) *queryList {
	q.queryListDo.DO = *(q.queryListDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *queryList) updateTableName(table string) *queryList {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt64(table, "id")
	q.CreatedAt = field.NewTime(table, "created_at")
	q.UpdatedAt = field.NewTime(table, "updated_at")
	q.DeletedAt = field.NewField(table, "deleted_at")
	q.UserID = field.NewInt64(table, "user_id")
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

func (q *queryList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *queryList) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 14)
	q.fieldMap["id"] = q.ID
	q.fieldMap["created_at"] = q.CreatedAt
	q.fieldMap["updated_at"] = q.UpdatedAt
	q.fieldMap["deleted_at"] = q.DeletedAt
	q.fieldMap["user_id"] = q.UserID
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

func (q queryList) clone(db *gorm.DB) queryList {
	q.queryListDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q queryList) replaceDB(db *gorm.DB) queryList {
	q.queryListDo.ReplaceDB(db)
	return q
}

type queryListDo struct{ gen.DO }

func (q queryListDo) Debug() *queryListDo {
	return q.withDO(q.DO.Debug())
}

func (q queryListDo) WithContext(ctx context.Context) *queryListDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q queryListDo) ReadDB() *queryListDo {
	return q.Clauses(dbresolver.Read)
}

func (q queryListDo) WriteDB() *queryListDo {
	return q.Clauses(dbresolver.Write)
}

func (q queryListDo) Session(config *gorm.Session) *queryListDo {
	return q.withDO(q.DO.Session(config))
}

func (q queryListDo) Clauses(conds ...clause.Expression) *queryListDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q queryListDo) Returning(value interface{}, columns ...string) *queryListDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q queryListDo) Not(conds ...gen.Condition) *queryListDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q queryListDo) Or(conds ...gen.Condition) *queryListDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q queryListDo) Select(conds ...field.Expr) *queryListDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q queryListDo) Where(conds ...gen.Condition) *queryListDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q queryListDo) Order(conds ...field.Expr) *queryListDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q queryListDo) Distinct(cols ...field.Expr) *queryListDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q queryListDo) Omit(cols ...field.Expr) *queryListDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q queryListDo) Join(table schema.Tabler, on ...field.Expr) *queryListDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q queryListDo) LeftJoin(table schema.Tabler, on ...field.Expr) *queryListDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q queryListDo) RightJoin(table schema.Tabler, on ...field.Expr) *queryListDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q queryListDo) Group(cols ...field.Expr) *queryListDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q queryListDo) Having(conds ...gen.Condition) *queryListDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q queryListDo) Limit(limit int) *queryListDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q queryListDo) Offset(offset int) *queryListDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q queryListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *queryListDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q queryListDo) Unscoped() *queryListDo {
	return q.withDO(q.DO.Unscoped())
}

func (q queryListDo) Create(values ...*model.QueryList) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q queryListDo) CreateInBatches(values []*model.QueryList, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q queryListDo) Save(values ...*model.QueryList) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q queryListDo) First() (*model.QueryList, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryList), nil
	}
}

func (q queryListDo) Take() (*model.QueryList, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryList), nil
	}
}

func (q queryListDo) Last() (*model.QueryList, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryList), nil
	}
}

func (q queryListDo) Find() ([]*model.QueryList, error) {
	result, err := q.DO.Find()
	return result.([]*model.QueryList), err
}

func (q queryListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QueryList, err error) {
	buf := make([]*model.QueryList, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q queryListDo) FindInBatches(result *[]*model.QueryList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q queryListDo) Attrs(attrs ...field.AssignExpr) *queryListDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q queryListDo) Assign(attrs ...field.AssignExpr) *queryListDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q queryListDo) Joins(fields ...field.RelationField) *queryListDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q queryListDo) Preload(fields ...field.RelationField) *queryListDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q queryListDo) FirstOrInit() (*model.QueryList, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryList), nil
	}
}

func (q queryListDo) FirstOrCreate() (*model.QueryList, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryList), nil
	}
}

func (q queryListDo) FindByPage(offset int, limit int) (result []*model.QueryList, count int64, err error) {
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

func (q queryListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q queryListDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q queryListDo) Delete(models ...*model.QueryList) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *queryListDo) withDO(do gen.Dao) *queryListDo {
	q.DO = *do.(*gen.DO)
	return q
}
