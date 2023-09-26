package main

import "fmt"

func stupidCode() {
	n := 0
	fmt.Println(1 / n)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	stupidCode()
}
