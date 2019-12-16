package mail

import (
	"net/smtp"
	)

func MailSend() {
	auth := smtp.PlainAuth("", "alexey@dshawk.info", "5uNo#g06", "mail.hosting.reg.ru")
	smtp.SendMail("smtp.yandex.ru:25", auth, "alexey@dshawk.info", []string{"ork.04@yandex.ru"}, []byte("Test"))
}