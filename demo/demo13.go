package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

// 定时任务
func main() {
	i := 0
	c := cron.New()
	spec := "*/1 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})
	c.Start()

	defer c.Stop()
	t1 := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-t1.C:
			print5()
		}
	}

}
func print10() {
	log.Println("Run 10s gap")
}

func print5() {
	log.Println("Run 5s gap")
}
