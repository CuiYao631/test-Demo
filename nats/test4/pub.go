package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	conn, err := nats.Connect("nats://0.0.0.0:14222")
	if err != nil {
		return
	}

	msg, err := conn.Request("hello", []byte("你好呀，我在等你，请给我回复"), time.Second)
	if err != nil {
		return
	}
	fmt.Println(string(msg.Data))

	select {}
}
