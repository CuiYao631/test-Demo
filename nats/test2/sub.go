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
	for i := 1; i <= 4; i++ {
		dummy := i
		conn.QueueSubscribe("hello", "go_queue", func(msg *nats.Msg) {
			fmt.Printf("消费者[%d]收到：%s\n", dummy, string(msg.Data))
		})
	}
	select {}
}
