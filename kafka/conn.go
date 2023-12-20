package kafkaConn

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	timeout    = 10 * time.Second
	writeRetry = 3
)

func New() *Conn {
	return (&Conn{}).SetContext(context.Background())
}

type Callback func(message kafka.Message) error
type Conn struct {
	brokers     []string
	topicPrefix string
	groupID     string

	ctx    context.Context
	cancel context.CancelFunc
}

func (c *Conn) SetBrokers(brokers ...string) *Conn {
	c.brokers = brokers
	return c
}
func (c *Conn) SetTopicPrefix(prefix string) *Conn {
	c.topicPrefix = prefix
	return c
}
func (c *Conn) SetGroupID(groupID string) *Conn {
	c.groupID = groupID
	return c
}
func (c *Conn) SetContext(ctx context.Context) *Conn {
	c.ctx, c.cancel = context.WithCancel(ctx)
	return c
}
func (c *Conn) Close() {
	c.ctx.Done()
	c.cancel()
}
func (c *Conn) Reader(topic string, partition int, cb Callback) {
	topic = c.topicPrefix + topic
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          c.brokers,
		Topic:            topic,
		Partition:        partition,
		MaxBytes:         10e6, // 10MB
		StartOffset:      kafka.LastOffset,
		GroupID:          c.groupID,
		ReadBatchTimeout: timeout,
	})
	defer r.Close()

	log.Println(
		"[subscribe]",
		"topic:", topic,
		"partition:", partition,
	)

	for {
		message, err := r.FetchMessage(c.ctx)
		if errors.Is(err, context.Canceled) {
			log.Println(
				"[subscribe]",
				"topic:", topic,
				"partition:", partition,
				"close",
			)
			break
		}
		if err != nil {
			log.Println(
				"[subscribe]",
				"reader message fail",
				"topic:", topic,
				"partition:", partition,
				"err:", err,
			)
			time.Sleep(time.Millisecond * 250)
			continue
		}

		log.Println(
			"[subscribe]",
			"reader",
			"topic:", message.Topic,
			"offset:", message.Offset,
			"partition:", message.Partition,
			"Key:", string(message.Key),
			"value:", string(message.Value),
		)
		if err := cb(message); err == nil {
			if err := r.CommitMessages(c.ctx, message); err != nil {
				log.Println(
					"[subscribe]",
					"failed to commit messages:",
					"topic:", message.Topic,
					"partition:", message.Partition,
					"err:", err,
				)
			}
		}
	}
}
func (c *Conn) Writer(topic string, messageList ...kafka.Message) error {
	w := &kafka.Writer{
		Addr:         kafka.TCP(c.brokers...),
		Topic:        c.topicPrefix + topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: timeout,
	}
	defer w.Close()

	var err error
	for i := 0; i < writeRetry; i++ {
		ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
		defer cancel()

		err = w.WriteMessages(ctx, messageList...)
		if err == nil {
			break
		}

		time.Sleep(time.Millisecond * 250)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			continue
		}
	}

	return err
}
