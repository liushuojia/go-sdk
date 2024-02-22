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

func newFollowPlanResultLabel(db *gorm.DB, opts ...gen.DOOption) followPlanResultLabel {
	_followPlanResultLabel := followPlanResultLabel{}

	_followPlanResultLabel.followPlanResultLabelDo.UseDB(db, opts...)
	_followPlanResultLabel.followPlanResultLabelDo.UseModel(&model.FollowPlanResultLabel{})

	tableName := _followPlanResultLabel.followPlanResultLabelDo.TableName()
	_followPlanResultLabel.ALL = field.NewAsterisk(tableName)
	_followPlanResultLabel.ID = field.NewInt64(tableName, "id")
	_followPlanResultLabel.CreatedAt = field.NewTime(tableName, "created_at")
	_followPlanResultLabel.UpdatedAt = field.NewTime(tableName, "updated_at")
	_followPlanResultLabel.DeletedAt = field.NewField(tableName, "deleted_at")
	_followPlanResultLabel.AutoFollowID = field.NewInt64(tableName, "auto_follow_id")
	_followPlanResultLabel.FollowPlanID = field.NewInt64(tableName, "follow_plan_id")
	_followPlanResultLabel.LabelName = field.NewString(tableName, "label_name")
	_followPlanResultLabel.LabelCode = field.NewString(tableName, "label_code")
	_followPlanResultLabel.LabelAbbreviation = field.NewString(tableName, "label_abbreviation")

	_followPlanResultLabel.fillFieldMap()

	return _followPlanResultLabel
}

type followPlanResultLabel struct {
	followPlanResultLabelDo

	ALL               field.Asterisk
	ID                field.Int64  // 自动编号
	CreatedAt         field.Time   // 创建时间
	UpdatedAt         field.Time   // 更新时间
	DeletedAt         field.Field  // 删除时间
	AutoFollowID      field.Int64  // 系统客户计划跟进id
	FollowPlanID      field.Int64  // 计划跟进id
	LabelName         field.String // 标签名称
	LabelCode         field.String // 标签编码
	LabelAbbreviation field.String // 标签简称 方便后续增加更多维度

	fieldMap map[string]field.Expr
}

func (f followPlanResultLabel) Table(newTableName string) *followPlanResultLabel {
	f.followPlanResultLabelDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f followPlanResultLabel) As(alias string) *followPlanResultLabel {
	f.followPlanResultLabelDo.DO = *(f.followPlanResultLabelDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *followPlanResultLabel) updateTableName(table string) *followPlanResultLabel {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")
	f.AutoFollowID = field.NewInt64(table, "auto_follow_id")
	f.FollowPlanID = field.NewInt64(table, "follow_plan_id")
	f.LabelName = field.NewString(table, "label_name")
	f.LabelCode = field.NewString(table, "label_code")
	f.LabelAbbreviation = field.NewString(table, "label_abbreviation")

	f.fillFieldMap()

	return f
}

func (f *followPlanResultLabel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *followPlanResultLabel) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 9)
	f.fieldMap["id"] = f.ID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
	f.fieldMap["auto_follow_id"] = f.AutoFollowID
	f.fieldMap["follow_plan_id"] = f.FollowPlanID
	f.fieldMap["label_name"] = f.LabelName
	f.fieldMap["label_code"] = f.LabelCode
	f.fieldMap["label_abbreviation"] = f.LabelAbbreviation
}

func (f followPlanResultLabel) clone(db *gorm.DB) followPlanResultLabel {
	f.followPlanResultLabelDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f followPlanResultLabel) replaceDB(db *gorm.DB) followPlanResultLabel {
	f.followPlanResultLabelDo.ReplaceDB(db)
	return f
}

type followPlanResultLabelDo struct{ gen.DO }

func (f followPlanResultLabelDo) Debug() *followPlanResultLabelDo {
	return f.withDO(f.DO.Debug())
}

func (f followPlanResultLabelDo) WithContext(ctx context.Context) *followPlanResultLabelDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f followPlanResultLabelDo) ReadDB() *followPlanResultLabelDo {
	return f.Clauses(dbresolver.Read)
}

func (f followPlanResultLabelDo) WriteDB() *followPlanResultLabelDo {
	return f.Clauses(dbresolver.Write)
}

func (f followPlanResultLabelDo) Session(config *gorm.Session) *followPlanResultLabelDo {
	return f.withDO(f.DO.Session(config))
}

func (f followPlanResultLabelDo) Clauses(conds ...clause.Expression) *followPlanResultLabelDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f followPlanResultLabelDo) Returning(value interface{}, columns ...string) *followPlanResultLabelDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f followPlanResultLabelDo) Not(conds ...gen.Condition) *followPlanResultLabelDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f followPlanResultLabelDo) Or(conds ...gen.Condition) *followPlanResultLabelDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f followPlanResultLabelDo) Select(conds ...field.Expr) *followPlanResultLabelDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f followPlanResultLabelDo) Where(conds ...gen.Condition) *followPlanResultLabelDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f followPlanResultLabelDo) Order(conds ...field.Expr) *followPlanResultLabelDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f followPlanResultLabelDo) Distinct(cols ...field.Expr) *followPlanResultLabelDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f followPlanResultLabelDo) Omit(cols ...field.Expr) *followPlanResultLabelDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f followPlanResultLabelDo) Join(table schema.Tabler, on ...field.Expr) *followPlanResultLabelDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f followPlanResultLabelDo) LeftJoin(table schema.Tabler, on ...field.Expr) *followPlanResultLabelDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f followPlanResultLabelDo) RightJoin(table schema.Tabler, on ...field.Expr) *followPlanResultLabelDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f followPlanResultLabelDo) Group(cols ...field.Expr) *followPlanResultLabelDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f followPlanResultLabelDo) Having(conds ...gen.Condition) *followPlanResultLabelDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f followPlanResultLabelDo) Limit(limit int) *followPlanResultLabelDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f followPlanResultLabelDo) Offset(offset int) *followPlanResultLabelDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f followPlanResultLabelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *followPlanResultLabelDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f followPlanResultLabelDo) Unscoped() *followPlanResultLabelDo {
	return f.withDO(f.DO.Unscoped())
}

func (f followPlanResultLabelDo) Create(values ...*model.FollowPlanResultLabel) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f followPlanResultLabelDo) CreateInBatches(values []*model.FollowPlanResultLabel, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f followPlanResultLabelDo) Save(values ...*model.FollowPlanResultLabel) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f followPlanResultLabelDo) First() (*model.FollowPlanResultLabel, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowPlanResultLabel), nil
	}
}

func (f followPlanResultLabelDo) Take() (*model.FollowPlanResultLabel, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowPlanResultLabel), nil
	}
}

func (f followPlanResultLabelDo) Last() (*model.FollowPlanResultLabel, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowPlanResultLabel), nil
	}
}

func (f followPlanResultLabelDo) Find() ([]*model.FollowPlanResultLabel, error) {
	result, err := f.DO.Find()
	return result.([]*model.FollowPlanResultLabel), err
}

func (f followPlanResultLabelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FollowPlanResultLabel, err error) {
	buf := make([]*model.FollowPlanResultLabel, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f followPlanResultLabelDo) FindInBatches(result *[]*model.FollowPlanResultLabel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f followPlanResultLabelDo) Attrs(attrs ...field.AssignExpr) *followPlanResultLabelDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f followPlanResultLabelDo) Assign(attrs ...field.AssignExpr) *followPlanResultLabelDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f followPlanResultLabelDo) Joins(fields ...field.RelationField) *followPlanResultLabelDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f followPlanResultLabelDo) Preload(fields ...field.RelationField) *followPlanResultLabelDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f followPlanResultLabelDo) FirstOrInit() (*model.FollowPlanResultLabel, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowPlanResultLabel), nil
	}
}

func (f followPlanResultLabelDo) FirstOrCreate() (*model.FollowPlanResultLabel, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowPlanResultLabel), nil
	}
}

func (f followPlanResultLabelDo) FindByPage(offset int, limit int) (result []*model.FollowPlanResultLabel, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f followPlanResultLabelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f followPlanResultLabelDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f followPlanResultLabelDo) Delete(models ...*model.FollowPlanResultLabel) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *followPlanResultLabelDo) withDO(do gen.Dao) *followPlanResultLabelDo {
	f.DO = *do.(*gen.DO)
	return f
}
