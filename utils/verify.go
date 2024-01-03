package utils

import (
	"regexp"
	"strings"
)

// VerifyEmail 验证邮箱
func VerifyEmail(email string) bool {
	if email == "" {
		return false
	}
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	email = strings.ToLower(email)
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyMobile 验证手机
func VerifyMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	regular := "^\\d{11}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

// VerifyNum Number
func VerifyNum(numString string) bool {
	regular := "^-?\\d+$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(numString)
}
