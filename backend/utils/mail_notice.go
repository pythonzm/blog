package utils

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func MailNotice(title string, articleID int) error {
	articleUrl := fmt.Sprintf("https://www.poorops.com/#/articles/?id=%d", articleID)
	aTag := fmt.Sprintf("<a href=" + articleUrl + ">" + articleUrl + "</a>")
	htmlMsg := "<h2>有人评论了文章： " + title + "，快去看看吧：" + aTag + "</h2>"

	m := gomail.NewMessage()
	m.SetHeader("From", MailInfo.User)
	m.SetHeader("To", MailInfo.User)
	m.SetHeader("Subject", "评论提醒!")
	m.SetBody("text/html", htmlMsg)

	d := gomail.NewDialer(MailInfo.Host, MailInfo.Port, MailInfo.User, MailInfo.Pass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
