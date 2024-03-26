package mysqlConn

import (
	"errors"
	"reflect"
	"sync"
)

// InitUpdateMap
// user := model.user
// InitUpdateMap(user, updateMap) 更新数据库前清理
//func InitUpdateMap(o interface{}, updateMap map[string]interface{}) error {
//	if o == nil {
//		return errors.New("更新结构为空")
//	}
//	if updateMap == nil {
//		return errors.New("更新内容为空")
//	}
//
//	delete(updateMap, "id")
//	delete(updateMap, "created_at")
//	delete(updateMap, "updated_at")
//	delete(updateMap, "deleted_at")
//	if len(updateMap) <= 0 {
//		return errors.New("更新内容为空")
//	}
//
//	fieldMap, err := GetFieldMap(o)
//	if err != nil {
//		return err
//	}
//
//	for k, v := range updateMap {
//		fieldValue, ok := fieldMap[k]
//
//		if !ok {
//			delete(updateMap, k)
//			continue
//		}
//
//		// 判断数据库类型，数据库只支持 int64 string
//		n := false
//		switch fieldValue {
//		case reflect.Int64:
//		case reflect.String:
//		default:
//			n = true
//		}
//		if n {
//			delete(updateMap, k)
//			continue
//		}
//
//		// 判断updateMap的值
//		switch t := v.(type) {
//		case int:
//			updateMap[k] = int64(t)
//		case int64:
//			if fieldValue != reflect.Int64 {
//				delete(updateMap, k)
//				break
//			}
//		case float64:
//			if fieldValue != reflect.Int64 {
//				delete(updateMap, k)
//				break
//			}
//			updateMap[k] = int64(t)
//		case string:
//			if fieldValue != reflect.String {
//				delete(updateMap, k)
//			}
//		default:
//			delete(updateMap, k)
//		}
//	}
//	return nil
//}

func GetFieldMap(o interface{}) (map[string]reflect.Kind, error) {
	tmp, err := GetModelType(o)
	if err != nil {
		return nil, err
	}
	if len(tmp.Json) <= 0 {
		return nil, errors.New("结构体JSON无字段")
	}
	return tmp.Json, nil
}

/*
modelTypeMap
*/
var modelTypeMap sync.Map

type typeInterface struct {
	Field      map[string]reflect.Kind
	Field2Json map[string]string
	Json       map[string]reflect.Kind
	Value      map[string]interface{}
}

func GetModelType(o interface{}) (*typeInterface, error) {
	getType := reflect.TypeOf(o)
	getValue := reflect.ValueOf(o)
	if getType.Kind() == reflect.Ptr {
		getType = getType.Elem()
		getValue = getValue.Elem()
	}
	if getType.Kind() != reflect.Struct {
		return nil, errors.New("只能处理结构变量")
	}
	if value, ok := modelTypeMap.Load(getType.Name()); ok {
		if data, ok := value.(*typeInterface); ok {
			return data, nil
		}
	}

	v := &typeInterface{
		Field:      make(map[string]reflect.Kind),
		Json:       make(map[string]reflect.Kind),
		Field2Json: make(map[string]string),
		Value:      make(map[string]interface{}),
	}

	for i := 0; i < getType.NumField(); i++ {
		fieldName := getType.Field(i).Name
		jsonName := getType.Field(i).Tag.Get("json")
		kindName := getValue.Field(i).Kind()

		v.Field[fieldName] = kindName
		v.Field2Json[fieldName] = jsonName
		v.Json[jsonName] = kindName
		switch kindName {
		case reflect.Int64:
			v.Value[jsonName] = getValue.Field(i).Int()
		case reflect.String:
			v.Value[jsonName] = getValue.Field(i).String()
		default:
		}
	}

	modelTypeMap.Store(getType.Name(), v)
	return v, nil
}
func GetUpdateMap(modelValue interface{}, updateValue interface{}) map[string]interface{} {
	updateMap := make(map[string]interface{})

	getType := reflect.TypeOf(updateValue)
	getValue := reflect.ValueOf(updateValue)
	if getType.Kind() == reflect.Ptr {
		getType = getType.Elem()
		getValue = getValue.Elem()
	}

	modelTypeData, err := GetModelType(modelValue)
	if err != nil {
		return nil
	}

	switch getType.Kind() {
	case reflect.Struct:
		for i := 0; i < getValue.NumField(); i++ {
			fieldName := getType.Field(i).Tag.Get("json")

			if fieldName != "" {
				// json
				m, ok := modelTypeData.Json[fieldName]
				if !ok || m != getValue.Field(i).Kind() {
					continue
				}

				switch m {
				case reflect.Int64:
					updateMap[fieldName] = getValue.Field(i).Int()
				case reflect.String:
					updateMap[fieldName] = getValue.Field(i).String()
				default:
				}
			} else {
				// FieldKey
				fieldName = getType.Field(i).Name
				m, ok := modelTypeData.Field[fieldName]
				if !ok || m != getValue.Field(i).Kind() {
					continue
				}

				fieldName = modelTypeData.Field2Json[fieldName]
				switch m {
				case reflect.Int64:
					updateMap[fieldName] = getValue.Field(i).Int()
				case reflect.String:
					updateMap[fieldName] = getValue.Field(i).String()
				default:
				}
			}
		}
	case reflect.Map:
		// golang json body 转过来
		// 字段格式 string float64
		dataMap, ok := updateValue.(map[string]interface{})
		if !ok {
			break
		}

		for k, v := range dataMap {
			m, ok := modelTypeData.Json[k]
			if !ok {
				continue
			}

			dataType := reflect.TypeOf(v)
			dataValue := reflect.ValueOf(v)
			if getType.Kind() == reflect.Ptr {
				dataType = dataType.Elem()
				dataValue = dataValue.Elem()
			}

			if m == dataValue.Kind() {
				// model 字段格式只有 int64 string
				switch m {
				case reflect.Int64:
					updateMap[k] = v.(int64)
				case reflect.String:
					updateMap[k] = v.(string)
				default:
				}
				continue
			}

			if m != reflect.Int64 {
				continue
			}

			switch dataValue.Kind() {
			case reflect.Float64:
				updateMap[k] = int64(dataValue.Float())
			case reflect.Int:
				updateMap[k] = dataValue.Int()
			default:
			}

		}
	default:
	}

	delete(updateMap, "id")
	delete(updateMap, "created_at")
	delete(updateMap, "updated_at")
	delete(updateMap, "deleted_at")
	return updateMap
}
