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

func newWeixinAccount(db *gorm.DB, opts ...gen.DOOption) weixinAccount {
	_weixinAccount := weixinAccount{}

	_weixinAccount.weixinAccountDo.UseDB(db, opts...)
	_weixinAccount.weixinAccountDo.UseModel(&model.WeixinAccount{})

	tableName := _weixinAccount.weixinAccountDo.TableName()
	_weixinAccount.ALL = field.NewAsterisk(tableName)
	_weixinAccount.ID = field.NewInt64(tableName, "id")
	_weixinAccount.CreatedAt = field.NewTime(tableName, "created_at")
	_weixinAccount.UpdatedAt = field.NewTime(tableName, "updated_at")
	_weixinAccount.DeletedAt = field.NewField(tableName, "deleted_at")
	_weixinAccount.AppID = field.NewString(tableName, "app_id")
	_weixinAccount.AppSecret = field.NewString(tableName, "app_secret")
	_weixinAccount.TokenWeb = field.NewString(tableName, "token_web")
	_weixinAccount.EncodingAesKey = field.NewString(tableName, "encoding_aes_key")
	_weixinAccount.EncodingType = field.NewInt64(tableName, "encoding_type")
	_weixinAccount.WeixinTitle = field.NewString(tableName, "weixin_title")
	_weixinAccount.WeixinToken = field.NewString(tableName, "weixin_token")
	_weixinAccount.TokenExp = field.NewInt64(tableName, "token_exp")
	_weixinAccount.WeixinUserID = field.NewString(tableName, "weixin_user_id")
	_weixinAccount.JsapiTicket = field.NewString(tableName, "jsapi_ticket")
	_weixinAccount.JsapiTicketExp = field.NewInt64(tableName, "jsapi_ticket_exp")
	_weixinAccount.WeixinImg = field.NewString(tableName, "weixin_img")
	_weixinAccount.LoginFlag = field.NewInt64(tableName, "login_flag")

	_weixinAccount.fillFieldMap()

	return _weixinAccount
}

type weixinAccount struct {
	weixinAccountDo

	ALL            field.Asterisk
	ID             field.Int64  // 自动编号
	CreatedAt      field.Time   // 创建时间
	UpdatedAt      field.Time   // 更新时间
	DeletedAt      field.Field  // 删除时间
	AppID          field.String // 公众号appid
	AppSecret      field.String // 公众号密钥
	TokenWeb       field.String // 被动token
	EncodingAesKey field.String // 被动密钥
	EncodingType   field.Int64  // 加密方式
	WeixinTitle    field.String // 名称
	WeixinToken    field.String // 主动token
	TokenExp       field.Int64  // 主动token有效期
	WeixinUserID   field.String // 微信原始名称
	JsapiTicket    field.String // js主动密钥
	JsapiTicketExp field.Int64  // js主动密钥有效期
	WeixinImg      field.String // 微信图片
	LoginFlag      field.Int64  // 默认登录

	fieldMap map[string]field.Expr
}

func (w weixinAccount) Table(newTableName string) *weixinAccount {
	w.weixinAccountDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w weixinAccount) As(alias string) *weixinAccount {
	w.weixinAccountDo.DO = *(w.weixinAccountDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *weixinAccount) updateTableName(table string) *weixinAccount {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.AppID = field.NewString(table, "app_id")
	w.AppSecret = field.NewString(table, "app_secret")
	w.TokenWeb = field.NewString(table, "token_web")
	w.EncodingAesKey = field.NewString(table, "encoding_aes_key")
	w.EncodingType = field.NewInt64(table, "encoding_type")
	w.WeixinTitle = field.NewString(table, "weixin_title")
	w.WeixinToken = field.NewString(table, "weixin_token")
	w.TokenExp = field.NewInt64(table, "token_exp")
	w.WeixinUserID = field.NewString(table, "weixin_user_id")
	w.JsapiTicket = field.NewString(table, "jsapi_ticket")
	w.JsapiTicketExp = field.NewInt64(table, "jsapi_ticket_exp")
	w.WeixinImg = field.NewString(table, "weixin_img")
	w.LoginFlag = field.NewInt64(table, "login_flag")

	w.fillFieldMap()

	return w
}

func (w *weixinAccount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *weixinAccount) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 17)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["app_id"] = w.AppID
	w.fieldMap["app_secret"] = w.AppSecret
	w.fieldMap["token_web"] = w.TokenWeb
	w.fieldMap["encoding_aes_key"] = w.EncodingAesKey
	w.fieldMap["encoding_type"] = w.EncodingType
	w.fieldMap["weixin_title"] = w.WeixinTitle
	w.fieldMap["weixin_token"] = w.WeixinToken
	w.fieldMap["token_exp"] = w.TokenExp
	w.fieldMap["weixin_user_id"] = w.WeixinUserID
	w.fieldMap["jsapi_ticket"] = w.JsapiTicket
	w.fieldMap["jsapi_ticket_exp"] = w.JsapiTicketExp
	w.fieldMap["weixin_img"] = w.WeixinImg
	w.fieldMap["login_flag"] = w.LoginFlag
}

func (w weixinAccount) clone(db *gorm.DB) weixinAccount {
	w.weixinAccountDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w weixinAccount) replaceDB(db *gorm.DB) weixinAccount {
	w.weixinAccountDo.ReplaceDB(db)
	return w
}

type weixinAccountDo struct{ gen.DO }

func (w weixinAccountDo) Debug() *weixinAccountDo {
	return w.withDO(w.DO.Debug())
}

func (w weixinAccountDo) WithContext(ctx context.Context) *weixinAccountDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w weixinAccountDo) ReadDB() *weixinAccountDo {
	return w.Clauses(dbresolver.Read)
}

func (w weixinAccountDo) WriteDB() *weixinAccountDo {
	return w.Clauses(dbresolver.Write)
}

func (w weixinAccountDo) Session(config *gorm.Session) *weixinAccountDo {
	return w.withDO(w.DO.Session(config))
}

func (w weixinAccountDo) Clauses(conds ...clause.Expression) *weixinAccountDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w weixinAccountDo) Returning(value interface{}, columns ...string) *weixinAccountDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w weixinAccountDo) Not(conds ...gen.Condition) *weixinAccountDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w weixinAccountDo) Or(conds ...gen.Condition) *weixinAccountDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w weixinAccountDo) Select(conds ...field.Expr) *weixinAccountDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w weixinAccountDo) Where(conds ...gen.Condition) *weixinAccountDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w weixinAccountDo) Order(conds ...field.Expr) *weixinAccountDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w weixinAccountDo) Distinct(cols ...field.Expr) *weixinAccountDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w weixinAccountDo) Omit(cols ...field.Expr) *weixinAccountDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w weixinAccountDo) Join(table schema.Tabler, on ...field.Expr) *weixinAccountDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w weixinAccountDo) LeftJoin(table schema.Tabler, on ...field.Expr) *weixinAccountDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w weixinAccountDo) RightJoin(table schema.Tabler, on ...field.Expr) *weixinAccountDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w weixinAccountDo) Group(cols ...field.Expr) *weixinAccountDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w weixinAccountDo) Having(conds ...gen.Condition) *weixinAccountDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w weixinAccountDo) Limit(limit int) *weixinAccountDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w weixinAccountDo) Offset(offset int) *weixinAccountDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w weixinAccountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *weixinAccountDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w weixinAccountDo) Unscoped() *weixinAccountDo {
	return w.withDO(w.DO.Unscoped())
}

func (w weixinAccountDo) Create(values ...*model.WeixinAccount) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w weixinAccountDo) CreateInBatches(values []*model.WeixinAccount, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w weixinAccountDo) Save(values ...*model.WeixinAccount) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w weixinAccountDo) First() (*model.WeixinAccount, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeixinAccount), nil
	}
}

func (w weixinAccountDo) Take() (*model.WeixinAccount, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeixinAccount), nil
	}
}

func (w weixinAccountDo) Last() (*model.WeixinAccount, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeixinAccount), nil
	}
}

func (w weixinAccountDo) Find() ([]*model.WeixinAccount, error) {
	result, err := w.DO.Find()
	return result.([]*model.WeixinAccount), err
}

func (w weixinAccountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WeixinAccount, err error) {
	buf := make([]*model.WeixinAccount, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w weixinAccountDo) FindInBatches(result *[]*model.WeixinAccount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w weixinAccountDo) Attrs(attrs ...field.AssignExpr) *weixinAccountDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w weixinAccountDo) Assign(attrs ...field.AssignExpr) *weixinAccountDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w weixinAccountDo) Joins(fields ...field.RelationField) *weixinAccountDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w weixinAccountDo) Preload(fields ...field.RelationField) *weixinAccountDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w weixinAccountDo) FirstOrInit() (*model.WeixinAccount, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeixinAccount), nil
	}
}

func (w weixinAccountDo) FirstOrCreate() (*model.WeixinAccount, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeixinAccount), nil
	}
}

func (w weixinAccountDo) FindByPage(offset int, limit int) (result []*model.WeixinAccount, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w weixinAccountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w weixinAccountDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w weixinAccountDo) Delete(models ...*model.WeixinAccount) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *weixinAccountDo) withDO(do gen.Dao) *weixinAccountDo {
	w.DO = *do.(*gen.DO)
	return w
}