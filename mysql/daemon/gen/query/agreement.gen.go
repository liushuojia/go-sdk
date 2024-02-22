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

func newAgreement(db *gorm.DB, opts ...gen.DOOption) agreement {
	_agreement := agreement{}

	_agreement.agreementDo.UseDB(db, opts...)
	_agreement.agreementDo.UseModel(&model.Agreement{})

	tableName := _agreement.agreementDo.TableName()
	_agreement.ALL = field.NewAsterisk(tableName)
	_agreement.ID = field.NewInt64(tableName, "id")
	_agreement.CreatedAt = field.NewTime(tableName, "created_at")
	_agreement.UpdatedAt = field.NewTime(tableName, "updated_at")
	_agreement.DeletedAt = field.NewField(tableName, "deleted_at")
	_agreement.DataID = field.NewInt64(tableName, "data_id")
	_agreement.DataType = field.NewInt64(tableName, "data_type")
	_agreement.DataCode = field.NewString(tableName, "data_code")
	_agreement.DataName = field.NewString(tableName, "data_name")
	_agreement.AgreementType = field.NewInt64(tableName, "agreement_type")
	_agreement.SettlementName = field.NewString(tableName, "settlement_name")
	_agreement.SettlementCode = field.NewString(tableName, "settlement_code")
	_agreement.ArrivalCity = field.NewString(tableName, "arrival_city")
	_agreement.ArrivalCityCode = field.NewString(tableName, "arrival_city_code")
	_agreement.AgreementNum = field.NewString(tableName, "agreement_num")
	_agreement.AgreementTitle = field.NewString(tableName, "agreement_title")
	_agreement.AgreementDesc = field.NewString(tableName, "agreement_desc")
	_agreement.AgreementStatus = field.NewInt64(tableName, "agreement_status")
	_agreement.FileSecrecyFlag = field.NewInt64(tableName, "file_secrecy_flag")
	_agreement.ValidityStart = field.NewInt64(tableName, "validity_start")
	_agreement.ValidityEnd = field.NewInt64(tableName, "validity_end")
	_agreement.UID = field.NewInt64(tableName, "uid")
	_agreement.Realname = field.NewString(tableName, "realname")
	_agreement.Mobile = field.NewString(tableName, "mobile")
	_agreement.Bid = field.NewInt64(tableName, "bid")
	_agreement.BankType = field.NewString(tableName, "bank_type")
	_agreement.BankName = field.NewString(tableName, "bank_name")
	_agreement.BankNum = field.NewString(tableName, "bank_num")
	_agreement.BankAccount = field.NewString(tableName, "bank_account")
	_agreement.OurCompanyID = field.NewInt64(tableName, "our_company_id")
	_agreement.OurCompanyName = field.NewString(tableName, "our_company_name")
	_agreement.OurBankID = field.NewInt64(tableName, "our_bank_id")
	_agreement.OurOpID = field.NewInt64(tableName, "our_op_id")
	_agreement.OurOpName = field.NewString(tableName, "our_op_name")
	_agreement.BelongOp = field.NewInt64(tableName, "belong_op")
	_agreement.BelongOpName = field.NewString(tableName, "belong_op_name")
	_agreement.BelongDepartmentCode = field.NewString(tableName, "belong_department_code")
	_agreement.BelongDepartmentName = field.NewString(tableName, "belong_department_name")

	_agreement.fillFieldMap()

	return _agreement
}

type agreement struct {
	agreementDo

	ALL                  field.Asterisk
	ID                   field.Int64  // 自动编号
	CreatedAt            field.Time   // 创建时间
	UpdatedAt            field.Time   // 更新时间
	DeletedAt            field.Field  // 删除时间
	DataID               field.Int64  // 业务所属自增id
	DataType             field.Int64  // 业务所属类型id
	DataCode             field.String // 业务所属编码
	DataName             field.String // 公司名称
	AgreementType        field.Int64  // 协议类型 0 供应商协议  1 客户协议
	SettlementName       field.String // 结算方式
	SettlementCode       field.String // 结算方式编码
	ArrivalCity          field.String // 目的地名称
	ArrivalCityCode      field.String // 目的地编码
	AgreementNum         field.String // 协议编号 唯一
	AgreementTitle       field.String // 协议名称
	AgreementDesc        field.String // 协议描述
	AgreementStatus      field.Int64  // 状态 -1 保存中 0 走审核流程中  1有效 2 过期(归档)
	FileSecrecyFlag      field.Int64  // 协议附件是否保密
	ValidityStart        field.Int64  // 有效 开始日期
	ValidityEnd          field.Int64  // 有效 结束日期
	UID                  field.Int64  // 协议联系人id
	Realname             field.String // 联系人
	Mobile               field.String // 联系电话
	Bid                  field.Int64  // 协议银行id
	BankType             field.String // 银行类型
	BankName             field.String // 开户行
	BankNum              field.String // 银行账号
	BankAccount          field.String // 账户名
	OurCompanyID         field.Int64  // 签约公司id
	OurCompanyName       field.String // 签约公司名称
	OurBankID            field.Int64  // 签约主体银行id
	OurOpID              field.Int64  // 签约人员 id
	OurOpName            field.String // 签约人员名称
	BelongOp             field.Int64  // 所属人OP
	BelongOpName         field.String // 所属人名称
	BelongDepartmentCode field.String // 所属部门编码
	BelongDepartmentName field.String // 所属部门名称

	fieldMap map[string]field.Expr
}

func (a agreement) Table(newTableName string) *agreement {
	a.agreementDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a agreement) As(alias string) *agreement {
	a.agreementDo.DO = *(a.agreementDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *agreement) updateTableName(table string) *agreement {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.DataID = field.NewInt64(table, "data_id")
	a.DataType = field.NewInt64(table, "data_type")
	a.DataCode = field.NewString(table, "data_code")
	a.DataName = field.NewString(table, "data_name")
	a.AgreementType = field.NewInt64(table, "agreement_type")
	a.SettlementName = field.NewString(table, "settlement_name")
	a.SettlementCode = field.NewString(table, "settlement_code")
	a.ArrivalCity = field.NewString(table, "arrival_city")
	a.ArrivalCityCode = field.NewString(table, "arrival_city_code")
	a.AgreementNum = field.NewString(table, "agreement_num")
	a.AgreementTitle = field.NewString(table, "agreement_title")
	a.AgreementDesc = field.NewString(table, "agreement_desc")
	a.AgreementStatus = field.NewInt64(table, "agreement_status")
	a.FileSecrecyFlag = field.NewInt64(table, "file_secrecy_flag")
	a.ValidityStart = field.NewInt64(table, "validity_start")
	a.ValidityEnd = field.NewInt64(table, "validity_end")
	a.UID = field.NewInt64(table, "uid")
	a.Realname = field.NewString(table, "realname")
	a.Mobile = field.NewString(table, "mobile")
	a.Bid = field.NewInt64(table, "bid")
	a.BankType = field.NewString(table, "bank_type")
	a.BankName = field.NewString(table, "bank_name")
	a.BankNum = field.NewString(table, "bank_num")
	a.BankAccount = field.NewString(table, "bank_account")
	a.OurCompanyID = field.NewInt64(table, "our_company_id")
	a.OurCompanyName = field.NewString(table, "our_company_name")
	a.OurBankID = field.NewInt64(table, "our_bank_id")
	a.OurOpID = field.NewInt64(table, "our_op_id")
	a.OurOpName = field.NewString(table, "our_op_name")
	a.BelongOp = field.NewInt64(table, "belong_op")
	a.BelongOpName = field.NewString(table, "belong_op_name")
	a.BelongDepartmentCode = field.NewString(table, "belong_department_code")
	a.BelongDepartmentName = field.NewString(table, "belong_department_name")

	a.fillFieldMap()

	return a
}

func (a *agreement) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *agreement) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 37)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["data_id"] = a.DataID
	a.fieldMap["data_type"] = a.DataType
	a.fieldMap["data_code"] = a.DataCode
	a.fieldMap["data_name"] = a.DataName
	a.fieldMap["agreement_type"] = a.AgreementType
	a.fieldMap["settlement_name"] = a.SettlementName
	a.fieldMap["settlement_code"] = a.SettlementCode
	a.fieldMap["arrival_city"] = a.ArrivalCity
	a.fieldMap["arrival_city_code"] = a.ArrivalCityCode
	a.fieldMap["agreement_num"] = a.AgreementNum
	a.fieldMap["agreement_title"] = a.AgreementTitle
	a.fieldMap["agreement_desc"] = a.AgreementDesc
	a.fieldMap["agreement_status"] = a.AgreementStatus
	a.fieldMap["file_secrecy_flag"] = a.FileSecrecyFlag
	a.fieldMap["validity_start"] = a.ValidityStart
	a.fieldMap["validity_end"] = a.ValidityEnd
	a.fieldMap["uid"] = a.UID
	a.fieldMap["realname"] = a.Realname
	a.fieldMap["mobile"] = a.Mobile
	a.fieldMap["bid"] = a.Bid
	a.fieldMap["bank_type"] = a.BankType
	a.fieldMap["bank_name"] = a.BankName
	a.fieldMap["bank_num"] = a.BankNum
	a.fieldMap["bank_account"] = a.BankAccount
	a.fieldMap["our_company_id"] = a.OurCompanyID
	a.fieldMap["our_company_name"] = a.OurCompanyName
	a.fieldMap["our_bank_id"] = a.OurBankID
	a.fieldMap["our_op_id"] = a.OurOpID
	a.fieldMap["our_op_name"] = a.OurOpName
	a.fieldMap["belong_op"] = a.BelongOp
	a.fieldMap["belong_op_name"] = a.BelongOpName
	a.fieldMap["belong_department_code"] = a.BelongDepartmentCode
	a.fieldMap["belong_department_name"] = a.BelongDepartmentName
}

func (a agreement) clone(db *gorm.DB) agreement {
	a.agreementDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a agreement) replaceDB(db *gorm.DB) agreement {
	a.agreementDo.ReplaceDB(db)
	return a
}

type agreementDo struct{ gen.DO }

func (a agreementDo) Debug() *agreementDo {
	return a.withDO(a.DO.Debug())
}

func (a agreementDo) WithContext(ctx context.Context) *agreementDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a agreementDo) ReadDB() *agreementDo {
	return a.Clauses(dbresolver.Read)
}

func (a agreementDo) WriteDB() *agreementDo {
	return a.Clauses(dbresolver.Write)
}

func (a agreementDo) Session(config *gorm.Session) *agreementDo {
	return a.withDO(a.DO.Session(config))
}

func (a agreementDo) Clauses(conds ...clause.Expression) *agreementDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a agreementDo) Returning(value interface{}, columns ...string) *agreementDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a agreementDo) Not(conds ...gen.Condition) *agreementDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a agreementDo) Or(conds ...gen.Condition) *agreementDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a agreementDo) Select(conds ...field.Expr) *agreementDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a agreementDo) Where(conds ...gen.Condition) *agreementDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a agreementDo) Order(conds ...field.Expr) *agreementDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a agreementDo) Distinct(cols ...field.Expr) *agreementDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a agreementDo) Omit(cols ...field.Expr) *agreementDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a agreementDo) Join(table schema.Tabler, on ...field.Expr) *agreementDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a agreementDo) LeftJoin(table schema.Tabler, on ...field.Expr) *agreementDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a agreementDo) RightJoin(table schema.Tabler, on ...field.Expr) *agreementDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a agreementDo) Group(cols ...field.Expr) *agreementDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a agreementDo) Having(conds ...gen.Condition) *agreementDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a agreementDo) Limit(limit int) *agreementDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a agreementDo) Offset(offset int) *agreementDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a agreementDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *agreementDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a agreementDo) Unscoped() *agreementDo {
	return a.withDO(a.DO.Unscoped())
}

func (a agreementDo) Create(values ...*model.Agreement) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a agreementDo) CreateInBatches(values []*model.Agreement, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a agreementDo) Save(values ...*model.Agreement) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a agreementDo) First() (*model.Agreement, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Agreement), nil
	}
}

func (a agreementDo) Take() (*model.Agreement, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Agreement), nil
	}
}

func (a agreementDo) Last() (*model.Agreement, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Agreement), nil
	}
}

func (a agreementDo) Find() ([]*model.Agreement, error) {
	result, err := a.DO.Find()
	return result.([]*model.Agreement), err
}

func (a agreementDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Agreement, err error) {
	buf := make([]*model.Agreement, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a agreementDo) FindInBatches(result *[]*model.Agreement, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a agreementDo) Attrs(attrs ...field.AssignExpr) *agreementDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a agreementDo) Assign(attrs ...field.AssignExpr) *agreementDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a agreementDo) Joins(fields ...field.RelationField) *agreementDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a agreementDo) Preload(fields ...field.RelationField) *agreementDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a agreementDo) FirstOrInit() (*model.Agreement, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Agreement), nil
	}
}

func (a agreementDo) FirstOrCreate() (*model.Agreement, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Agreement), nil
	}
}

func (a agreementDo) FindByPage(offset int, limit int) (result []*model.Agreement, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a agreementDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a agreementDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a agreementDo) Delete(models ...*model.Agreement) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *agreementDo) withDO(do gen.Dao) *agreementDo {
	a.DO = *do.(*gen.DO)
	return a
}