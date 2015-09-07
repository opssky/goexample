package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"

	"github.com/cuixin/csv4g"
)

type Test struct {
	Id    int
	Name  string
	Email string
}

// SSL/TLS Email Example

func main() {
	from := mail.Address{"qq", "service@qq.net"}

	// Connect to the SMTP Server
	servername := "smtp.qq.net:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", "service@qq.net", "xxx", host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

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

		to := mail.Address{"", tt.Email}
		subj := "qq"
		body := `
QQ

        `
		fmt.Println(tt)
		fmt.Printf("send email to %s\n", to.Address)

		// Setup headers
		headers := make(map[string]string)
		headers["From"] = from.String()
		headers["To"] = to.String()
		headers["Subject"] = subj
		headers["Content-Type"] = "text/plain; charset=UTF-8"

		// Setup message
		message := ""
		for k, v := range headers {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		message += "\r\n" + body

		// Here is the key, you need to call tls.Dial instead of smtp.Dial
		// for smtp servers running on 465 that require an ssl connection
		// from the very beginning (no starttls)
		conn, err := tls.Dial("tcp", servername, tlsconfig)
		if err != nil {
			log.Panic(err)
		}

		c, err := smtp.NewClient(conn, host)
		if err != nil {
			log.Panic(err)
		}

		// Auth
		if err = c.Auth(auth); err != nil {
			log.Panic(err)
		}

		// To && From
		if err = c.Mail(from.Address); err != nil {
			log.Panic(err)
		}

		if err = c.Rcpt(to.Address); err != nil {
			log.Panic(err)
		}

		// Data
		w, err := c.Data()
		if err != nil {
			log.Panic(err)
		}

		_, err = w.Write([]byte(message))
		if err != nil {
			log.Panic(err)
		}

		err = w.Close()
		if err != nil {
			log.Panic(err)
		}

		c.Quit()
	}

}
