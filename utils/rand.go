package utils

import (
	"math/rand"
	"time"
)

func RandString(l int) string {
	bytesTmp := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytesTmp[r.Intn(len(bytesTmp))])
	}
	return string(result)
}

func RandNumber(l int) string {
	bytesTmp := []byte("0123456789")
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytesTmp[r.Intn(len(bytesTmp))])
	}
	return string(result)
}
