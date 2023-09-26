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

	subjs := []string{"user", "user.*", "user.*.>", "user.*.*"}
	for _, subj := range subjs {
		dummy := subj
		conn.Subscribe(dummy, func(msg *nats.Msg) {
			fmt.Printf("消费者主题[%s]收到：%s\n", dummy, string(msg.Data))
		})
	}

	select {}
}
