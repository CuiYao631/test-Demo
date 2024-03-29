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

	// Get the IP addresses associated with the local machine
	localIPs, err := getLocalIPs()
	if err != nil {
		fmt.Println("Error getting local IP addresses:", err)
		return
	}

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

		// Check if the IP is in the list of local IPs
		if !contains(localIPs, ip) {
			// Print discovered user
			fmt.Printf("Discovered user: %s at %s\n", hostname, ip)
		}
	}
}

// getLocalIPs returns a list of local machine IP addresses
func getLocalIPs() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var localIPs []string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIPs = append(localIPs, ipnet.IP.String())
			}
		}
	}

	return localIPs, nil
}

// contains checks if a string is present in a slice of strings
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
