package redisConn

import (
	"fmt"
	"github.com/liushuojia/go-sdk/ssh"
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"
	"time"
)

// 目录下 go test
func TestConnectRedis(t *testing.T) {
	log.Println("redis test")

	client := ssh.SSH{
		Host:    "www.liushuojia.com",
		User:    "root",
		Port:    22,
		KeyFile: "/Users/liushuojia/.ssh/id_rsa",
		Type:    "KEY", // PASSWORD or KEY
	}
	c, _ := client.Dial()

	r := New().SetAddr("127.0.0.1:6379").SetPassword("liushuojia").SetDB(1).ConnectWithSSH(c)

	fmt.Println("aaa")
	fmt.Println(r.Set(r.ctx, "a", "nihao", time.Second*100).Err())
	fmt.Println("bbb")

	go r.Reader(
		func(channel, message string) error {
			fmt.Println("channel:", channel, "message:", message)
			return nil
		},
		"topic",
		"topic2",
	)

	//time.Sleep(3 * time.Second)
	//
	//fmt.Println("r.close")
	//r.Close()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 3)
		r.Publish(r.ctx, "topic", "topic a "+strconv.Itoa(i))
	}

	//fmt.Println("end")
	time.Sleep(60 * time.Second)
}
