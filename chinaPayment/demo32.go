package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	username := "6d018c04-4fca-4559-8521-edf09c39b621"                             // 设置用户名
	password := "864916b7b153107e1e9f219d9357abcfddb34c71c9b6de3db913c93bec79708b" // 设置密码

	// 将用户名和密码进行组合并转换为字节数组
	authString := fmt.Sprintf("%s:%s", username, password)
	byteAuthString := []byte(authString)

	// 对字节数组进行Base64编码
	encodedAuthString := base64.StdEncoding.EncodeToString(byteAuthString)

	// 构建Authorization header
	authorizationHeader := http.CanonicalHeaderKey("Authorization") + ": Basic " + encodedAuthString

	fmt.Println(authorizationHeader)
}
