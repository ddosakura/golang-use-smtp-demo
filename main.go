package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/ddosakura/gklang"
	"github.com/joho/godotenv"
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")

	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err

	// fmt.Println("", user, password, hp[0])
	// fmt.Println(host, user, send_to, string(msg))
	// return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		gklang.Er(err)
	}
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	host := "smtp.qq.com:25"
	to := os.Getenv("TEST_TARGET")

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		Hello World!
		</h3>
		<p>配置来源于环境变量</p>
		</body>
		</html>
		`
	fmt.Println("send email")
	err = SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}
