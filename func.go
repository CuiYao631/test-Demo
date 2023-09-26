package main

import "fmt"

func A() {

	fmt.Println("A")
}

func Use(f func()) {
	f()
}

func CreateApproval(f func()) {

	f()
}

func main() {
	CreateApproval(A)
}
