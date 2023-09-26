package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect("nats://0.0.0.0:14222")
	if err != nil {
		return
	}

	conn.Subscribe("hello", func(msg *nats.Msg) {
		fmt.Printf("消费者主题[%s]收到：%s\n", "hello", string(msg.Data))
		conn.Publish(msg.Reply, []byte("这是我的回复"))
	})

	select {}
}
