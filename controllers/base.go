package controllers

import (
	"github.com/astaxie/beego"
	"krypton-server/utils/mail"
	"krypton-server/utils/session"
)

var (
	mailer         *mail.Mail
	sessionManager *session.SessionManger
)

func init() {
	mailer = mail.NewMail(
		beego.AppConfig.String("mail_host"),
		beego.AppConfig.String("mail_email"),
		beego.AppConfig.String("mail_password"),
	)

	expiredTime, err := beego.AppConfig.Int64("expired_time")
	if err != nil {
		panic(err)
	}

	sessionManager = &session.SessionManger{
		Secret:      beego.AppConfig.String("secret"),
		Issuer:      beego.AppConfig.String("issuer"),
		ExpiredTime: expiredTime,
	}
}
