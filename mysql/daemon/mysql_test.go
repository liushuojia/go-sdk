package cmd

import (
	"fmt"
	mysqlConn "github.com/liushuojia/go-sdk/mysql"
	"github.com/liushuojia/go-sdk/mysql/daemon/model"
	"log"
	"testing"
)

func TestMysql(t *testing.T) {
	t.Log("mysql test")

	db, err := mysqlConn.GormDB("root", "liushuojia", "127.0.0.1", 3306, "abc")
	if err != nil {
		log.Fatalln(err)
	}

	mysqlConn.UseDB(db)
	//fmt.Println((&model.FileDefault{
	//	DataCode: "company",
	//}).TableName())
	//db.AutoMigrate(&model.Admin{})
	//db.AutoMigrate(&model.FileDefault{})

	admin := &model.Admin{}
	fmt.Println(admin.Select(1))

	fmt.Println(
		admin.Query(
			map[string]any{
				"Username_like": "xiao",
			},
			0,
			0,
		),
	)

	admin2 := &model.Admin{}
	admin2.SetTableNameSuffix("you")
	fmt.Println(admin2.Select(1))
	fmt.Println(admin2.GetTableNameSuffix())

	//admin3 := &model.Admin{}
	//admin3.SetTableNameCode("you_aabbcc").SetDB(db)

	//db.Scopes(query.TableOfCode(admin, "you")).AutoMigrate(admin)

	//db.
	//	Scopes(query.TableOfCode(admin, "you")).
	//	AutoMigrate(admin)

	//admin.Create()
	//fmt.Println(query.DataFieldList.GetFieldList(admin))
}
