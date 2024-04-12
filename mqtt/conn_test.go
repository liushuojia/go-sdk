package mqttConn

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestMQTT(t *testing.T) {
	m := New().SetBroker("127.0.0.1", 1883).
		SetAccount("go_mqtt_client", "emqx_u", "liushuojia")

	m.Subscribe("tt", func(topic string, payload []byte) {
		fmt.Println(topic, ":", string(payload))
	})

	if err := m.Connect(); err != nil {
		log.Fatalln(err.Error())
	}
	defer m.Close()

	m.Subscribe("bb", func(topic string, payload []byte) {
		fmt.Println(topic, ":", string(payload))
	})
	m.Subscribe("cc", func(topic string, payload []byte) {
		fmt.Println(topic, ":", string(payload))
	})

	go func() {
		time.Sleep(time.Second * 3)
		publish(m, "tt")
	}()
	go func() {
		time.Sleep(time.Second * 5)
		publish(m, "bb")
	}()
	time.Sleep(3 * time.Minute)
}

func publish(m *Conn, topic string) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		err := m.Publish(topic, []byte(text))
		fmt.Println(err)
		time.Sleep(time.Second)
	}
}
