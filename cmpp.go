package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var (
	TableTemplate = `{
    "ecName":"{{.EcNames}}",
    "apId":"{{.ApID}}",
    "mobiles":"{{.Mobiles}}",
    "content":"{{.Content}}",
    "sign":"{{.Sign}}",
    "addSerial":"{{.AddSerial}}",
    "mac":"{{.Mac}}"
}`
)

type SmsParams struct {
	EcNames   string `json:"ecName"`
	ApID      string `json:"apId"`
	SecretKey string `json:"secretKey"`
	Mobiles   string `json:"mobiles"`
	Content   string `json:"content"`
	Sign      string `json:"sign"`
	AddSerial string `json:"addSerial"`
	Mac       string `json:"mac"`
}

func makeSmsParams() *SmsParams {
	return &SmsParams{}
}

func (s *SmsParams) GetSmsParams() string {
	s.EcNames = "西安市莲湖区人力资源和社会保障局"
	s.ApID = "lianhu"
	s.SecretKey = "Passwd_13"
	s.Mobiles = "13609287424"
	s.Content = "xiaocui_Test"
	s.Sign = "JwMg8UfeM"
	s.AddSerial = ""
	h := md5.New()
	h.Write([]byte(s.EcNames + s.ApID + s.SecretKey + s.Mobiles + s.Content + s.Sign + s.AddSerial))
	s.Mac = hex.EncodeToString(h.Sum(nil))

	templ, err := template.New("TableTemplate").Parse(TableTemplate)
	if err != nil {
		log.Fatal("create table template error", err)
	}

	sqlCmd := new(strings.Builder)
	err = templ.Execute(sqlCmd, SmsParams{
		EcNames:   "西安市莲湖区人力资源和社会保障局",
		ApID:      "lianhu",
		Mobiles:   "13609287424",
		Content:   "xiaocui_Test",
		Sign:      "JwMg8UfeM",
		AddSerial: "",
		Mac:       hex.EncodeToString(h.Sum(nil)),
	})

	fmt.Println(sqlCmd.String())

	decodeBytes := base64.StdEncoding.EncodeToString([]byte(sqlCmd.String()))

	var jsonstr = []byte(decodeBytes) //转换二进制
	buffer := bytes.NewBuffer(jsonstr)
	request, err := http.NewRequest("POST", "http://112.35.1.155:1992/sms/norsubmit", buffer)
	if err != nil {
		fmt.Printf("http.NewRequest%v", err)
		return "err"
	}

	client := http.Client{}         //创建客户端
	resp, err := client.Do(request) //发送请求
	if err != nil {
		fmt.Printf("client.Do%v", err)
		return "err"
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll%v", err)
		return "err"
	}
	fmt.Println(respBytes)

	return decodeBytes
}

func main() {
	s := makeSmsParams()

	log.Println(s.GetSmsParams())
}
