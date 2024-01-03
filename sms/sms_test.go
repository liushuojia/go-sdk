package smsConn

import (
	"testing"
)

// 目录下 go test
func TestSMS(t *testing.T) {
	t.Log("sms test")

	New().SetMobile("13725588389").SetType("code").SetTemplateParam("code", "123456").Send()
}
