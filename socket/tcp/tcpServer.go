package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Create a TCP address to listen on
	addr, err := net.ResolveTCPAddr("tcp", ":7777")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a TCP listener
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP chat server is listening on", listener.Addr())

	for {
		// Accept incoming TCP connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("New connection from %s\n", conn.RemoteAddr())

	// Create a goroutine to read messages from the client
	go readMessages(conn)

	// Create a scanner to read user input from the console
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Read user input
		fmt.Print("Enter message to send: ")
		scanner.Scan()
		input := scanner.Text()

		// Send the message to the client
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}

func readMessages(conn net.Conn) {
	// Create a reader to read messages from the client
	reader := bufio.NewReader(conn)

	for {
		// Read a message from the client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// Print the received message
		fmt.Printf("Received message from %s: %s", conn.RemoteAddr(), message)
	}
}
