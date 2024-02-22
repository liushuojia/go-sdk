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

func newConsultation(db *gorm.DB, opts ...gen.DOOption) consultation {
	_consultation := consultation{}

	_consultation.consultationDo.UseDB(db, opts...)
	_consultation.consultationDo.UseModel(&model.Consultation{})

	tableName := _consultation.consultationDo.TableName()
	_consultation.ALL = field.NewAsterisk(tableName)
	_consultation.ID = field.NewInt64(tableName, "id")
	_consultation.CreatedAt = field.NewTime(tableName, "created_at")
	_consultation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_consultation.DeletedAt = field.NewField(tableName, "deleted_at")
	_consultation.UID = field.NewInt64(tableName, "uid")
	_consultation.UIDName = field.NewString(tableName, "uid_name")
	_consultation.UIDSex = field.NewInt64(tableName, "uid_sex")
	_consultation.Mobile = field.NewString(tableName, "mobile")
	_consultation.Email = field.NewString(tableName, "email")
	_consultation.Weixin = field.NewString(tableName, "weixin")
	_consultation.Aliwangwang = field.NewString(tableName, "aliwangwang")
	_consultation.MafengwoUID = field.NewString(tableName, "mafengwo_uid")
	_consultation.MafengwoIm = field.NewString(tableName, "mafengwo_im")
	_consultation.CtripAccount = field.NewString(tableName, "ctrip_account")
	_consultation.CtripGrab = field.NewString(tableName, "ctrip_grab")
	_consultation.QunaerPhone = field.NewString(tableName, "qunaer_phone")
	_consultation.DepartureCityCode = field.NewString(tableName, "departure_city_code")
	_consultation.DepartureCity = field.NewString(tableName, "departure_city")
	_consultation.ArrivalCityCode = field.NewString(tableName, "arrival_city_code")
	_consultation.ArrivalCity = field.NewString(tableName, "arrival_city")
	_consultation.AdultNum = field.NewInt64(tableName, "adult_num")
	_consultation.ChildrenNum = field.NewInt64(tableName, "children_num")
	_consultation.TravelDate = field.NewInt64(tableName, "travel_date")
	_consultation.TravelDateEnd = field.NewInt64(tableName, "travel_date_end")
	_consultation.TravelDayNum = field.NewInt64(tableName, "travel_day_num")
	_consultation.ReturnDate = field.NewInt64(tableName, "return_date")
	_consultation.ReturnDateEnd = field.NewInt64(tableName, "return_date_end")
	_consultation.Budget = field.NewString(tableName, "budget")
	_consultation.BudgetCode = field.NewString(tableName, "budget_code")
	_consultation.DemandNumber = field.NewString(tableName, "demand_number")
	_consultation.TransferNumber = field.NewString(tableName, "transfer_number")
	_consultation.SourceCode = field.NewString(tableName, "source_code")
	_consultation.SourceName = field.NewString(tableName, "source_name")
	_consultation.SourceAbbreviation = field.NewString(tableName, "source_abbreviation")
	_consultation.MethodCode = field.NewString(tableName, "method_code")
	_consultation.MethodName = field.NewString(tableName, "method_name")
	_consultation.FollowStatusCode = field.NewString(tableName, "follow_status_code")
	_consultation.FollowStatusName = field.NewString(tableName, "follow_status_name")
	_consultation.ProductTypeCode = field.NewString(tableName, "product_type_code")
	_consultation.ProductType = field.NewString(tableName, "product_type")
	_consultation.ProductTypeAbbreviation = field.NewString(tableName, "product_type_abbreviation")
	_consultation.CreateOp = field.NewInt64(tableName, "create_op")
	_consultation.CreateOpName = field.NewString(tableName, "create_op_name")
	_consultation.BelongOp = field.NewInt64(tableName, "belong_op")
	_consultation.BelongOpName = field.NewString(tableName, "belong_op_name")
	_consultation.BelongDepartmentCode = field.NewString(tableName, "belong_department_code")
	_consultation.BelongDepartmentName = field.NewString(tableName, "belong_department_name")
	_consultation.FollowPlanFlag = field.NewInt64(tableName, "follow_plan_flag")
	_consultation.FollowPlanDay = field.NewInt64(tableName, "follow_plan_day")
	_consultation.Status = field.NewInt64(tableName, "status")
	_consultation.StatusOrder = field.NewInt64(tableName, "status_order")
	_consultation.BecomeOrderAt = field.NewInt64(tableName, "become_order_at")
	_consultation.BecomeOrderSecond = field.NewInt64(tableName, "become_order_second")
	_consultation.StatusClose = field.NewInt64(tableName, "status_close")
	_consultation.CloseReason = field.NewString(tableName, "close_reason")
	_consultation.CloseReasonCode = field.NewString(tableName, "close_reason_code")
	_consultation.LastFollowTime = field.NewInt64(tableName, "last_follow_time")
	_consultation.SystemAutoFlag = field.NewInt64(tableName, "system_auto_flag")
	_consultation.SystemAutoTime = field.NewInt64(tableName, "system_auto_time")

	_consultation.fillFieldMap()

	return _consultation
}

type consultation struct {
	consultationDo

	ALL                     field.Asterisk
	ID                      field.Int64  // 自动编号
	CreatedAt               field.Time   // 创建时间
	UpdatedAt               field.Time   // 更新时间
	DeletedAt               field.Field  // 删除时间
	UID                     field.Int64  // 用户编码
	UIDName                 field.String // 用户名字
	UIDSex                  field.Int64  // 用户性别
	Mobile                  field.String // 手机号码
	Email                   field.String // 邮箱
	Weixin                  field.String // 微信号
	Aliwangwang             field.String // 阿里旺旺
	MafengwoUID             field.String // 马蜂窝UID
	MafengwoIm              field.String // 马蜂窝IM
	CtripAccount            field.String // 携程子账号
	CtripGrab               field.String // 携程抢单号
	QunaerPhone             field.String // 去哪儿电话
	DepartureCityCode       field.String // 出发城市编码
	DepartureCity           field.String // 出发城市名称
	ArrivalCityCode         field.String // 目的地编码
	ArrivalCity             field.String // 目的地名称
	AdultNum                field.Int64  // 成人数
	ChildrenNum             field.Int64  // 儿童数
	TravelDate              field.Int64  // 出行日期
	TravelDateEnd           field.Int64  // 出行日期截止
	TravelDayNum            field.Int64  // 出行天数
	ReturnDate              field.Int64  // 返程日期
	ReturnDateEnd           field.Int64  // 返程日期截止
	Budget                  field.String // 预算
	BudgetCode              field.String // 预算编码
	DemandNumber            field.String // 需求单号
	TransferNumber          field.String // 转接号
	SourceCode              field.String // 系统来源编码
	SourceName              field.String // 系统来源名称
	SourceAbbreviation      field.String // 系统来源简称
	MethodCode              field.String // 咨询方式编码
	MethodName              field.String // 咨询方式名称
	FollowStatusCode        field.String // 咨询跟进标识编码
	FollowStatusName        field.String // 咨询跟进标识名称
	ProductTypeCode         field.String // 产品类型编码
	ProductType             field.String // 产品类型
	ProductTypeAbbreviation field.String // 产品类型简称
	CreateOp                field.Int64  // 创建opID
	CreateOpName            field.String // 创建op姓名
	BelongOp                field.Int64  // 所属op
	BelongOpName            field.String // 所属op姓名
	BelongDepartmentCode    field.String // 所属部门编码
	BelongDepartmentName    field.String // 所属部门名称
	FollowPlanFlag          field.Int64  // 是否远期计划
	FollowPlanDay           field.Int64  // 远期计划日期
	Status                  field.Int64  // 状态 0 跟进中 1 历史咨询 2 远期计划
	StatusOrder             field.Int64  // 成单状态 0 未成单 1 已成单
	BecomeOrderAt           field.Int64  // 成单时间
	BecomeOrderSecond       field.Int64  // 成单时长 单位秒
	StatusClose             field.Int64  // 关闭状态	0 未关闭	1 已关闭	2 关闭中
	CloseReason             field.String // 关闭原因
	CloseReasonCode         field.String // 关闭原因编码
	LastFollowTime          field.Int64  // 最后处理时间
	SystemAutoFlag          field.Int64  // 系统自动分配
	SystemAutoTime          field.Int64  // 咨询分配op时间

	fieldMap map[string]field.Expr
}

func (c consultation) Table(newTableName string) *consultation {
	c.consultationDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c consultation) As(alias string) *consultation {
	c.consultationDo.DO = *(c.consultationDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *consultation) updateTableName(table string) *consultation {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.UID = field.NewInt64(table, "uid")
	c.UIDName = field.NewString(table, "uid_name")
	c.UIDSex = field.NewInt64(table, "uid_sex")
	c.Mobile = field.NewString(table, "mobile")
	c.Email = field.NewString(table, "email")
	c.Weixin = field.NewString(table, "weixin")
	c.Aliwangwang = field.NewString(table, "aliwangwang")
	c.MafengwoUID = field.NewString(table, "mafengwo_uid")
	c.MafengwoIm = field.NewString(table, "mafengwo_im")
	c.CtripAccount = field.NewString(table, "ctrip_account")
	c.CtripGrab = field.NewString(table, "ctrip_grab")
	c.QunaerPhone = field.NewString(table, "qunaer_phone")
	c.DepartureCityCode = field.NewString(table, "departure_city_code")
	c.DepartureCity = field.NewString(table, "departure_city")
	c.ArrivalCityCode = field.NewString(table, "arrival_city_code")
	c.ArrivalCity = field.NewString(table, "arrival_city")
	c.AdultNum = field.NewInt64(table, "adult_num")
	c.ChildrenNum = field.NewInt64(table, "children_num")
	c.TravelDate = field.NewInt64(table, "travel_date")
	c.TravelDateEnd = field.NewInt64(table, "travel_date_end")
	c.TravelDayNum = field.NewInt64(table, "travel_day_num")
	c.ReturnDate = field.NewInt64(table, "return_date")
	c.ReturnDateEnd = field.NewInt64(table, "return_date_end")
	c.Budget = field.NewString(table, "budget")
	c.BudgetCode = field.NewString(table, "budget_code")
	c.DemandNumber = field.NewString(table, "demand_number")
	c.TransferNumber = field.NewString(table, "transfer_number")
	c.SourceCode = field.NewString(table, "source_code")
	c.SourceName = field.NewString(table, "source_name")
	c.SourceAbbreviation = field.NewString(table, "source_abbreviation")
	c.MethodCode = field.NewString(table, "method_code")
	c.MethodName = field.NewString(table, "method_name")
	c.FollowStatusCode = field.NewString(table, "follow_status_code")
	c.FollowStatusName = field.NewString(table, "follow_status_name")
	c.ProductTypeCode = field.NewString(table, "product_type_code")
	c.ProductType = field.NewString(table, "product_type")
	c.ProductTypeAbbreviation = field.NewString(table, "product_type_abbreviation")
	c.CreateOp = field.NewInt64(table, "create_op")
	c.CreateOpName = field.NewString(table, "create_op_name")
	c.BelongOp = field.NewInt64(table, "belong_op")
	c.BelongOpName = field.NewString(table, "belong_op_name")
	c.BelongDepartmentCode = field.NewString(table, "belong_department_code")
	c.BelongDepartmentName = field.NewString(table, "belong_department_name")
	c.FollowPlanFlag = field.NewInt64(table, "follow_plan_flag")
	c.FollowPlanDay = field.NewInt64(table, "follow_plan_day")
	c.Status = field.NewInt64(table, "status")
	c.StatusOrder = field.NewInt64(table, "status_order")
	c.BecomeOrderAt = field.NewInt64(table, "become_order_at")
	c.BecomeOrderSecond = field.NewInt64(table, "become_order_second")
	c.StatusClose = field.NewInt64(table, "status_close")
	c.CloseReason = field.NewString(table, "close_reason")
	c.CloseReasonCode = field.NewString(table, "close_reason_code")
	c.LastFollowTime = field.NewInt64(table, "last_follow_time")
	c.SystemAutoFlag = field.NewInt64(table, "system_auto_flag")
	c.SystemAutoTime = field.NewInt64(table, "system_auto_time")

	c.fillFieldMap()

	return c
}

func (c *consultation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *consultation) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 59)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["uid"] = c.UID
	c.fieldMap["uid_name"] = c.UIDName
	c.fieldMap["uid_sex"] = c.UIDSex
	c.fieldMap["mobile"] = c.Mobile
	c.fieldMap["email"] = c.Email
	c.fieldMap["weixin"] = c.Weixin
	c.fieldMap["aliwangwang"] = c.Aliwangwang
	c.fieldMap["mafengwo_uid"] = c.MafengwoUID
	c.fieldMap["mafengwo_im"] = c.MafengwoIm
	c.fieldMap["ctrip_account"] = c.CtripAccount
	c.fieldMap["ctrip_grab"] = c.CtripGrab
	c.fieldMap["qunaer_phone"] = c.QunaerPhone
	c.fieldMap["departure_city_code"] = c.DepartureCityCode
	c.fieldMap["departure_city"] = c.DepartureCity
	c.fieldMap["arrival_city_code"] = c.ArrivalCityCode
	c.fieldMap["arrival_city"] = c.ArrivalCity
	c.fieldMap["adult_num"] = c.AdultNum
	c.fieldMap["children_num"] = c.ChildrenNum
	c.fieldMap["travel_date"] = c.TravelDate
	c.fieldMap["travel_date_end"] = c.TravelDateEnd
	c.fieldMap["travel_day_num"] = c.TravelDayNum
	c.fieldMap["return_date"] = c.ReturnDate
	c.fieldMap["return_date_end"] = c.ReturnDateEnd
	c.fieldMap["budget"] = c.Budget
	c.fieldMap["budget_code"] = c.BudgetCode
	c.fieldMap["demand_number"] = c.DemandNumber
	c.fieldMap["transfer_number"] = c.TransferNumber
	c.fieldMap["source_code"] = c.SourceCode
	c.fieldMap["source_name"] = c.SourceName
	c.fieldMap["source_abbreviation"] = c.SourceAbbreviation
	c.fieldMap["method_code"] = c.MethodCode
	c.fieldMap["method_name"] = c.MethodName
	c.fieldMap["follow_status_code"] = c.FollowStatusCode
	c.fieldMap["follow_status_name"] = c.FollowStatusName
	c.fieldMap["product_type_code"] = c.ProductTypeCode
	c.fieldMap["product_type"] = c.ProductType
	c.fieldMap["product_type_abbreviation"] = c.ProductTypeAbbreviation
	c.fieldMap["create_op"] = c.CreateOp
	c.fieldMap["create_op_name"] = c.CreateOpName
	c.fieldMap["belong_op"] = c.BelongOp
	c.fieldMap["belong_op_name"] = c.BelongOpName
	c.fieldMap["belong_department_code"] = c.BelongDepartmentCode
	c.fieldMap["belong_department_name"] = c.BelongDepartmentName
	c.fieldMap["follow_plan_flag"] = c.FollowPlanFlag
	c.fieldMap["follow_plan_day"] = c.FollowPlanDay
	c.fieldMap["status"] = c.Status
	c.fieldMap["status_order"] = c.StatusOrder
	c.fieldMap["become_order_at"] = c.BecomeOrderAt
	c.fieldMap["become_order_second"] = c.BecomeOrderSecond
	c.fieldMap["status_close"] = c.StatusClose
	c.fieldMap["close_reason"] = c.CloseReason
	c.fieldMap["close_reason_code"] = c.CloseReasonCode
	c.fieldMap["last_follow_time"] = c.LastFollowTime
	c.fieldMap["system_auto_flag"] = c.SystemAutoFlag
	c.fieldMap["system_auto_time"] = c.SystemAutoTime
}

func (c consultation) clone(db *gorm.DB) consultation {
	c.consultationDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c consultation) replaceDB(db *gorm.DB) consultation {
	c.consultationDo.ReplaceDB(db)
	return c
}

type consultationDo struct{ gen.DO }

func (c consultationDo) Debug() *consultationDo {
	return c.withDO(c.DO.Debug())
}

func (c consultationDo) WithContext(ctx context.Context) *consultationDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c consultationDo) ReadDB() *consultationDo {
	return c.Clauses(dbresolver.Read)
}

func (c consultationDo) WriteDB() *consultationDo {
	return c.Clauses(dbresolver.Write)
}

func (c consultationDo) Session(config *gorm.Session) *consultationDo {
	return c.withDO(c.DO.Session(config))
}

func (c consultationDo) Clauses(conds ...clause.Expression) *consultationDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c consultationDo) Returning(value interface{}, columns ...string) *consultationDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c consultationDo) Not(conds ...gen.Condition) *consultationDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c consultationDo) Or(conds ...gen.Condition) *consultationDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c consultationDo) Select(conds ...field.Expr) *consultationDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c consultationDo) Where(conds ...gen.Condition) *consultationDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c consultationDo) Order(conds ...field.Expr) *consultationDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c consultationDo) Distinct(cols ...field.Expr) *consultationDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c consultationDo) Omit(cols ...field.Expr) *consultationDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c consultationDo) Join(table schema.Tabler, on ...field.Expr) *consultationDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c consultationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *consultationDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c consultationDo) RightJoin(table schema.Tabler, on ...field.Expr) *consultationDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c consultationDo) Group(cols ...field.Expr) *consultationDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c consultationDo) Having(conds ...gen.Condition) *consultationDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c consultationDo) Limit(limit int) *consultationDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c consultationDo) Offset(offset int) *consultationDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c consultationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *consultationDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c consultationDo) Unscoped() *consultationDo {
	return c.withDO(c.DO.Unscoped())
}

func (c consultationDo) Create(values ...*model.Consultation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c consultationDo) CreateInBatches(values []*model.Consultation, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c consultationDo) Save(values ...*model.Consultation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c consultationDo) First() (*model.Consultation, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Consultation), nil
	}
}

func (c consultationDo) Take() (*model.Consultation, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Consultation), nil
	}
}

func (c consultationDo) Last() (*model.Consultation, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Consultation), nil
	}
}

func (c consultationDo) Find() ([]*model.Consultation, error) {
	result, err := c.DO.Find()
	return result.([]*model.Consultation), err
}

func (c consultationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Consultation, err error) {
	buf := make([]*model.Consultation, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c consultationDo) FindInBatches(result *[]*model.Consultation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c consultationDo) Attrs(attrs ...field.AssignExpr) *consultationDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c consultationDo) Assign(attrs ...field.AssignExpr) *consultationDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c consultationDo) Joins(fields ...field.RelationField) *consultationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c consultationDo) Preload(fields ...field.RelationField) *consultationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c consultationDo) FirstOrInit() (*model.Consultation, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Consultation), nil
	}
}

func (c consultationDo) FirstOrCreate() (*model.Consultation, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Consultation), nil
	}
}

func (c consultationDo) FindByPage(offset int, limit int) (result []*model.Consultation, count int64, err error) {
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

func (c consultationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c consultationDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c consultationDo) Delete(models ...*model.Consultation) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *consultationDo) withDO(do gen.Dao) *consultationDo {
	c.DO = *do.(*gen.DO)
	return c
}
