package mail

import (
	"net/smtp"
	"sensei/conf"
)

func SendMail(to string, message string) error {
	conf := conf.Get()

	auth := smtp.PlainAuth("", conf.Mail.User, conf.Mail.Pass, conf.Mail.SmtpHost)
	toArray := []string{to}
	err := smtp.SendMail(conf.Mail.SmtpHost+":"+conf.Mail.SmtpPort, auth, conf.Mail.User, toArray, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
