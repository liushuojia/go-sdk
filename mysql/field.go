package mysqlConn

import (
	"errors"
	"reflect"
	"sync"
)

var structMap sync.Map

// InitUpdateMap
// user := model.user
// InitUpdateMap(user, updateMap) 更新数据库前清理
func InitUpdateMap(o interface{}, updateMap map[string]interface{}) error {
	if o == nil {
		return errors.New("更新结构为空")
	}
	if updateMap == nil {
		return errors.New("更新内容为空")
	}

	delete(updateMap, "id")
	delete(updateMap, "created_at")
	delete(updateMap, "updated_at")
	delete(updateMap, "deleted_at")
	if len(updateMap) <= 0 {
		return errors.New("更新内容为空")
	}

	fieldMap, err := GetFieldMap(o)
	if err != nil {
		return err
	}

	for k, v := range updateMap {
		fieldValue, ok := fieldMap[k]
		if !ok {
			delete(updateMap, k)
			continue
		}

		n := false
		switch fieldValue {
		case reflect.Int64:
		case reflect.String:
		default:
			n = true
		}
		if n {
			delete(updateMap, k)
			continue
		}

		switch v.(type) {
		case float64:
			if fieldValue != reflect.Int64 {
				delete(updateMap, k)
				break
			}
			updateMap[k] = int64(v.(float64))
		case string:
			if fieldValue != reflect.String {
				delete(updateMap, k)
			}
		default:
			delete(updateMap, k)
		}
	}
	return nil
}

func GetFieldMap(o interface{}) (map[string]reflect.Kind, error) {

	fieldMap := make(map[string]reflect.Kind)

	getType := reflect.TypeOf(o)
	getValue := reflect.ValueOf(o)
	if getType.Kind() == reflect.Ptr {
		getType = getType.Elem()
		getValue = getValue.Elem()
	}
	if getType.Kind() != reflect.Struct {
		return nil, errors.New("只能处理结构变量")
	}

	if value, ok := structMap.Load(getType.Name()); ok {
		if v, ok := value.(map[string]reflect.Kind); ok {
			return v, nil
		}
	}

	for i := 0; i < getValue.NumField(); i++ {
		fieldMap[getType.Field(i).Tag.Get("json")] = getValue.Field(i).Kind()
	}
	if len(fieldMap) <= 0 {
		return nil, errors.New("结构体无字段")
	}
	structMap.Store(getType.Name(), fieldMap)

	return fieldMap, nil
}
