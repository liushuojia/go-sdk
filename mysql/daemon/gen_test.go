package cmd

import (
	mysqlConn "github.com/liushuojia/go-sdk/mysql"
	"testing"
)

func TestGen(t *testing.T) {
	mysqlConn.Build("root", "liushuojia", "127.0.0.1", 3306, "app", "./gen/model")

	//db, err := mysqlConn.GormDB("root", "liushuojia", "127.0.0.1", 3306, "abc")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//query.SetDefault(db)
	//query.Admin.Where(query.Admin.ID.Eq(1)).Take()
}
