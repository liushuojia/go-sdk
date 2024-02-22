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

func newFollowDesc(db *gorm.DB, opts ...gen.DOOption) followDesc {
	_followDesc := followDesc{}

	_followDesc.followDescDo.UseDB(db, opts...)
	_followDesc.followDescDo.UseModel(&model.FollowDesc{})

	tableName := _followDesc.followDescDo.TableName()
	_followDesc.ALL = field.NewAsterisk(tableName)
	_followDesc.ID = field.NewInt64(tableName, "id")
	_followDesc.CreatedAt = field.NewTime(tableName, "created_at")
	_followDesc.UpdatedAt = field.NewTime(tableName, "updated_at")
	_followDesc.DeletedAt = field.NewField(tableName, "deleted_at")
	_followDesc.FollowType = field.NewString(tableName, "follow_type")
	_followDesc.FollowTypeCode = field.NewString(tableName, "follow_type_code")
	_followDesc.FollowTypeIco = field.NewString(tableName, "follow_type_ico")
	_followDesc.FollowTypeAbbreviation = field.NewString(tableName, "follow_type_abbreviation")
	_followDesc.DataCode = field.NewString(tableName, "data_code")
	_followDesc.DataType = field.NewInt64(tableName, "data_type")
	_followDesc.DataID = field.NewInt64(tableName, "data_id")
	_followDesc.BelongOp = field.NewInt64(tableName, "belong_op")
	_followDesc.BelongOpName = field.NewString(tableName, "belong_op_name")
	_followDesc.BelongDepartmentCode = field.NewString(tableName, "belong_department_code")
	_followDesc.BelongDepartmentName = field.NewString(tableName, "belong_department_name")
	_followDesc.FollowTitle = field.NewString(tableName, "follow_title")
	_followDesc.FollowContent = field.NewString(tableName, "follow_content")
	_followDesc.FollowPlanDayFlag = field.NewInt64(tableName, "follow_plan_day_flag")
	_followDesc.FollowPlanDay = field.NewInt64(tableName, "follow_plan_day")
	_followDesc.FollowPlanDayID = field.NewInt64(tableName, "follow_plan_day_id")
	_followDesc.AutoFollowID = field.NewInt64(tableName, "auto_follow_id")
	_followDesc.OpChangeFlag = field.NewInt64(tableName, "op_change_flag")
	_followDesc.CompanyID = field.NewInt64(tableName, "company_id")
	_followDesc.CompanyName = field.NewString(tableName, "company_name")
	_followDesc.UserID = field.NewInt64(tableName, "user_id")
	_followDesc.UserRealname = field.NewString(tableName, "user_realname")
	_followDesc.ResultJSON = field.NewString(tableName, "result_json")
	_followDesc.SpendTimeCode = field.NewString(tableName, "spend_time_code")
	_followDesc.SpendTimeName = field.NewString(tableName, "spend_time_name")
	_followDesc.ProcessingResultCode = field.NewString(tableName, "processing_result_code")
	_followDesc.ProcessingResultName = field.NewString(tableName, "processing_result_name")
	_followDesc.FollowUpTypeCode = field.NewString(tableName, "follow_up_type_code")
	_followDesc.FollowUpTypeName = field.NewString(tableName, "follow_up_type_name")
	_followDesc.Module = field.NewString(tableName, "module")

	_followDesc.fillFieldMap()

	return _followDesc
}

type followDesc struct {
	followDescDo

	ALL                    field.Asterisk
	ID                     field.Int64  // 自动编号
	CreatedAt              field.Time   // 创建时间
	UpdatedAt              field.Time   // 更新时间
	DeletedAt              field.Field  // 删除时间
	FollowType             field.String // 跟进类型
	FollowTypeCode         field.String // 跟进类型编码
	FollowTypeIco          field.String // 跟进类型图标
	FollowTypeAbbreviation field.String // 跟进类型分类索引
	DataCode               field.String // 业务所属编码
	DataType               field.Int64  // 业务所属类型id
	DataID                 field.Int64  // 业务所属自增id
	BelongOp               field.Int64  // 所属OP
	BelongOpName           field.String // 所属op姓名
	BelongDepartmentCode   field.String // 所属部门编码
	BelongDepartmentName   field.String // 所属部门名称
	FollowTitle            field.String // 跟进标题
	FollowContent          field.String // 跟进内容
	FollowPlanDayFlag      field.Int64  // 远期跟进标识
	FollowPlanDay          field.Int64  // 远期跟进日期
	FollowPlanDayID        field.Int64  // 远期跟进 索引id
	AutoFollowID           field.Int64  // 系统计划跟进id 索引id
	OpChangeFlag           field.Int64  // 客服转交标识
	CompanyID              field.Int64  // 所属公司id
	CompanyName            field.String // 所属公司名字
	UserID                 field.Int64  // 所属用户id
	UserRealname           field.String // 所属用户名字
	ResultJSON             field.String // 结果标签 json 数组类型 如[客户明确拒绝,无法联系客户]
	SpendTimeCode          field.String // 处理时间编码
	SpendTimeName          field.String // 处理时间
	ProcessingResultCode   field.String // 处理结果编码
	ProcessingResultName   field.String // 处理结果
	FollowUpTypeCode       field.String // 跟进结果编码
	FollowUpTypeName       field.String // 跟进结果名称
	Module                 field.String // 所属模块

	fieldMap map[string]field.Expr
}

func (f followDesc) Table(newTableName string) *followDesc {
	f.followDescDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f followDesc) As(alias string) *followDesc {
	f.followDescDo.DO = *(f.followDescDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *followDesc) updateTableName(table string) *followDesc {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")
	f.FollowType = field.NewString(table, "follow_type")
	f.FollowTypeCode = field.NewString(table, "follow_type_code")
	f.FollowTypeIco = field.NewString(table, "follow_type_ico")
	f.FollowTypeAbbreviation = field.NewString(table, "follow_type_abbreviation")
	f.DataCode = field.NewString(table, "data_code")
	f.DataType = field.NewInt64(table, "data_type")
	f.DataID = field.NewInt64(table, "data_id")
	f.BelongOp = field.NewInt64(table, "belong_op")
	f.BelongOpName = field.NewString(table, "belong_op_name")
	f.BelongDepartmentCode = field.NewString(table, "belong_department_code")
	f.BelongDepartmentName = field.NewString(table, "belong_department_name")
	f.FollowTitle = field.NewString(table, "follow_title")
	f.FollowContent = field.NewString(table, "follow_content")
	f.FollowPlanDayFlag = field.NewInt64(table, "follow_plan_day_flag")
	f.FollowPlanDay = field.NewInt64(table, "follow_plan_day")
	f.FollowPlanDayID = field.NewInt64(table, "follow_plan_day_id")
	f.AutoFollowID = field.NewInt64(table, "auto_follow_id")
	f.OpChangeFlag = field.NewInt64(table, "op_change_flag")
	f.CompanyID = field.NewInt64(table, "company_id")
	f.CompanyName = field.NewString(table, "company_name")
	f.UserID = field.NewInt64(table, "user_id")
	f.UserRealname = field.NewString(table, "user_realname")
	f.ResultJSON = field.NewString(table, "result_json")
	f.SpendTimeCode = field.NewString(table, "spend_time_code")
	f.SpendTimeName = field.NewString(table, "spend_time_name")
	f.ProcessingResultCode = field.NewString(table, "processing_result_code")
	f.ProcessingResultName = field.NewString(table, "processing_result_name")
	f.FollowUpTypeCode = field.NewString(table, "follow_up_type_code")
	f.FollowUpTypeName = field.NewString(table, "follow_up_type_name")
	f.Module = field.NewString(table, "module")

	f.fillFieldMap()

	return f
}

func (f *followDesc) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *followDesc) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 34)
	f.fieldMap["id"] = f.ID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
	f.fieldMap["follow_type"] = f.FollowType
	f.fieldMap["follow_type_code"] = f.FollowTypeCode
	f.fieldMap["follow_type_ico"] = f.FollowTypeIco
	f.fieldMap["follow_type_abbreviation"] = f.FollowTypeAbbreviation
	f.fieldMap["data_code"] = f.DataCode
	f.fieldMap["data_type"] = f.DataType
	f.fieldMap["data_id"] = f.DataID
	f.fieldMap["belong_op"] = f.BelongOp
	f.fieldMap["belong_op_name"] = f.BelongOpName
	f.fieldMap["belong_department_code"] = f.BelongDepartmentCode
	f.fieldMap["belong_department_name"] = f.BelongDepartmentName
	f.fieldMap["follow_title"] = f.FollowTitle
	f.fieldMap["follow_content"] = f.FollowContent
	f.fieldMap["follow_plan_day_flag"] = f.FollowPlanDayFlag
	f.fieldMap["follow_plan_day"] = f.FollowPlanDay
	f.fieldMap["follow_plan_day_id"] = f.FollowPlanDayID
	f.fieldMap["auto_follow_id"] = f.AutoFollowID
	f.fieldMap["op_change_flag"] = f.OpChangeFlag
	f.fieldMap["company_id"] = f.CompanyID
	f.fieldMap["company_name"] = f.CompanyName
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["user_realname"] = f.UserRealname
	f.fieldMap["result_json"] = f.ResultJSON
	f.fieldMap["spend_time_code"] = f.SpendTimeCode
	f.fieldMap["spend_time_name"] = f.SpendTimeName
	f.fieldMap["processing_result_code"] = f.ProcessingResultCode
	f.fieldMap["processing_result_name"] = f.ProcessingResultName
	f.fieldMap["follow_up_type_code"] = f.FollowUpTypeCode
	f.fieldMap["follow_up_type_name"] = f.FollowUpTypeName
	f.fieldMap["module"] = f.Module
}

func (f followDesc) clone(db *gorm.DB) followDesc {
	f.followDescDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f followDesc) replaceDB(db *gorm.DB) followDesc {
	f.followDescDo.ReplaceDB(db)
	return f
}

type followDescDo struct{ gen.DO }

func (f followDescDo) Debug() *followDescDo {
	return f.withDO(f.DO.Debug())
}

func (f followDescDo) WithContext(ctx context.Context) *followDescDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f followDescDo) ReadDB() *followDescDo {
	return f.Clauses(dbresolver.Read)
}

func (f followDescDo) WriteDB() *followDescDo {
	return f.Clauses(dbresolver.Write)
}

func (f followDescDo) Session(config *gorm.Session) *followDescDo {
	return f.withDO(f.DO.Session(config))
}

func (f followDescDo) Clauses(conds ...clause.Expression) *followDescDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f followDescDo) Returning(value interface{}, columns ...string) *followDescDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f followDescDo) Not(conds ...gen.Condition) *followDescDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f followDescDo) Or(conds ...gen.Condition) *followDescDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f followDescDo) Select(conds ...field.Expr) *followDescDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f followDescDo) Where(conds ...gen.Condition) *followDescDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f followDescDo) Order(conds ...field.Expr) *followDescDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f followDescDo) Distinct(cols ...field.Expr) *followDescDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f followDescDo) Omit(cols ...field.Expr) *followDescDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f followDescDo) Join(table schema.Tabler, on ...field.Expr) *followDescDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f followDescDo) LeftJoin(table schema.Tabler, on ...field.Expr) *followDescDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f followDescDo) RightJoin(table schema.Tabler, on ...field.Expr) *followDescDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f followDescDo) Group(cols ...field.Expr) *followDescDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f followDescDo) Having(conds ...gen.Condition) *followDescDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f followDescDo) Limit(limit int) *followDescDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f followDescDo) Offset(offset int) *followDescDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f followDescDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *followDescDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f followDescDo) Unscoped() *followDescDo {
	return f.withDO(f.DO.Unscoped())
}

func (f followDescDo) Create(values ...*model.FollowDesc) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f followDescDo) CreateInBatches(values []*model.FollowDesc, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f followDescDo) Save(values ...*model.FollowDesc) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f followDescDo) First() (*model.FollowDesc, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowDesc), nil
	}
}

func (f followDescDo) Take() (*model.FollowDesc, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowDesc), nil
	}
}

func (f followDescDo) Last() (*model.FollowDesc, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowDesc), nil
	}
}

func (f followDescDo) Find() ([]*model.FollowDesc, error) {
	result, err := f.DO.Find()
	return result.([]*model.FollowDesc), err
}

func (f followDescDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FollowDesc, err error) {
	buf := make([]*model.FollowDesc, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f followDescDo) FindInBatches(result *[]*model.FollowDesc, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f followDescDo) Attrs(attrs ...field.AssignExpr) *followDescDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f followDescDo) Assign(attrs ...field.AssignExpr) *followDescDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f followDescDo) Joins(fields ...field.RelationField) *followDescDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f followDescDo) Preload(fields ...field.RelationField) *followDescDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f followDescDo) FirstOrInit() (*model.FollowDesc, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowDesc), nil
	}
}

func (f followDescDo) FirstOrCreate() (*model.FollowDesc, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowDesc), nil
	}
}

func (f followDescDo) FindByPage(offset int, limit int) (result []*model.FollowDesc, count int64, err error) {
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

func (f followDescDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f followDescDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f followDescDo) Delete(models ...*model.FollowDesc) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *followDescDo) withDO(do gen.Dao) *followDescDo {
	f.DO = *do.(*gen.DO)
	return f
}