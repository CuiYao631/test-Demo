package main

// 模版
import (
	"fmt"
	"log"
	"strings"
	"text/template"
)

type SmsParamsp struct {
	EcNames   string `json:"ecName"`
	ApID      string `json:"apId"`
	Mobiles   string `json:"mobiles"`
	Content   string `json:"content"`
	Sign      string `json:"sign"`
	AddSerial string `json:"addSerial"`
	Mac       string `json:"mac"`
}

var (
	TableTemplates = `{
    "ecName":"{{.EcNames}}",
    "apId":"{{.ApID}}",
    "mobiles":"{{.Mobiles}}",
    "content":"{{.Content}}",
    "sign":"{{.Sign}}",
    "addSerial":"{{.AddSerial}}",
    "mac":"{{.Mac}}"
}`
)

func main() {
	templ, err := template.New("TableTemplate").Parse(TableTemplates)
	if err != nil {
		log.Fatal("create table template error", err)
	}

	sqlCmd := new(strings.Builder)
	err = templ.Execute(sqlCmd, SmsParamsp{
		EcNames:   "456",
		ApID:      "",
		Mobiles:   "",
		Content:   "",
		Sign:      "",
		AddSerial: "",
		Mac:       "",
	})

	fmt.Println(sqlCmd.String())

}
