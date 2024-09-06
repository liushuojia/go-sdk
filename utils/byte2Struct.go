package utils

import "unsafe"

type sliceData struct {
	value interface{}
}

type sliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func Encode(data any) []byte {
	v := &sliceData{
		value: data,
	}
	size := unsafe.Sizeof(*v)
	testBytes := &sliceMock{
		addr: uintptr(unsafe.Pointer(v)),
		cap:  int(size),
		len:  int(size),
	}
	return *(*[]byte)(unsafe.Pointer(testBytes))
}
func Decode(data []byte) any {
	ptestStruct := *(**sliceData)(unsafe.Pointer(&data))
	return ptestStruct.value
}
