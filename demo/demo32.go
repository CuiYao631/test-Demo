package main

import (
	"github.com/deckarep/gosx-notifier"
	"net/http"
	"strings"
	"time"
)

//a slice of string sites that you are interested in watching
var sites []string = []string{
	"http://www.123.com",
	"http://www.google.com",
	"http://www.bing.com"}

func main() {
	ch := make(chan string)

	for _, s := range sites {
		go pinger(ch, s)
	}

	for {
		select {
		case result := <-ch:
			if strings.HasPrefix(result, "-") {
				s := strings.Trim(result, "-")
				showNotification("紧急, 无法访问网站: " + s)
			}
		}
	}
}

func showNotification(message string) {

	note := gosxnotifier.NewNotification(message)
	note.Title = "⚠️"
	note.Sound = gosxnotifier.Default
	//note.Sender = "com.apple.messages"
	note.AppIcon = "/Users/xiaocui/GolandProjects/test-Demo/example.png"
	note.ContentImage = "/Users/xiaocui/GolandProjects/test-Demo/example.png"

	note.Push()
}

//Prefixing a site with a + means it's up, while - means it's down
func pinger(ch chan string, site string) {
	for {
		res, err := http.Get(site)

		if err != nil {
			ch <- "-" + site
		} else {
			if res.StatusCode != 200 {
				ch <- "-" + site
			} else {
				ch <- "+" + site
			}
			res.Body.Close()
		}
		time.Sleep(30 * time.Second)
	}
}
