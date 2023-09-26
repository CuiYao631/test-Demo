package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numClients := 100

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientNum int) {
			defer wg.Done()

			// 建立TCP连接
			conn, err := net.Dial("tcp", "localhost:8732")
			if err != nil {
				fmt.Printf("无法建立连接：%v\n", err)
				return
			}
			defer conn.Close()
			//所以每次随机数都是随机的
			rand.Seed(time.Now().UnixNano())
			num := rand.Intn(1000)
			// 发送消息
			message := "*CMDR,OM," + strconv.Itoa(num) + ",000000000000,Q0,338,0#\n"
			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Printf("发送消息时出错：%v\n", err)
				return
			}
			select {}
			//// 接收服务器的响应
			//buffer := make([]byte, 1024)
			//numBytes, err := conn.Read(buffer)
			//if err != nil {
			//	fmt.Printf("接收响应时出错：%v\n", err)
			//	return
			//}
			//
			//response := string(buffer[:numBytes])
			//fmt.Printf("Client %d 收到服务器的响应：%s\n", clientNum, response)
		}(i + 1)
	}

	wg.Wait()
	fmt.Println("所有客户端连接已关闭")
}
