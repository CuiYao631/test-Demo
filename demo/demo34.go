package main

import (
	"github.com/go-toast/toast"
	"log"
)

func main() {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "大佬",
		Message: "只是一个通知",
		Icon:    "C:\\path\\to\\your\\logo.png", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "按钮1", "https://www.google.com/"},
			{"protocol", "按钮2", "https://github.com/"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Println(err)
	}
}
