package notifier

import (
	"context"
	"fmt"
	"net/smtp"
	"notification/internal/config"
	"notification/internal/domain/model"
)

type Email struct {
	from    string
	auth    smtp.Auth
	address string
}

func NewEmail(cfg config.SMTP) *Email {
	return &Email{
		from:    cfg.OrgEmail,
		auth:    smtp.PlainAuth("", cfg.OrgEmail, cfg.Password, cfg.Host),
		address: cfg.Host + ":" + cfg.Port,
	}
}

func (n *Email) SendMessage(ctx context.Context, message model.EmailMessage) error {
	subject := fmt.Sprintf("Subject: %s\n", message.Subject)
	from := fmt.Sprintf("From: %s\n", n.from)
	to := fmt.Sprintf("To: %s\n", message.ToEmail[0])
	contentType := "Content-Type: text/plain; charset=UTF-8\n"

	msg := []byte(from + to + subject + contentType + "\n" + message.Body)

	err := smtp.SendMail(n.address, n.auth, n.from, message.ToEmail, msg)
	if err != nil {
		return err
	}

	return nil
}
