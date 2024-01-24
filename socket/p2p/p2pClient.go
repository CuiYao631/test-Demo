package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 从命令行参数获取目标 IP 地址
	if len(os.Args) < 2 {
		fmt.Println("Usage: client <target_ip>")
		return
	}
	targetIP := os.Args[1]

	// 创建一个 TCP 地址用于连接
	serverAddr, err := net.ResolveTCPAddr("tcp", targetIP+":7777")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}

	// 创建一个 TCP 连接到目标节点
	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to node:", conn.RemoteAddr())

	// 启动 goroutine 读取来自节点的消息
	go readMessages(conn)

	// 创建一个 scanner 从控制台读取用户输入
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// 读取用户输入
		fmt.Print("Enter message to send: ")
		scanner.Scan()
		input := scanner.Text()

		// 将消息发送到节点
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}

func readMessages(conn net.Conn) {
	// 创建一个 reader 读取消息从节点
	reader := bufio.NewReader(conn)

	for {
		// 从节点读取消息
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// 打印收到的消息
		fmt.Print("Received message from node: ", message)
	}
}
