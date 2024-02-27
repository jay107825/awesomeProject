package main

/**
 * @Author: admin
 * @Description:
 * @File: 163邮箱
 * @Version: 1.0.0
 * @Date: 2024/2/27 22:32
 */

import (
	"fmt"
	"net/smtp"
	"strings"
)

func sendEmailWith163(to []string, subject string, body string, smtpServer, smtpPort, authUsername, authPassword string) error {
	// 设置邮件头信息
	message := []byte("To: " + strings.Join(to, ", ") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// 发送邮件
	addr := smtpServer + ":" + smtpPort
	auth := smtp.PlainAuth("", authUsername, authPassword, smtpServer)
	err := smtp.SendMail(addr, auth, authUsername, to, message)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	to := []string{"recipient@example.com"}                            // 收件人邮箱地址
	subject := "Hello from 163 SMTP"                                   // 邮件主题
	body := "This is a test email sent from 163 SMTP server using Go." // 邮件正文

	smtpServer := "smtp.163.com"          // 163 SMTP服务器地址
	smtpPort := "465"                     // 163 SMTP服务器端口，使用SSL时通常为465，使用TLS时通常为587
	authUsername := "@163.com"            // 你的163邮箱地址
	authPassword := "your-smtp-auth-code" // 你的163邮箱SMTP授权码，不是邮箱密码

	// 发送邮件
	err := sendEmailWith163(to, subject, body, smtpServer, smtpPort, authUsername, authPassword)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
