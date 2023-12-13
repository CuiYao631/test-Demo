package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// 将sender@example.com替换为“发件人”地址。 此地址必须与Amazon SES进行验证。
	Sender = "oetech24@gmail.com"

	// 将recipient@example.com替换为“收件人”地址。如果你的账户
	Recipient = "oetech24@gmail.com"

	// 电子邮件的主题。
	Subject = "Amazon SES Test (AWS SDK for Go)"

	// 电子邮件的HTML正文。
	HtmlBody = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	// 非html电子邮件客户端的收件人的电子邮件主体。
	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

	// 电子邮件的字符编码。
	CharSet = "UTF-8"
)

func main() {

	// 设置访问密钥和密钥 ID
	accessKeyID := "AKIA4NIHJX534PCETVMA"
	secretAccessKey := "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"

	// 创建 AWS 会话
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-2"),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	// 创建会话错误
	if err != nil {
		fmt.Println("Error creating session:", err)
		os.Exit(1)
	}
	fmt.Println("Session created successfully")
	// 组装电子邮件。.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
	}

	svc := ses.New(sess)
	// 尝试发送电子邮件。
	fmt.Println("Attempting to send email to " + Recipient)
	result, err := svc.SendEmail(input)

	// 如果出现错误，则显示错误消息。
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected: // 消息被拒绝
				fmt.Println("消息被拒绝")
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException: // 未验证域名
				fmt.Println("未验证域名")
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException: // 配置集不存在
				fmt.Println("配置集不存在")
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println("发送失败：", aerr.Error())
			}
		} else {
			// 打印错误的消息。
			fmt.Println("发送失败：", err.Error())
		}

		return
	}

	fmt.Println("Email Sent to address: " + Recipient)
	fmt.Println(result)
}
