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

func newSystemAreaType(db *gorm.DB, opts ...gen.DOOption) systemAreaType {
	_systemAreaType := systemAreaType{}

	_systemAreaType.systemAreaTypeDo.UseDB(db, opts...)
	_systemAreaType.systemAreaTypeDo.UseModel(&model.SystemAreaType{})

	tableName := _systemAreaType.systemAreaTypeDo.TableName()
	_systemAreaType.ALL = field.NewAsterisk(tableName)
	_systemAreaType.ID = field.NewInt64(tableName, "id")
	_systemAreaType.CreatedAt = field.NewTime(tableName, "created_at")
	_systemAreaType.UpdatedAt = field.NewTime(tableName, "updated_at")
	_systemAreaType.DeletedAt = field.NewField(tableName, "deleted_at")
	_systemAreaType.TypeName = field.NewString(tableName, "type_name")
	_systemAreaType.DataType = field.NewInt64(tableName, "data_type")
	_systemAreaType.DataValueType = field.NewInt64(tableName, "data_value_type")
	_systemAreaType.TypeAbbreviation = field.NewString(tableName, "type_abbreviation")
	_systemAreaType.TypeClass = field.NewInt64(tableName, "type_class")
	_systemAreaType.Weight = field.NewInt64(tableName, "weight")
	_systemAreaType.Status = field.NewInt64(tableName, "status")
	_systemAreaType.TypeSeparator = field.NewString(tableName, "type_separator")
	_systemAreaType.LabelJSON = field.NewString(tableName, "label_json")

	_systemAreaType.fillFieldMap()

	return _systemAreaType
}

type systemAreaType struct {
	systemAreaTypeDo

	ALL              field.Asterisk
	ID               field.Int64  // 自动编号
	CreatedAt        field.Time   // 创建时间
	UpdatedAt        field.Time   // 更新时间
	DeletedAt        field.Field  // 删除时间
	TypeName         field.String // 分类名称
	DataType         field.Int64  // 默认 类型 0 默认分类选择  1 select  2 radio 3 checkbox
	DataValueType    field.Int64  // 值类型 0 字符串 1 数字
	TypeAbbreviation field.String // 分类标识
	TypeClass        field.Int64  // 分类类型 0 多级分类 1 单级分类
	Weight           field.Int64  // 权重排序 越大越靠前
	Status           field.Int64  // 状态 1有效 0 停用
	TypeSeparator    field.String // 分隔符
	LabelJSON        field.String // 扩展组件 自定义分类属性

	fieldMap map[string]field.Expr
}

func (s systemAreaType) Table(newTableName string) *systemAreaType {
	s.systemAreaTypeDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s systemAreaType) As(alias string) *systemAreaType {
	s.systemAreaTypeDo.DO = *(s.systemAreaTypeDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *systemAreaType) updateTableName(table string) *systemAreaType {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.TypeName = field.NewString(table, "type_name")
	s.DataType = field.NewInt64(table, "data_type")
	s.DataValueType = field.NewInt64(table, "data_value_type")
	s.TypeAbbreviation = field.NewString(table, "type_abbreviation")
	s.TypeClass = field.NewInt64(table, "type_class")
	s.Weight = field.NewInt64(table, "weight")
	s.Status = field.NewInt64(table, "status")
	s.TypeSeparator = field.NewString(table, "type_separator")
	s.LabelJSON = field.NewString(table, "label_json")

	s.fillFieldMap()

	return s
}

func (s *systemAreaType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *systemAreaType) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 13)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["type_name"] = s.TypeName
	s.fieldMap["data_type"] = s.DataType
	s.fieldMap["data_value_type"] = s.DataValueType
	s.fieldMap["type_abbreviation"] = s.TypeAbbreviation
	s.fieldMap["type_class"] = s.TypeClass
	s.fieldMap["weight"] = s.Weight
	s.fieldMap["status"] = s.Status
	s.fieldMap["type_separator"] = s.TypeSeparator
	s.fieldMap["label_json"] = s.LabelJSON
}

func (s systemAreaType) clone(db *gorm.DB) systemAreaType {
	s.systemAreaTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s systemAreaType) replaceDB(db *gorm.DB) systemAreaType {
	s.systemAreaTypeDo.ReplaceDB(db)
	return s
}

type systemAreaTypeDo struct{ gen.DO }

func (s systemAreaTypeDo) Debug() *systemAreaTypeDo {
	return s.withDO(s.DO.Debug())
}

func (s systemAreaTypeDo) WithContext(ctx context.Context) *systemAreaTypeDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemAreaTypeDo) ReadDB() *systemAreaTypeDo {
	return s.Clauses(dbresolver.Read)
}

func (s systemAreaTypeDo) WriteDB() *systemAreaTypeDo {
	return s.Clauses(dbresolver.Write)
}

func (s systemAreaTypeDo) Session(config *gorm.Session) *systemAreaTypeDo {
	return s.withDO(s.DO.Session(config))
}

func (s systemAreaTypeDo) Clauses(conds ...clause.Expression) *systemAreaTypeDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemAreaTypeDo) Returning(value interface{}, columns ...string) *systemAreaTypeDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemAreaTypeDo) Not(conds ...gen.Condition) *systemAreaTypeDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemAreaTypeDo) Or(conds ...gen.Condition) *systemAreaTypeDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemAreaTypeDo) Select(conds ...field.Expr) *systemAreaTypeDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemAreaTypeDo) Where(conds ...gen.Condition) *systemAreaTypeDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemAreaTypeDo) Order(conds ...field.Expr) *systemAreaTypeDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemAreaTypeDo) Distinct(cols ...field.Expr) *systemAreaTypeDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemAreaTypeDo) Omit(cols ...field.Expr) *systemAreaTypeDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemAreaTypeDo) Join(table schema.Tabler, on ...field.Expr) *systemAreaTypeDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemAreaTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *systemAreaTypeDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemAreaTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) *systemAreaTypeDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemAreaTypeDo) Group(cols ...field.Expr) *systemAreaTypeDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemAreaTypeDo) Having(conds ...gen.Condition) *systemAreaTypeDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemAreaTypeDo) Limit(limit int) *systemAreaTypeDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemAreaTypeDo) Offset(offset int) *systemAreaTypeDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemAreaTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *systemAreaTypeDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemAreaTypeDo) Unscoped() *systemAreaTypeDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemAreaTypeDo) Create(values ...*model.SystemAreaType) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemAreaTypeDo) CreateInBatches(values []*model.SystemAreaType, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemAreaTypeDo) Save(values ...*model.SystemAreaType) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemAreaTypeDo) First() (*model.SystemAreaType, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAreaType), nil
	}
}

func (s systemAreaTypeDo) Take() (*model.SystemAreaType, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAreaType), nil
	}
}

func (s systemAreaTypeDo) Last() (*model.SystemAreaType, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAreaType), nil
	}
}

func (s systemAreaTypeDo) Find() ([]*model.SystemAreaType, error) {
	result, err := s.DO.Find()
	return result.([]*model.SystemAreaType), err
}

func (s systemAreaTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemAreaType, err error) {
	buf := make([]*model.SystemAreaType, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemAreaTypeDo) FindInBatches(result *[]*model.SystemAreaType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemAreaTypeDo) Attrs(attrs ...field.AssignExpr) *systemAreaTypeDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemAreaTypeDo) Assign(attrs ...field.AssignExpr) *systemAreaTypeDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemAreaTypeDo) Joins(fields ...field.RelationField) *systemAreaTypeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemAreaTypeDo) Preload(fields ...field.RelationField) *systemAreaTypeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemAreaTypeDo) FirstOrInit() (*model.SystemAreaType, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAreaType), nil
	}
}

func (s systemAreaTypeDo) FirstOrCreate() (*model.SystemAreaType, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAreaType), nil
	}
}

func (s systemAreaTypeDo) FindByPage(offset int, limit int) (result []*model.SystemAreaType, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s systemAreaTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s systemAreaTypeDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s systemAreaTypeDo) Delete(models ...*model.SystemAreaType) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *systemAreaTypeDo) withDO(do gen.Dao) *systemAreaTypeDo {
	s.DO = *do.(*gen.DO)
	return s
}
