package kafkaConn

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"strconv"
	"testing"
	"time"
)

// 目录下 go test
func TestHelloWorld(t *testing.T) {
	t.Log("send email")

	go read()
	write()

	time.Sleep(time.Second * 100)
}

func read() {
	New().SetBrokers("localhost:9092").SetTopicPrefix("ppp-").SetGroupID("my002").
		Reader("my-topic", 0, func(message kafka.Message) error {
			fmt.Println(message)
			return nil
		})
}

func write() {
	c := New().SetBrokers("localhost:9092").SetTopicPrefix("ppp-").SetGroupID("my002")

	for i := 1; i < 3; i++ {
		time.Sleep(1 * time.Second)

		c.Writer(
			"my-topic",
			kafka.Message{
				Key:   []byte("Key-A " + strconv.Itoa(i)),
				Value: []byte("Hello World!" + strconv.Itoa(i)),
			},
			//kafka.Message{
			//	Key:   []byte("Key-B"),
			//	Value: []byte("One!"),
			//},
			//kafka.Message{
			//	Key:   []byte("Key-C"),
			//	Value: []byte("Two!"),
			//},
		)
	}
}
