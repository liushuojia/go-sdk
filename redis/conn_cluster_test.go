package redisConn

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 目录下 go test
func TestRedisCluster(t *testing.T) {
	t.Log("redis cluster test")

	r := NewCluster().SetAddr(":7001").Connect()

	go r.Reader(
		func(channel, message string) {
			fmt.Println("channel:", channel, "message:", message)
		},
		"cluster-topic",
		"cluster-topic2",
	)

	//time.Sleep(3 * time.Second)
	//
	//fmt.Println("r.close")
	//r.Close()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		r.Publish(r.ctx, "cluster-topic", "cluster-topic a "+strconv.Itoa(i))
	}

	//fmt.Println("end")
	time.Sleep(2 * time.Second)
}
