package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"time"
)

/*
奇怪的连接方式，需要在电脑上运行一个程序，然后通过这个程序来连接esp8266
https://gobot.io/documentation/platforms/esp8266
*/

func main() {
	firmataAdaptor := firmata.NewTCPAdaptor("192.168.1.26:3030")
	led := gpio.NewLedDriver(firmataAdaptor, "2")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
