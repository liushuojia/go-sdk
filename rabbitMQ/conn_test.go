package rabbitMQConn

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
	"testing"
	"time"
)

// 目录下 go test
func TestRabbitMQConn(t *testing.T) {

	c := New().SetUrl("admin", "liushuojia", "106.55.61.225", 33003, "/")
	if err := c.Connect(); err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(10 * time.Second)
		c.Close()
	}()

	//c.CreateExchange("logs-top", amqp.ExchangeTopic)
	//c.CreateExchange("logs", amqp.ExchangeTopic)
	//c.CreateExchangeBind("logs", "logs-top", "log", "tt", "")
	//c.CreateQueue("page", "email", "firehose")
	//c.CreateQueueBind("page", "logs", "alert", "a")
	//c.CreateQueueBind("email", "logs", "info", "a", "b")
	//c.CreateQueueBind("firehose", "logs", "#", "a", "c")
	//c.SubscribeExchange("my-topic", amqp.ExchangeTopic, "aa")

	go c.SubscribeExchange(
		func(exchange, routingKey string, body []byte) error {
			fmt.Println(
				"exchange:", exchange,
				"routingKey:", routingKey,
				"body:", string(body),
			)
			return nil
		},
		"my-topic",
		amqp.ExchangeTopic,
		"my-queue",
		"aa",
	)

	//go func() {
	//	for i := 0; i < 30; i++ {
	//		c.PublishExchange("my-topic", amqp.ExchangeTopic, "aa", []byte("hi "+strconv.Itoa(i)))
	//		time.Sleep(time.Second)
	//	}
	//}()

	go c.SubscribeQueue(
		func(exchange, routingKey string, body []byte) error {
			fmt.Println(
				"exchange:", exchange,
				"routingKey:", routingKey,
				"body:", string(body),
			)
			return nil
		},
		"my-queue",
	)

	go func() {
		for i := 0; i < 30; i++ {
			s := []byte("hi " + strconv.Itoa(i))
			fmt.Println("============================================")
			fmt.Println(string(s))
			fmt.Println(c.PublishQueue("my-queue", s))
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Minute * 3)
}
