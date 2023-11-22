package main

import (
	"fmt"
	"time"
)

// 时间转换
func main() {
	//now := time.Now()
	//
	//// 18:00
	//eighteen := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, now.Location())
	//
	//// 6:00 第二天
	//six := time.Date(now.Year(), now.Month(), now.Day()+1, 6, 0, 0, 0, now.Location())
	//
	//if now.After(eighteen) && now.Before(six) {
	//	fmt.Println("当前时间在 18:00 到第二天 6:00 之间")
	//} else {
	//	fmt.Println("当前时间不在 18:00 到第二天 6:00 之间")
	//}

	timeStamp := time.Date(2022, 5, 20, 13, 14, 0, 0, time.Local).Unix()
	fmt.Println("时间转时间戳：", timeStamp)

	timeLayout := "15:04:05"
	timeStr := time.Unix(18868702633333334, 0).Format(timeLayout)
	fmt.Println(timeStr)
}
