package mysqlConn

import (
	"fmt"
	"github.com/liushuojia/go-sdk/ssh"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestSSH(t *testing.T) {
	client := ssh.SSH{
		Host:    "www.liushuojia.com",
		User:    "root",
		Port:    22,
		KeyFile: "/Users/liushuojia/.ssh/id_rsa",
		Type:    "KEY", // PASSWORD or KEY
	}

	c, err := client.Dial()
	if err != nil {
		log.Fatalln(err)
	}

	my := MySQL{
		Host:     "localhost",
		User:     "root",
		Password: "liushuojia",
		Port:     3306,
		Database: "APPUSERDB",
	}

	conn, err := my.ConnWithSSH(c)
	if err != nil {
		log.Fatalln(err)
		return
	}

	val := make(map[string]interface{})
	if err := conn.Table("admin").Where("id = ?", 1).Find(&val).Error; err != nil {
		log.Fatalf("mysql query error: %s", err.Error())
		return
	}
	fmt.Println(val)
}
