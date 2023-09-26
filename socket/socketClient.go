package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

func sendMessages(conn net.Conn, messages <-chan string) {
	for message := range messages {
		_, err := conn.Write([]byte(message))
		if err != nil {
			log.Println("Failed to send message:", err)
			continue
		}
	}
}

func receiveResponses(conn net.Conn, responses chan<- string) {
	reader := bufio.NewReader(conn)
	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read response:", err)
			close(responses)
			return
		}
		responses <- response
	}
}

func main() {
	for i := 0; i < 500; i++ {
		conn, err := net.Dial("tcp", "localhost:8732")
		if err != nil {
			log.Fatal("Failed to connect to server:", err)
		}
		defer conn.Close()

		messages := make(chan string)
		responses := make(chan string)

		// Start goroutines for sending and receiving messages concurrently
		go sendMessages(conn, messages)
		go receiveResponses(conn, responses)

		// Read messages from user and send them to the server
		//go func() {
		//	reader := bufio.NewReader(os.Stdin)
		//	for {
		//		fmt.Print("Enter a message: ")
		//		message, err := reader.ReadString('\n')
		//		if err != nil {
		//			log.Println("Failed to read input:", err)
		//			continue
		//		}
		//		messages <- message
		//	}
		//}()
		messages <- "*CMDR,OM," + strconv.Itoa(i) + ",000000000000,Q0,338,0#\n"

		// Print server responses
		for response := range responses {
			fmt.Println("Server response:", response)
		}
	}

}

//*CMDR,OM,869731053930911,000000000000,Q0,338,0#
//*CMDR,OM,869731053930913,000000000000,L0,0,1,1687315677#
