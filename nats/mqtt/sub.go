package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func NewTLSConfig() *tls.Config {

	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("/Users/xiaocui/Downloads/AmazonRootCA1.pem")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair(
		"/Users/xiaocui/Downloads/c599ad107d5462af786c50e837bd52e44665b1513104411253a0f4d915180299-certificate.pem.crt",
		"/Users/xiaocui/Downloads/c599ad107d5462af786c50e837bd52e44665b1513104411253a0f4d915180299-private.pem.key")
	if err != nil {
		panic(err)
	}

	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	//fmt.Println(cert.Leaf)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}
}

func main() {
	//tlsconfig := NewTLSConfig()
	var broker = "54.153.156.96"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("mqtt://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("test")
	opts.SetPassword("test")
	//opts.SetTLSConfig(tlsconfig)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetWill("test", "go_mqtt_client lost connection", 1, false)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	//publish(client)

	//client.Disconnect(250)

	select {}
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		//text := fmt.Sprintf("Message %d", i)
		//token := client.Publish("test", 0, false, text)
		//token.Wait()
		//client.
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "/user/kingbonn/ebike/comm/ttface/post"
	token := client.Subscribe(topic, 0, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
