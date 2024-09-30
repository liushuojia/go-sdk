package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"reflect"
)

/*
GobRegister 使用前需要注册下
*/
func GobRegister(value interface{}) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Ptr:
	case reflect.Struct:
		gob.Register(value)
	default:
	}
}
func GobAny2Byte(value interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	GobRegister(value)

	enc := gob.NewEncoder(buf)
	if err := enc.Encode(&value); err != nil {
		return nil, err
	}

	//if err := binary.Write(buf, binary.BigEndian, value); err != nil {
	//	return nil, err
	//}

	return buf.Bytes(), nil
}
func GobByte2Any(valueBytes []byte, value interface{}) error {

	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(value); err != nil {
		return err
	}

	//if err := binary.Read(buf, binary.BigEndian, &value); err != nil {
	//	return err
	//}
	return nil
}

func BinaryAny2Byte(value interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func BinaryByte2Any(valueBytes []byte, value interface{}) error {
	buf := bytes.NewBuffer(valueBytes)
	if err := binary.Read(buf, binary.BigEndian, value); err != nil {
		return err
	}
	return nil
}
