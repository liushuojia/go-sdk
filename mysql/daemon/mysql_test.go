package cmd

import (
	"testing"
)

func TestMysql(t *testing.T) {
	//t.Log("mysql test")

	//
	//mysqlConn.UseDB(db)
	//fmt.Println((&model.FileDefault{
	//	DataCode: "company",
	//}).TableName())
	//db.AutoMigrate(&model.Admin{})
	//db.AutoMigrate(&model.FileDefault{})
	//
	//admin := model.Admin{}
	//admin.Select(1)
	//admin.SetTableNameSuffix("you")
	//admin.Create()

	//var adminList []*model.Admin
	//err = mysqlConn.QueryData(
	//	db,
	//	[]string{"id", "realname"},
	//	[]string{"id", "realname"},
	//	[]string{"id desc"},
	//	&adminList,
	//)
	//fmt.Println(err)
	//for _, v := range adminList {
	//	fmt.Println(v.ID, v.Realname)
	//}

	//admin.Select(3)

	//admin.Realname = "liushuojia"
	//admin.Update()
	//
	//fmt.Println(admin)

	//admin.Change(map[string]interface{}{
	//	"username": "xiao liu 123",
	//})
	//fmt.Println(admin)
	//fmt.Println(admin.Select(1))
	//return

	//fmt.Println(
	//	admin.Query(
	//		map[string]any{
	//			"Username_like": "xiao",
	//			"order_by": []string{
	//				"realname",
	//				"id desc",
	//				"abbcc",
	//			},
	//			"id_in": []int64{
	//				1, 2, 3, 4,
	//			},
	//			"updated_at_min": time.Now().AddDate(0, -3, 0),
	//		},
	//		0,
	//		0,
	//	),
	//)

	//return
	//admin2 := &model.Admin{}
	//admin2.SetTableNameSuffix("you")
	//fmt.Println(admin2.Select(1))
	//fmt.Println(admin2.GetTableNameSuffix())

	//admin3 := &model.Admin{}
	//admin3.SetTableNameCode("you_aabbcc").SetDB(db)

	//db.Scopes(query.TableOfCode(admin, "you")).AutoMigrate(admin)

	//db.
	//	Scopes(query.TableOfCode(admin, "you")).
	//	AutoMigrate(admin)

	//admin.Create()
	//fmt.Println(query.DataFieldList.GetFieldList(admin))
}
