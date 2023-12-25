package main

import (
	"flag"
	"fmt"
	"github.com/liushuojia/go-sdk/example/model"
	"github.com/liushuojia/go-sdk/example/query"
	mysqlConn "github.com/liushuojia/go-sdk/mysql"
	"log"
)

var (
	username = flag.String("u", "root", "-u username")
	password = flag.String("p", "liushuojia", "-p password")
	host     = flag.String("h", "127.0.0.1", "-h host")
	port     = flag.Int("P", 3306, "-P port")
	db       = flag.String("db", "my_test", "-db database")
)

func main() {

	conn, err := mysqlConn.GormDB(*username, *password, *host, *port, *db)
	if err != nil {
		log.Fatalln(err)
	}
	query.SetDefault(conn)

	user := query.User
	userCreate := model.User{
		Name: "张三",
		Addr: "深圳",
	}
	user2Create := model.User{
		Name: "李四",
		Addr: "广州",
	}

	fmt.Println(user.Create(&userCreate, &user2Create))
	fmt.Println(userCreate)
	fmt.Println(user2Create)

	fmt.Println(user.Where(user.ID.Eq(userCreate.ID)).UpdateSimple(
		user.Name.Value("深圳市1"),
		user.Addr.Value("深圳市2"),
	))

	fmt.Println(user.Find())

	fmt.Println(user.Where(user.ID.Eq(userCreate.ID)).Delete())

}
