package kafkaConn

import (
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"
	"time"
)

// 目录下 go test
func TestKafka(t *testing.T) {
	t.Log("test kafka")

	n := New().SetBrokers("localhost:9092").SetTopicPrefix("ppp-").SetGroupID("my002")

	go n.Reader("my-topic", 0, func(topic string, key, value []byte) error {
		log.Println(
			"[subscribe]",
			"reader",
			"topic:", topic,
			"Key:", string(key),
			"value:", string(value),
		)
		return nil
	})
	go n.Reader("my-topic-one", 0, func(topic string, key, value []byte) error {
		log.Println(
			"[subscribe]",
			"reader",
			"topic:", topic,
			"Key:", string(key),
			"value:", string(value),
		)
		return nil
	})
	go n.Reader("my-topic-two", 0, func(topic string, key, value []byte) error {
		log.Println(
			"[subscribe]",
			"reader",
			"topic:", topic,
			"Key:", string(key),
			"value:", string(value),
		)
		return nil
	})

	//time.Sleep(3 * time.Second)
	//n.Close()

	//write()
	time.Sleep(time.Second * 100)
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
