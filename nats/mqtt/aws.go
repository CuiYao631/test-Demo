package main

import (
	"crypto/tls"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

var fallbackFun MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	cer, err := tls.LoadX509KeyPair(
		"/Users/xiaocui/Downloads/c599ad107d5462af786c50e837bd52e44665b1513104411253a0f4d915180299-certificate.pem.crt",
		"/Users/xiaocui/Downloads/c599ad107d5462af786c50e837bd52e44665b1513104411253a0f4d915180299-private.pem.key")
	//fmt.Println("cer==", cer)
	check(err)

	cid := "basicPubSub"
	host := "au1uri55taesh-ats.iot.ap-southeast-2.amazonaws.com"
	port := 8883

	// AutoReconnect option is true by default
	// CleanSession option is true by default
	// KeepAlive option is 30 seconds by default
	connOpts := MQTT.NewClientOptions()
	// This line is different, we use the constructor function instead of creating the instance ourselves.
	connOpts.SetClientID(cid)
	connOpts.SetMaxReconnectInterval(1 * time.Second)
	connOpts.SetTLSConfig(&tls.Config{Certificates: []tls.Certificate{cer}})
	connOpts.SetDefaultPublishHandler(fallbackFun)
	connOpts.SetBinaryWill("over", []byte("123"), 0, false)

	//path := "/mqtt"

	brokerURL := fmt.Sprintf("tcps://%s:%d", host, port)
	connOpts.AddBroker(brokerURL)

	mqttClient := MQTT.NewClient(connOpts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("[MQTT] Connected")
	//	payload := `{
	//    "state": {
	//        "desired" : {
	//            "color" : { "r" : 11 },
	//            "engine" : "ON"
	//        }
	//    }
	//}`
	//fmt.Println("Sending payload.", payload)
	//
	//for {
	//	if token := mqttClient.Publish("test", 0, false, payload); token.Wait() && token.Error() != nil {
	//		log.Fatalf("failed to send update: %v", token.Error())
	//	}
	//	fmt.Println(payload)
	//	fmt.Println("Sending Successfully.")
	//	time.Sleep(1 * time.Second)
	//}
	subtest(mqttClient)
	subtest2(mqttClient)
	//publish(client)

	//client.Disconnect(250)

	select {}

}

func subtest(client MQTT.Client) {
	topic := "test"
	token := client.Subscribe(topic, 0, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
func subtest2(client MQTT.Client) {
	topic := "123"
	token := client.Subscribe(topic, 0, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
