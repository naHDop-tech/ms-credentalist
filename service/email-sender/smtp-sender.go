package email_sender

import (
	"net/smtp"

	"github.com/naHDop-tech/ms-credentalist/utils"
)

type SmtpSender struct {
	config utils.Config
	auth   smtp.Auth
}

func NewSmtpSender(config utils.Config) Sender {
	auth := smtp.PlainAuth("", config.SmtpLogin, config.SmtpPassword, config.SmtpHost)
	return &SmtpSender{
		config: config,
		auth:   auth,
	}
}

func (s *SmtpSender) Sent(from string, to []string, msg []byte) error {
	err := smtp.SendMail(s.config.SmtpHost+":"+s.config.SmtpPort, s.auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}
