package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// 获取本机 IP 地址
	localIP, err := getLocalIP()
	if err != nil {
		fmt.Println("Error getting local IP address:", err)
		return
	}

	fmt.Println("Your IP address is:", localIP)

	// 创建一个 TCP 地址来监听
	addr, err := net.ResolveTCPAddr("tcp", localIP+":7777")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// 创建一个 TCP 监听器
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("P2P chat node is listening on", listener.Addr())

	// 创建一个 map 用于存储已连接的节点
	connectedNodes := make(map[string]net.Conn)

	for {
		// 接受新的连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 获取连接的远程地址
		remoteAddr := conn.RemoteAddr().String()
		fmt.Printf("New connection from %s\n", remoteAddr)

		// 存储连接
		connectedNodes[remoteAddr] = conn

		// 启动 goroutine 处理节点的输入
		go handleMessages(conn, remoteAddr, connectedNodes)
	}
}

func handleMessages(conn net.Conn, remoteAddr string, connectedNodes map[string]net.Conn) {
	defer conn.Close()

	// 创建一个 reader 读取来自节点的消息
	reader := bufio.NewReader(conn)

	for {
		// 读取节点的消息
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Connection with %s closed\n", remoteAddr)
			delete(connectedNodes, remoteAddr)
			break
		}

		// 打印收到的消息
		fmt.Printf("Received message from %s: %s", remoteAddr, message)

		// 向其他连接的节点广播消息
		broadcastMessage(remoteAddr, message, connectedNodes)
	}
}

func broadcastMessage(senderAddr string, message string, connectedNodes map[string]net.Conn) {
	for addr, conn := range connectedNodes {
		if addr != senderAddr {
			go func(conn net.Conn) {
				_, err := conn.Write([]byte(message))
				if err != nil {
					fmt.Println("Error writing to node:", err)
				}
			}(conn)
		}
	}
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("No suitable IP address found")
}
