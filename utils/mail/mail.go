package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	Host     string `json:"host"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewMail(host, email, password string) *Mail {
	return &Mail{
		Host:     host,
		Email:    email,
		Password: password,
	}
}

/*
   user : example@example.com login smtp server user
   password: xxxxx login smtp server password
   host: smtp.example.com:port   smtp.163.com:25
   to: example@example.com;example1@163.com;example2@sina.com.cn;...
   subject:The subject of mail
   body: The content of mail
   mailtyoe: mail type html or text
*/
func (m *Mail) SendMail(to, subject, body, mailtype string) error {
	hp := strings.Split(m.Host, ":")
	auth := smtp.PlainAuth("", m.Email, m.Password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + m.Email + "<" + m.Email + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(m.Host, auth, m.Email, send_to, msg)
	return err
}

func (m *Mail) SendRegisterMail(to, content string) {
	subject := "Active email from Krypton"

	body := `
    <html>
    <body>
    <h3>
        Welcome to register our APP click this to active your account
        ` + content +
		`
    </h3>
    </body>
    </html>
    `
	err := m.SendMail(to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
}
