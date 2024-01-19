package main

import (
	"fmt"
	"net"
	"strings"
)

const discoveryPort = 55555

type DiscoveredUser struct {
	Hostname string
	IP       string
}

func main() {
	// Create a UDP address to listen on
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", discoveryPort))
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

	fmt.Println("User discovery client is listening on", conn.LocalAddr())

	buffer := make([]byte, 1024)

	for {
		// Read data from the connection
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		// Extract hostname and IP address from received data
		data := string(buffer[:n])
		parts := strings.Split(data, "|")
		if len(parts) != 2 {
			fmt.Println("Invalid data format:", data)
			continue
		}

		hostname := parts[0]
		ip := parts[1]

		// Print discovered user
		fmt.Printf("Discovered user: %s at %s\n", hostname, ip)
	}
}
