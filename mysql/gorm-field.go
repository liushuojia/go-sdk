package mysqlConn

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var field = NewTableField()

func DefaultField() *Table {
	return field
}

func NewTableField() *Table {
	return &Table{}
}

type TableField struct {
	SQLFieldList []string                `json:"sql_field_list"` // 数据库字段
	GormMap      map[string]GormField    `json:"gorm_map"`       // key sql字段			: value 类型
	FieldType    map[string]reflect.Type `json:"field_type"`     // key model KEY		: value 类型
	FieldMap     map[string]string       `json:"field_map"`      // key model KEY		: sql字段
	JsonMap      map[string]string       `json:"json_map"`       // key jsonName		: sql字段
	AutoMigrate  map[string]bool         `json:"auto_migrate"`   // 表是否初始化=
}

type GormField struct {
	Column        string `json:"column"`
	Type          string `json:"type"`
	TypeLength    string `json:"type_length"`
	Comment       string `json:"comment"`
	Default       string `json:"default"`
	PrimaryKey    bool   `json:"primary_key"`
	Unique        bool   `json:"unique"`
	NotNull       bool   `json:"not_null"`
	AutoIncrement bool   `json:"auto_increment"`
	Index         bool   `json:"index"`
}

type m interface {
	TableName() string
	GetDB() *gorm.DB
}

type Table struct {
	data        sync.Map
	autoMigrate sync.Map
}

func (t *Table) Get(name string) *TableField {
	if tableFieldMap, ok := t.data.Load(name); ok {
		if d, ok := tableFieldMap.(TableField); ok {
			return &d
		}
	}
	return nil
}
func (t *Table) Set(o m) error {
	if d := t.Get(o.TableName()); d != nil {
		return nil
	}
	getType := reflect.TypeOf(o)
	if getType.Kind() != reflect.Struct {
		getType = getType.Elem()
	}
	if getType.Kind() != reflect.Struct {
		return errors.New("只能处理结构变量")
	}

	tableFieldData := TableField{
		GormMap:   make(map[string]GormField),
		FieldType: make(map[string]reflect.Type),
		FieldMap:  make(map[string]string),
		JsonMap:   make(map[string]string),
	}

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		gormString := field.Tag.Get("gorm")
		if gormString == "" || gormString == "-" {
			continue
		}

		stringMap := make(map[string]string)
		boolMap := make(map[string]bool)
		for _, v := range strings.Split(gormString, ";") {
			v = strings.Trim(v, " ")
			if v == "" {
				continue
			}
			tmp := strings.Split(v, ":")
			if len(tmp) == 1 {
				boolMap[tmp[0]] = true
				continue
			}
			stringMap[tmp[0]] = tmp[1]
		}

		typeString := strings.ToLower(stringMap["type"])
		typeLength := ""
		tmp1 := strings.Split(typeString, "(")
		if len(tmp1) > 1 {
			typeString = tmp1[0]
			typeLength = strings.Trim(tmp1[1], ")")
		}

		gormField := GormField{
			Column:        stringMap["column"],
			Type:          typeString,
			TypeLength:    typeLength,
			Comment:       stringMap["comment"],
			Default:       stringMap["default"],
			PrimaryKey:    boolMap["primaryKey"],
			Unique:        boolMap["unique"],
			NotNull:       boolMap["not null"],
			AutoIncrement: boolMap["autoIncrement"],
			Index:         boolMap["index"],
		}

		jsonName := field.Tag.Get("json")
		if index := strings.Index(jsonName, ","); index >= 0 {
			jsonName = jsonName[:index]
		}
		jsonName = strings.Trim(jsonName, " ")

		tableFieldData.GormMap[gormField.Column] = gormField
		tableFieldData.FieldType[field.Name] = field.Type
		tableFieldData.FieldMap[field.Name] = gormField.Column
		tableFieldData.JsonMap[jsonName] = gormField.Column
		tableFieldData.SQLFieldList = append(tableFieldData.SQLFieldList, gormField.Column)
	}

	t.data.Store(o.TableName(), tableFieldData)
	return nil
}
func (t *Table) GetAndSet(o m) *TableField {
	if tableFieldMap := t.Get(o.TableName()); tableFieldMap != nil {
		return tableFieldMap
	}
	if err := t.Set(o); err != nil {
		return nil
	}
	return t.Get(o.TableName())
}
func (t *Table) GetFieldList(o m) []string {
	if d := t.GetAndSet(o); d != nil {
		return d.SQLFieldList
	}
	return nil
}
func (t *Table) CleanUpdateMap(o m, updateMap map[string]interface{}) map[string]interface{} {
	tableField := t.GetAndSet(o)
	if tableField == nil {
		return nil
	}

	for _, key := range []string{"id", "created_at", "updated_at", "deleted_at"} {
		delete(updateMap, key)
	}

	returnUpdateMap := make(map[string]interface{})
	if len(updateMap) > len(tableField.GormMap) {
		for k := range tableField.GormMap {
			if value, ok := updateMap[k]; ok {
				returnUpdateMap[k] = value
			}
		}
	} else {
		for k, value := range updateMap {
			if _, ok := tableField.GormMap[k]; ok {
				returnUpdateMap[k] = value
			}
		}
	}

	return returnUpdateMap
}
func (t *Table) GetAutoMigrate(name string) bool {
	if f, ok := t.autoMigrate.Load(name); ok {
		if d, ok := f.(bool); ok {
			return d
		}
	}
	return false
}
func (t *Table) SetAutoMigrate(name string, f bool) {
	t.autoMigrate.Store(name, f)
}
func (t *Table) QueryDB(db *gorm.DB, fieldMap map[string]string, searchMap map[string]any) *gorm.DB {
	for field, sqlField := range fieldMap {
		field = strings.Trim(field, " ")
		sqlField = strings.Trim(sqlField, " ")
		if field == "" || sqlField == "" {
			continue
		}

		if value, ok := searchMap[field]; ok && any2String(value) != "" {
			db = db.Where(sqlField+" = ?", value)
		}
		if value, ok := searchMap[field+"_min"]; ok && any2String(value) != "" {
			db = db.Where(sqlField+" >= ?", value)
		}
		if value, ok := searchMap[field+"_max"]; ok && any2String(value) != "" {
			db = db.Where(sqlField+" <= ?", value)
		}

		if value, ok := searchMap[field+"_is_null"]; ok && any2String(value) != "" {
			db = db.Where(sqlField + " is null")
		}
		if value, ok := searchMap[field+"_is_empty"]; ok && any2String(value) != "" {
			db = db.Where(sqlField + " = ''")
		}

		if value, ok := searchMap[field+"_len"]; ok {
			if l, err := any2Number(value); err == nil && l > 0 {
				db = db.Where("length("+sqlField+") = ? ", l)
			}
		}
		if value, ok := searchMap[field+"_len_min"]; ok && value != "" {
			if l, err := any2Number(value); err == nil && l > 0 {
				db = db.Where("length("+sqlField+") >= ? ", l)
			}
		}
		if value, ok := searchMap[field+"_len_max"]; ok && value != "" {
			if l, err := any2Number(value); err == nil && l > 0 {
				db = db.Where("length("+sqlField+") <= ? ", l)
			}
		}
		if value, ok := searchMap[field+"_len_min_or_len_0"]; ok && value != "" {
			if l, err := any2Number(value); err == nil && l > 0 {
				db = db.Where("length("+sqlField+") >= ? or length("+sqlField+") = 0", l)
			}
		}

		// 全文检索 尽量少用
		if value, ok := searchMap[field+"_not"]; ok && any2String(value) != "" {
			db = db.Where(sqlField+" <> ?", value)
		}
		if value, ok := searchMap[field+"_like"]; ok && any2String(value) != "" {
			db = db.Where(sqlField+" LIKE ?", "%"+fmt.Sprintf("%v", value)+"%")
		}
		if value, ok := searchMap[field+"_left_like"]; ok && any2String(value) != "" {
			db = db.Where(sqlField+" LIKE ?", "%"+fmt.Sprintf("%v", value))
		}
		if value, ok := searchMap[field+"_right_like"]; ok && any2String(value) != "" {
			db = db.Where(sqlField+" LIKE ?", fmt.Sprintf("%v", value)+"%")
		}

		if value, ok := searchMap[field+"_in"]; ok {
			var tmp []interface{}
			if v, ok := value.(string); ok && v != "" {
				tt := strings.Split(v, ",")
				for _, vTmp := range tt {
					vTmp = strings.Trim(vTmp, " ")
					if vTmp != "" {
						tmp = append(tmp, vTmp)
					}
				}
			}
			if v, ok := value.([]interface{}); ok && len(v) > 0 {
				tmp = v
			}

			if len(tmp) > 0 {
				db = db.Where(sqlField+" in (?)", tmp)
			}
		}
		if value, ok := searchMap[field+"_not_in"]; ok {
			var tmp []interface{}
			if v, ok := value.(string); ok && v != "" {
				tt := strings.Split(v, ",")
				for _, vTmp := range tt {
					vTmp = strings.Trim(vTmp, " ")
					if vTmp != "" {
						tmp = append(tmp, vTmp)
					}
				}
			}
			if v, ok := value.([]interface{}); ok && len(v) > 0 {
				tmp = v
			}
			if len(tmp) > 0 {
				db = db.Where(sqlField+" not in (?)", tmp)
			}
		}

		if value, ok := searchMap[field+"_right_like_in"]; ok {
			sqlString := ""
			var interfaceList []interface{}
			if v, ok := value.(string); ok && v != "" {
				tt := strings.Split(v, ",")
				for _, vTmp := range tt {
					vTmp = strings.Trim(vTmp, " ")
					if vTmp == "" {
						continue
					}
					sqlString += " or " + sqlField + " like ?"
					interfaceList = append(interfaceList, v+"%")
				}
			}
			if v, ok := value.([]interface{}); ok && len(v) > 0 {
				for _, vTmp := range v {
					sqlString += " or " + sqlField + " like ?"
					interfaceList = append(interfaceList, fmt.Sprintf("%v", vTmp)+"%")
				}
			}
			if len(interfaceList) > 0 {
				db = db.Where("1=2"+sqlString, interfaceList...)
			}
		}
	}
	return db
}
func (t *Table) Query(o m, searchMap map[string]any) *gorm.DB {
	db := o.GetDB()
	tableField := t.GetAndSet(o)
	if tableField == nil {
		return db
	}

	gormMap := make(map[string]string)
	for k := range tableField.GormMap {
		gormMap[k] = k
	}
	db = t.QueryDB(db, gormMap, searchMap)

	//db = t.QueryDB(db, tableField.FieldMap, searchMap)
	//db = t.QueryDB(db, tableField.JsonMap, searchMap)
	return db
}

func any2String(v any) string {
	return fmt.Sprintf("%v", v)
}
func any2Number(v any) (int64, error) {
	var value int64
	switch d := v.(type) {
	case string:
		n, err := strconv.Atoi(d)
		if err != nil {
			return value, err
		}
		value = int64(n)
	case uint:
		value = int64(d)
	case uint8:
		value = int64(d)
	case uint16:
		value = int64(d)
	case uint32:
		value = int64(d)
	case uint64:
		value = int64(d)
	case float32:
		value = int64(d)
	case float64:
		value = int64(d)
	case int:
		value = int64(d)
	case int8:
		value = int64(d)
	case int16:
		value = int64(d)
	case int32:
		value = int64(d)
	case int64:
		value = d
	default:
		return value, errors.New("is not number")
	}
	return value, nil
}
