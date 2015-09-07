package main

import (
	"fmt"
	"net/smtp"
	"strings"
	// "time"

	"github.com/cuixin/csv4g"
)

/*
 *  user : example@example.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */

type Test struct {
	Id    int
	Name  string
	Email string
}

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	user := "service@qq.net"
	password := "xxx"
	host := "smtp.qq.net:25"

	subject := "qq"

	body := `
    <html>
    <body>
    <p>hello world</p>
    </body>
    </html>
    `

	csv, err := csv4g.New("./email.csv", ',', Test{}, 1)
	if err != nil {
		fmt.Errorf("Error %v\n", err)
		return
	}
	for i := 0; i < csv.LineLen; i++ {
		tt := &Test{}
		err = csv.Parse(tt)
		if err != nil {
			fmt.Printf("Error on parse %v\n", err)
			return
		}
		fmt.Println(tt)
		fmt.Printf("send email to %s", tt.Email)
		to := tt.Email
		err := SendMail(user, password, host, to, subject, body, "html")
		if err != nil {
			fmt.Println("send mail error!")
			fmt.Println(err)
		} else {
			fmt.Println("send mail success!")
		}
		// timersleep := time.NewTimer(time.Second * 3)
		// //此处在等待channel中的信号，执行此段代码时会阻塞三秒
		// <-timersleep.C
		// fmt.Println("3s expired")
	}
}
