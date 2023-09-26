package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response interface {
	ErrCode() int
	ErrMsg() string
}
type authenticationResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
type JsonResult struct {
	Msg string `json:"message"`
}

func main() {
	client := http.Client{} //创建客户端
	req := `{"name"":"admin","password","admin"}`

	var jsonstr = []byte(req) //转换二进制
	buffer := bytes.NewBuffer(jsonstr)
	request, err := http.NewRequest("POST", "https://www.xcuitech.com/user/login", buffer)
	if err != nil {
		fmt.Printf("http.NewRequest%v", err)

	}
	resp, err := client.Do(request) //发送请求
	if err != nil {
		fmt.Printf("client.Do%v", err)

	}
	defer resp.Body.Close()
	var resData JsonResult

	err = json.NewDecoder(resp.Body).Decode(&resData)
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}

	//content, _ := ioutil.ReadAll(resp.Body)

	log.Println(resData)
	//log.Println(string(content))
}
