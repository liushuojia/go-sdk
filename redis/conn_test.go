package redisConn

import (
	"context"
	"fmt"
	"github.com/liushuojia/go-sdk/ssh"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"testing"
	"time"
)

type Hook1 struct{}

func (Hook1) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}
func (Hook1) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		fmt.Println("hook-1 start")
		if err := next(ctx, cmd); err != nil {
			return err
		}
		fmt.Println("hook-1 end")
		return nil
	}
}
func (Hook1) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}

type Hook2 struct{}

func (Hook2) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}
func (Hook2) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		fmt.Println("hook-2 start")
		fmt.Println(cmd.Name())
		fmt.Println(cmd.Args())
		fmt.Println(cmd.String())
		fmt.Println(cmd.FullName())
		if err := next(ctx, cmd); err != nil {
			return err
		}
		fmt.Println("hook-2 end")
		return nil
	}
}
func (Hook2) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}

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
	r.SetAddr("www.liushuojia.com:33001").Connect()
	r.AddHook(Hook1{})
	r.AddHook(Hook2{})

	fmt.Println("")
	fmt.Println(r.Set(r.ctx, "a", "nihao", time.Second*100).Result())
	fmt.Println("")
	fmt.Println(r.Get(r.ctx, "a").Result())

	return

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
