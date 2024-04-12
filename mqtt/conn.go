package mqttConn

import (
	"context"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

/*
	这里就不介入账号管理的逻辑了默认认为账号已创建
*/

func New() *Conn {
	return (&Conn{}).SetContext(context.Background())
}

// type Callback func(client mqtt.Client, msg mqtt.Message) mqtt.MessageHandler
type Callback func(topic string, payload []byte)

type Conn struct {
	broker      string
	port        int
	topicPrefix string

	clientID string
	username string
	password string

	ctx    context.Context
	cancel context.CancelFunc
	client mqtt.Client

	subscribe sync.Map
}

func (m *Conn) Close() {
	m.cancel()
	m.client.Disconnect(250)
}
func (m *Conn) SetContext(ctx context.Context) *Conn {
	m.ctx, m.cancel = context.WithCancel(ctx)
	return m
}
func (m *Conn) SetBroker(broker string, port int) *Conn {
	m.broker = broker
	m.port = port
	return m
}
func (m *Conn) SetTopicPrefix(prefix string) *Conn {
	m.topicPrefix = prefix
	return m
}
func (m *Conn) SetAccount(clientID, username, password string) *Conn {
	m.clientID = clientID
	m.username = username
	m.password = password
	return m
}
func (m *Conn) Connect() error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", m.broker, m.port))
	opts.SetClientID(m.clientID)
	opts.SetUsername(m.username)
	opts.SetPassword(m.password)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})
	opts.SetPingTimeout(1 * time.Second)
	opts.SetCleanSession(false)
	opts.ConnectRetry = true
	opts.OnConnect = func(client mqtt.Client) {
		log.Println("Connected success!")
		m.SubscribeAll()
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Printf("Connect lost: %v", err)
	}

	m.client = mqtt.NewClient(opts)
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
func (m *Conn) SubscribeAll() {
	m.subscribe.Range(func(k, v interface{}) bool {
		key, okKey := k.(string)
		cb, okCB := v.(Callback)
		if !okKey || !okCB {
			m.subscribe.Delete(k)
			return false
		}
		if err := m.Subscribe(key, cb); err != nil {
			return false
		}
		return true
	})
}
func (m *Conn) Subscribe(topic string, cb Callback) error {
	topic = m.topicPrefix + topic
	if m.client != nil && m.client.IsConnected() {
		if token := m.client.Subscribe(topic, 2, func(client mqtt.Client, message mqtt.Message) {
			cb(message.Topic(), message.Payload())
		}); token.Wait() && token.Error() != nil {
			log.Println("[mqtt]", "subscribe", "topic", topic, "fail")
			return token.Error()
		}
	}
	m.subscribe.LoadOrStore(topic, cb)
	log.Println("[mqtt]", "subscribe", "topic", topic)
	return nil
}
func (m *Conn) UnSubscribe(topic string) error {
	topic = m.topicPrefix + topic
	if _, ok := m.subscribe.LoadAndDelete(topic); !ok {
		return errors.New("subscribe not found")
	}

	if token := m.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Println("[mqtt]", "unsubscribe", "topic", topic, "fail", token.Error())
		return token.Error()
	}
	log.Println("[mqtt]", "unsubscribe", "topic", topic)
	return nil
}
func (m *Conn) Publish(topic string, payload []byte) error {
	topic = m.topicPrefix + topic
	if token := m.client.Publish(topic, 2, false, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
