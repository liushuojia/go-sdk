package email

import (
	"testing"
)

// 目录下 go test
func TestEmail(t *testing.T) {
	t.Log("send email")

	New().
		SetHost(
			"****@qq.com",
			"****",
			"****",
			"smtp.qq.com",
			465,
		).
		SetBody("主体内容").
		SetSubject("标题").
		AddMailTo("刘硕嘉", "liushuojia@qq.com").
		Send()
}
