package main

import "github.com/gen2brain/beeep"

func main() {

	err := beeep.Alert("Title", "Message body", "assets/warning.png")
	if err != nil {
		panic(err)
	}
}
