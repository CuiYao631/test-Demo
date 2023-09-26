package main

import (
	"log"
	"math/rand"
	"time"
)

//随机数测试
func main() {
	strs := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	nums := GenerateRandomNumber(strs, 0, len(strs), len(strs)+3)
	log.Println(nums)
}

func GenerateRandomNumber(input []string, start, end, count int) []string {
	if count > len(input) {
		count = len(input)
	}
	if end < start || (end-start) < count {
		return nil
	}
	output := make([]string, 0)
	nums := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn(end-start) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
			output = append(output, input[num])
		}
	}
	log.Println(nums)
	return output
}
