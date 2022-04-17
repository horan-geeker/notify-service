package services

import (
	"gopkg.in/gomail.v2"
	"log"
	"net/mail"
	"notify-service/config"
	"time"
)

const TENCENT_SMTP_SERVER = "smtp.exmail.qq.com"
const TENCENT_SMTP_PORT = 465

func SendMailByQQ(mailTo string, from string, title string, content string) {
	m := gomail.NewMessage()
	fromMail := mail.Address{from, config.Mail.TENCENT_MAIL_ADDRESS}
	m.SetHeader("From", fromMail.String())
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(TENCENT_SMTP_SERVER, TENCENT_SMTP_PORT, config.Mail.TENCENT_MAIL_ADDRESS, config.Mail.TENCENT_MAIL_PASSWORD)

	begin := time.Now()
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return
	}
	log.Println("send mail success, duration:", time.Since(begin), "to:", mailTo, "title:", title)
}
