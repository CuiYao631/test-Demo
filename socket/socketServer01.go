package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
	"log"
)

//ProxyConfig 配置
type MsgConfig struct {
	Type string `json:"type,omitempty"`
	Uid  string `json:"uid,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

//var connMap = make(map[string]*websocket.Conn)

func sendMessage(replyMsg MsgConfig, conn *websocket.Conn, connUid string) {
	msg := replyMsg.Uid + "说:" + replyMsg.Msg
	if connUid == replyMsg.Uid {
		fmt.Println(msg)
		if replyMsg.Type == "login" {
			msg = "你好！我是你的AI助理，有什么可以帮助你的吗？"
		} else {
			msg = "你说：" + replyMsg.Msg
		}

	}

	if err := websocket.Message.Send(conn, msg); err != nil {
		fmt.Println("Can't send")
	}

}

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		var err error
		for {
			var reply string

			if err = websocket.Message.Receive(ws, &reply); err != nil {
				fmt.Println("Can't receive")
				break
			}

			log.Println("接收消息", reply)
			//replyMsg := MsgConfig{}
			//json.Unmarshal([]byte(reply), &replyMsg)
			//
			//if replyMsg.Type == "login" && replyMsg.Uid != "" {
			//	connMap[replyMsg.Uid] = ws
			//	fmt.Println(connMap)
			//}
			//for k, v := range connMap {
			//
			//	go sendMessage(replyMsg, v, k)
			//}

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Static("/", "../public")
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":8732"))
}
