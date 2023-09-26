package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

const (
	url  = "nats://0.0.0.0:4222"
	subj = "weather"
)

var (
	nc  *nats.Conn
	err error
)

func init() {
	if nc, err = nats.Connect(url); checkErr(err) {
		//
	}
}

func main() {

	startServer(subj, "s1")
	startServer(subj, "s2")
	startServer(subj, "s3")
	//wait for subscribe complete
	time.Sleep(1 * time.Second)

	startClient(subj)

	select {}
}

//send message to server
func startClient(subj string) {
	nc.Publish(subj, []byte("Sun"))
	nc.Publish(subj, []byte("Rain"))
	nc.Publish(subj, []byte("Fog"))
	nc.Publish(subj, []byte("Cloudy"))
}

//receive message
func startServer(subj, name string) {
	go sync(nc, subj, name)
	go async(nc, subj, name)
}

func async(nc *nats.Conn, subj, name string) {
	nc.Subscribe(subj, func(msg *nats.Msg) {
		fmt.Println(name, "Received a message From Async : ", string(msg.Data))
	})
}

func sync(nc *nats.Conn, subj, name string) {
	subscription, err := nc.SubscribeSync(subj)
	checkErr(err)
	if msg, err := subscription.NextMsg(10 * time.Second); checkErr(err) {
		if msg != nil {
			fmt.Println(name, "Received a message From Sync : ", string(msg.Data))
		}
	} else {
		//handle timeout
	}

}

func checkErr(err error) bool {
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
