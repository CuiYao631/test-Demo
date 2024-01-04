package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"log"
)

func getFirstLetter(username string) string {
	// 将汉字转换为拼音
	p := pinyin.NewArgs()
	result := pinyin.Pinyin(username, p)

	// 提取每个拼音的第一个字母
	log.Println(result)
	firstLetter := result[0][0]

	return firstLetter
}

func main() {
	username := "zzz"
	firstLetter := getFirstLetter(username)
	fmt.Printf("First letter of %s: %s\n", username, firstLetter)
}
