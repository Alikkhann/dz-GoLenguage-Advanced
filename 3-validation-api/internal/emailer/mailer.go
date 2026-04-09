package emailer

import (
	"go/project-3/configs"
	"go/project-3/internal/payload"
	"net/smtp"
	"strings"
	"github.com/jordan-wright/email"
)

func MailerData(JsonStr payload.DataRequest) error {
  smtpConfig := configs.NewConfig()
	e := email.NewEmail()
	e.From = smtpConfig.Email
	e.To = []string{JsonStr.Email}
	e.Subject = "Подтверждение почты"
	e.Text = []byte("Для подтверждения почты перейдите по ссылке http://localhost:8081/verify/" + JsonStr.Hash )
	
	addr := smtpConfig.Address
	host := strings.Split(smtpConfig.Address, ":")[0]
	err := e.Send(addr, smtp.PlainAuth("", smtpConfig.Email, smtpConfig.Password, host))
	if err != nil {
		return err
	}
	return nil
}