package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Get the server IP address from command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: client <server_ip>")
		return
	}
	serverIP := os.Args[1]

	// Create a TCP address to connect to
	serverAddr, err := net.ResolveTCPAddr("tcp", serverIP+":7777")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}

	// Create a TCP connection to the server
	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server:", conn.RemoteAddr())

	// Create a goroutine to read messages from the server
	go readMessages(conn)

	// Create a scanner to read user input from the console
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Read user input
		fmt.Print("Enter message to send: ")
		scanner.Scan()
		input := scanner.Text()

		// Send the message to the server
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}

func readMessages(conn net.Conn) {
	// Create a reader to read messages from the server
	reader := bufio.NewReader(conn)

	for {
		// Read a message from the server
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// Print the received message
		fmt.Print("Received message from server: ", message)
	}
}
