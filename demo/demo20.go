package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func generateOrderNumber(userID string) string {
	// 生成唯一ID
	orderID := uuid.New().String()

	// 获取当前时间戳
	timestamp := time.Now().Unix()

	// 结合用户ID和时间戳生成订单号
	orderNumber := fmt.Sprintf("%s-%s-%d", userID, orderID, timestamp)

	return orderNumber
}

func main() {
	userID := "user123"
	orderNumber := generateOrderNumber(userID)
	fmt.Println("Generated Order Number:", orderNumber)
}
