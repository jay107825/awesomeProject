package main

/**
 * @Author: admin
 * @Description: 发送邮件
 * @File: main
 * @Version: 1.0.0
 * @Date: 2024/2/27 22:09
 */
import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
)

func sendEmail(to []string, subject string, body string, smtpServer, smtpPort, authUsername, authPassword string) error {
	// 设置邮件头部
	header := make(map[string]string)
	header["From"] = authUsername
	header["To"] = strings.Join(to, ", ")
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "7bit"

	// 构建邮件内容
	var buffer bytes.Buffer
	for k, v := range header {
		buffer.WriteString(k + ": " + v + "\r\n")
	}
	buffer.WriteString("\r\n")
	buffer.WriteString(body)

	// 发送邮件
	addr := smtpServer + ":" + smtpPort
	auth := smtp.PlainAuth("", authUsername, authPassword, smtpServer)
	err := smtp.SendMail(addr, auth, authUsername, to, []byte(buffer.String()))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("jay mail test")
	to := []string{"mail.163.com"}
	subject := "Subject: Your Email Subject Here\r\n"
	body := "This is the body of the email message."
	smtpServer := "mail.163.com"
	smtpPort := "56" // 或者 "25", 取决于你的SMTP服务器设置
	authUsername := "17798128831@163.com"
	authPassword := "Chai107825.."

	err := sendEmail(to, subject, body, smtpServer, smtpPort, authUsername, authPassword)
	if err != nil {
		fmt.Println("SendMail error:", err)
		return
	}

	fmt.Println("Email Sent!")
}
