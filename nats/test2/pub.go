package main

import "github.com/nats-io/nats.go"

func main() {
	conn, err := nats.Connect("nats://0.0.0.0:14222")
	if err != nil {
		return
	}
	for i := 1; i <= 2; i++ {
		conn.Publish("hello", []byte("你好呀，我是王老五"))
	}
	select {}
}
