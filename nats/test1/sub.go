package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i <= 10; i++ {
		dummy := i
		conn.Subscribe("hello", func(msg *nats.Msg) {
			fmt.Printf("消费者[%d]收到：%s\n", dummy, string(msg.Data))
		})
	}
	select {}
}
