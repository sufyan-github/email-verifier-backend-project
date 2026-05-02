package validator

import (
	"net"
	"net/smtp"
	"time"

	"email-verifier/pkg/config"
)

func VerifySMTP(email, mxHost string) bool {
	timeout := time.Duration(config.AppConfig.SMTPTimeout) * time.Second

	conn, err := net.DialTimeout("tcp", mxHost+":25", timeout)
	if err != nil {
		return false
	}

	client, err := smtp.NewClient(conn, mxHost)
	if err != nil {
		return false
	}
	defer client.Close()

	client.Hello("example.com")

	from := "verify@example.com"
	if err = client.Mail(from); err != nil {
		return false
	}

	if err = client.Rcpt(email); err != nil {
		return false
	}

	return true
}