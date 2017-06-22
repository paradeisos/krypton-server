package controllers

import (
	"github.com/astaxie/beego"
	"krypton-server/utils/mail"
)

var (
	mailer *mail.Mail
)

func init() {
	mailer = mail.NewMail(
		beego.AppConfig.String("mail_host"),
		beego.AppConfig.String("mail_email"),
		beego.AppConfig.String("mail_password"),
	)
}
