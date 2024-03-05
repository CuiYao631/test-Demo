package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	sysType := runtime.GOOS
	fmt.Println(runtime.GOARCH)

	log.Println(sysType)

	if sysType == "linux" {
		// LINUX系统
		fmt.Println("Linux system")
	}

	if sysType == "windows" {
		// windows系统
		fmt.Println("Windows system")
	}
	if sysType == "darwin" {
		// MAC系统
		fmt.Println("Mac system")
	}

}
