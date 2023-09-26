package main

import "github.com/nats-io/nats.go"

func main() {
	conn, err := nats.Connect("nats://0.0.0.0:14222")
	if err != nil {
		return
	}
	conn.Publish("user", []byte("user"))
	conn.Publish("user.hello", []byte("user.hello"))
	conn.Publish("user.hello.lufei", []byte("user.hello.lufei"))
	conn.Publish("user.hello.lufei.namei", []byte("user.hello.lufei.namei"))
	select {}
}
