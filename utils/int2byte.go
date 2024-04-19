package utils

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// IntToBytes 整形转换成字节
func IntToBytes(n int64) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()

}

// BytesToInt 字节转换成整形
func BytesToInt(b []byte) int64 {
	var x int32
	bytesBuffer := bytes.NewBuffer(b)
	if err := binary.Read(bytesBuffer, binary.BigEndian, &x); err == nil {
		return int64(x)
	}

	i, err := strconv.Atoi(string(b))
	if err != nil {
		return 0
	}
	return int64(i)
}
