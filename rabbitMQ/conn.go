package rabbitMQConn

import (
	"context"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

const (
	reconnectDelay   = 8 * time.Second // 连接断开后多久重连
	reconnectMaxTime = 0               // 发送消息或订阅时，等待重连次数 0 一直重复连接

)

func New() *Conn {
	return (&Conn{}).SetContext(context.Background())
}

type Callback func(message amqp.Delivery) error

type Conn struct {
	connection    *amqp.Connection
	channel       *amqp.Channel
	url           string
	ctx           context.Context
	cancel        context.CancelFunc
	m             sync.Mutex       // 对closeChan关闭上锁
	isConnected   bool             // 是否已连接
	notifyClose   chan *amqp.Error // 异常关闭通知
	notifyConnect chan struct{}    // 连接成功通知
	subScribeMap  sync.Map         // 订阅或事件回调
}

func (r *Conn) Connect() error {
	log.Println("[rabbitMQ]", "connect rabbitMQ")
	if err := r.init().connect(); err != nil {
		return errors.New("rabbitMQ connect fail")
	}
	go r.connectLoop()
	return nil
}

// SetContext 该函数必须在Connect前使用， 否则可能存在重连问题
func (r *Conn) SetContext(ctx context.Context) *Conn {
	r.ctx, r.cancel = context.WithCancel(ctx)
	return r
}
func (r *Conn) SetUrl(user, password, host string, port int, vhost string) *Conn {
	r.url = fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		user, password, host, port, vhost,
	)
	return r
}
func (r *Conn) Close() {
	r.cancel()
	if r.isConnected {
		r.closeConnect()
	}
}
func (r *Conn) closeConnect() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.connection != nil && !r.connection.IsClosed() {
		r.connection.Close()
	}
	r.init()
}
func (r *Conn) init() *Conn {
	r.m.Lock()
	defer r.m.Unlock()
	r.isConnected = false
	r.notifyConnect = make(chan struct{})
	return r
}
func (r *Conn) connect() error {
	var err error
	r.connection, err = amqp.Dial(r.url)
	if err != nil {
		return err
	}
	r.channel, err = r.connection.Channel()
	if err != nil {
		return err
	}

	r.m.Lock()
	defer r.m.Unlock()
	r.isConnected = true

	// 这个必须在这里重新初始化
	r.notifyClose = make(chan *amqp.Error)
	r.connection.NotifyClose(r.notifyClose)
	close(r.notifyConnect)
	return err
}
func (r *Conn) connectLoop() {
	i := 1
	for {
		if !r.isConnected {
			log.Println("[rabbitMQ]", "test to connect")
			if err := r.connect(); err != nil {
				log.Println("[rabbitMQ]", i, err.Error(), "Failed to connect rabbitMQ. Retrying...")
				i++
				time.Sleep(reconnectDelay)
			} else {
				i = 1
			}
		}
		select {
		case <-r.ctx.Done():
			goto END
		case <-r.notifyClose:
			r.closeConnect()
		}
	}
END:
	fmt.Println("[rabbitMQ]", "close")
}
func (r *Conn) wait() error {
	i := 1
	for {
		if r.isConnected {
			return nil
		}
		select {
		case <-r.ctx.Done():
			goto END
		case <-r.notifyConnect:
		case <-r.notifyClose:
			time.Sleep(reconnectDelay)
		}
		if reconnectMaxTime > 0 && i >= reconnectMaxTime {
			goto END
		}
		i++
	}
END:
	return errors.New("close rabbitMQ")
}

/*
	amqp.ExchangeFanout
	amqp.ExchangeDirect
	amqp.ExchangeTopic
	amqp.ExchangeHeaders

	topic 举例：
	item.# ：能够匹配 item.insert.abc 或者 item.insert
	item.* ：只能匹配 item.insert
*/

func (r *Conn) CreateExchangeAction(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	return r.channel.ExchangeDeclare(
		name,       // name
		kind,       // type
		durable,    // durable
		autoDelete, // auto-deleted
		internal,   // internal true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		noWait,     // no-wait
		args,       // arguments
	)
}
func (r *Conn) CreateQueueAction(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	// 队列不存在创建
	return r.channel.QueueDeclare(
		name,       // name 队列名称 为空时，名称随机
		durable,    // durable 是否持久化
		autoDelete, // delete when unused 是否自动删除
		exclusive,  // exclusive 是否设置排他
		noWait,     // no-wait 是否非阻塞
		args,       // arguments 参数
	)
}

func (r *Conn) CreateExchange(name, kind string) error {
	return r.CreateExchangeAction(
		name,  // name
		kind,  // type
		true,  // durable
		false, // auto-deleted
		false, // internal true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false, // no-wait
		nil,   // arguments
	)
}
func (r *Conn) CreateExchangeBind(name string, exchange string, routingKeys ...string) error {
	for _, k := range routingKeys {
		err := r.channel.ExchangeBind(
			name,     // name
			k,        // routing key
			exchange, // exchange
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Conn) CreateQueue(nameList ...string) error {
	// 队列不存在创建
	for _, name := range nameList {
		_, err := r.CreateQueueAction(
			name,  // name 队列名称 为空时，名称随机
			true,  // durable 是否持久化
			false, // delete when unused 是否自动删除
			false, // exclusive 是否设置排他
			false, // no-wait 是否非阻塞
			nil,   // arguments 参数
		)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Conn) CreateQueueBind(name string, exchange string, routingKeys ...string) error {
	for _, k := range routingKeys {
		err := r.channel.QueueBind(
			name,     // queue name
			k,        // routing key
			exchange, // exchange
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
	publish
*/

func (r *Conn) PublishExchange(name, kind, key string, body []byte) error {
	if err := r.wait(); err != nil {
		return err
	}
	if err := r.CreateExchange(name, kind); err != nil {
		return err
	}
	return r.channel.Publish(
		name,  // exchange
		key,   // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
func (r *Conn) PublishQueue(name string, body []byte) error {
	if err := r.wait(); err != nil {
		return err
	}
	if err := r.CreateQueue(name); err != nil {
		return err
	}
	return r.channel.Publish(
		"",    // exchange
		name,  // name
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}

/*
	subScribe
	go SubscribeExchange
	go SubscribeQueue
*/

func (r *Conn) SubscribeExchange(callback Callback, name, kind string, routingKeys ...string) {
START:
	if err := r.wait(); err != nil {
		return
	}

	log.Println("[subscribe]", "rabbitMQ", kind, name, routingKeys, "start")
	if err := r.CreateExchange(name, kind); err != nil {
		goto START
	}
	q, err := r.CreateQueueAction("", false, true, false, false, nil)
	if err != nil {
		goto START
	}
	if routingKeys == nil || len(routingKeys) <= 0 {
		routingKeys = []string{""}
	}
	if err := r.CreateQueueBind(q.Name, name, routingKeys...); err != nil {
		goto START
	}
	message, err := r.channel.Consume(
		q.Name, // name
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		goto START
	}
	for {
		select {
		case d, msgIsOpen := <-message:
			if !msgIsOpen {
				break
			}
			if err := callback(d); err == nil {
				d.Ack(true)
			}
		case <-r.ctx.Done():
			goto END
		case <-r.notifyClose:
			goto START
		}
	}
END:
	log.Println("[subscribe]", "rabbitMQ", kind, name, routingKeys, "end")
}
func (r *Conn) SubscribeQueue(callback Callback, name string) {
START:
	if err := r.wait(); err != nil {
		return
	}

	log.Println("[subscribe]", "rabbitMQ", "queue", name, "start")
	if err := r.CreateQueue(name); err != nil {
		goto START
	}
	message, err := r.channel.Consume(
		name, // name
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		goto START
	}

	for {
		select {
		case d, msgIsOpen := <-message:
			if !msgIsOpen {
				break
			}
			if err := callback(d); err == nil {
				d.Ack(true)
			}
		case <-r.ctx.Done():
			goto END
		case <-r.notifyClose:
			goto START
		}
	}
END:
	log.Println("[subscribe]", "rabbitMQ", "queue", name, "end")
}
