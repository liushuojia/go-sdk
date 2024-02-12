package main

import (
	"flag"
	mysqlConn "github.com/liushuojia/go-sdk/mysql"
)

var (
	username = flag.String("u", "root", "-u username")
	password = flag.String("p", "liushuojia", "-p password")
	host     = flag.String("h", "127.0.0.1", "-h host")
	port     = flag.Int("P", 3306, "-P port")
	db       = flag.String("db", "abc", "-db database")
)

func main() {
	mysqlConn.Build(*username, *password, *host, *port, *db, "/Volumes/work/app-aliyun/sdk/example/query")
}
