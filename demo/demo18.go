package main

import (
	"log"
	"time"
)

// 时间
func main() {
	str := "2023-06-05T16:00:00.000Z"
	parse, err := time.Parse("2006-1-2T15:04:05.000Z", str)
	if err != nil {
		log.Println(err)
		return
	}

	//location, err := time.ParseInLocation("2006/01/02 15:04:05", str, time.Local)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	log.Println(parse)
	//log.Println(location)

}
