package utils

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

/*
Copier
from to 类型必须是struct
*/
func Copier(to, from any) error {
	if from == nil {
		return errors.New("data is empty")
	}
	switch reflect.TypeOf(from).Kind() {
	case reflect.Map:
		return MapCopier(to, from)
	case reflect.Struct:
		return StructCopier(to, from)
	case reflect.Invalid:
		return errors.New("data is empty")
	default:
		return JsonCopier(to, from)
	}
}

func JsonCopier(to, from any) error {
	j, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(j, to)
}

func MapCopier(to any, m any) error {
	return mapstructure.Decode(m, &to)
}

func StructCopier(to, from any) error {
	return copier.Copy(to, from)
}
