package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Create a UDP address to listen on
	addr, err := net.ResolveUDPAddr("udp", ":55555")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a UDP listener
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("User discovery server is listening on", conn.LocalAddr())

	buffer := make([]byte, 1024)

	for {
		// Read data from the connection
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		// Extract username from received data
		username := strings.TrimSpace(string(buffer[:n]))

		// Print discovered user
		fmt.Printf("Discovered user: %s at %s\n", username, remoteAddr)
	}
}
