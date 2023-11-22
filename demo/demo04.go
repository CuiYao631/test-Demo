package main

import "log"

//呼吸灯算法
var (
	brightness = 0
	val        = 0
	SUB        = 0
	ADD        = 255
)

func main() {
	for {
		if val >= 255 {
			brightness = SUB
		}

		if val <= 0 {
			brightness = ADD
		}

		if brightness == SUB {
			val -= 1
		} else if brightness == ADD {
			val += 1
		}
		log.Println(val)
	}
}
