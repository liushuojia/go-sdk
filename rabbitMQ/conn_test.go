package rabbitMQConn

import (
	"fmt"
	"github.com/streadway/amqp"
	"testing"
	"time"
)

// 目录下 go test
func TestRabbitMQConn(t *testing.T) {

	c := New()
	c.SetUrl("admin", "admin", "localhost", 5672, "/")
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
		func(message amqp.Delivery) error {
			fmt.Println(message)
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
		func(message amqp.Delivery) error {
			fmt.Println(message)
			return nil
		},
		"my-queue",
	)

	//go func() {
	//	for i := 0; i < 30; i++ {
	//		s := []byte("hi " + strconv.Itoa(i))
	//		fmt.Println("============================================")
	//		fmt.Println(string(s))
	//		fmt.Println(c.PublishQueue("my-queue", s))
	//		time.Sleep(time.Second)
	//	}
	//}()

	time.Sleep(time.Minute * 3)
}
