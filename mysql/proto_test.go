package mysqlConn

import (
	"testing"
)

// 目录下 go test
func TestProto(t *testing.T) {
	BuildProto(
		"./proto.message",
		"127.0.0.1",
		3306,
		"root",
		"liushuojia",
		"APP_admin",
		"deleted_at",
	)
}
