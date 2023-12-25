package redisConn

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 目录下 go test
func TestRedis(t *testing.T) {
	t.Log("redis test")

	r := New().SetAddr("127.0.0.1:6379").SetPassword("liushuojia").Connect()

	r.Set(r.ctx, "a", "nihao", time.Second*100).Err()

	go r.Reader(
		func(channel, message string) {
			fmt.Println("channel:", channel, "message:", message)
		},
		"topic",
		"topic2",
	)

	//time.Sleep(3 * time.Second)
	//
	//fmt.Println("r.close")
	//r.Close()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		r.Publish(r.ctx, "topic", "topic a "+strconv.Itoa(i))
	}

	//fmt.Println("end")
	time.Sleep(2 * time.Second)
}
