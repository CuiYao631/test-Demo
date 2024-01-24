package main

import (
	"fmt"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("无法获取当前用户信息:", err)
		return
	}

	fmt.Println("当前用户的用户名:", currentUser.Username)
}
