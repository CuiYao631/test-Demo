package main

import (
	"github.com/samber/lo"
	"log"
)

func main() {
	stream1 := make(chan int, 42)
	stream2 := make(chan int, 42)
	stream3 := make(chan int, 42)

	all := lo.FanIn(100, stream1, stream2, stream3)
	// <-chan int
	log.Println(all)
}
