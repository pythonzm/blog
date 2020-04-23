package utils

import (
	"gopkg.in/gomail.v2"
)

func MailNotice(title string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", MailInfo.User)
	m.SetHeader("To", MailInfo.User)
	m.SetHeader("Subject", "评论提醒!")
	m.SetBody("text/html", "<h2>有人评论了文章： "+title+"</h2>")

	d := gomail.NewDialer(MailInfo.Host, MailInfo.Port, MailInfo.User, MailInfo.Pass)

	return d.DialAndSend(m)
}
