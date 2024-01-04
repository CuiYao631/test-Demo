package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

type mail struct {
	user   string
	passwd string
}

//初始化用户名和密码
func New(u string, p string) mail {
	temp := mail{user: u, passwd: p}
	return temp
}

//标题 文本 目标邮箱
func (m mail) Send(title string, text string, toId string) {
	auth := smtp.PlainAuth("", m.user, m.passwd, "smtp.gmail.com")

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.gmail.com",
	}

	conn, err := tls.Dial("tcp", "smtp.gmail.com:465", tlsconfig)
	if err != nil {
		log.Println("Dial")
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, "smtp.gmail.com")
	if err != nil {
		log.Println("NewClient")
		log.Panic(err)
	}

	if err = client.Auth(auth); err != nil {
		log.Println("Auth")
		log.Panic(err)
	}

	if err = client.Mail(m.user); err != nil {
		log.Println("Mail")
		log.Panic(err)
	}

	if err = client.Rcpt(toId); err != nil {
		log.Println("Rcpt")
		log.Panic(err)
	}

	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s", toId, "support@oeonline.com.au", "support@oeonline.com.au", "ceshi", contentType, text)
	//msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", "support@oeonline.com.au", toId, title, text)

	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()
}
func main() {

	foo := New("oetech24@gmail.com", "ckvl zitl trdh tnig")

	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Your Title Here</title>
    <style>
        body {
            display: flex;
            align-items: center;
            justify-content: center;
            height: 100vh;
            margin: 0;
            text-align: center;
            font-family: Arial, sans-serif;
        }

        .container {
            max-width: 400px;
        }

        h1 {
            color: #333;
        }

        p {
            margin: 20px 0;
            color: #666;
        }

        button {
            padding: 10px 20px;
            font-size: 16px;
            background-color: #007BFF;
            color: #fff;
            border: none;
            cursor: pointer;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>OEOnline</h1>
        <p>Click on the link to Login</p>
        <a href="https://your-link-here.com">Click me</a>
    </div>
</body>
</html>
`

	//title:邮件标题		text:邮件内容	told:发送到的邮箱
	foo.Send("ceshi", body, "cuiyao07@gmail.com")
}
