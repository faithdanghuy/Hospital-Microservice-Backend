package provider

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

type SMTPEmailClient struct {
	host    string
	port    int
	user    string
	pass    string
	timeout time.Duration
}

func NewSMTPEmailClient(host string, port int, user, pass string, timeout time.Duration) *SMTPEmailClient {
	return &SMTPEmailClient{host: host, port: port, user: user, pass: pass, timeout: timeout}
}

// provider/smtp_email_client.go
func (s *SMTPEmailClient) SendEmail(ctx context.Context, to []string, subject, body string) error {
	if len(to) == 0 {
		return nil
	}
	server := fmt.Sprintf("%s:%d", s.host, s.port)

	// Only use AUTH when creds are provided (avoids "unencrypted connection" with MailHog).
	var auth smtp.Auth
	if s.user != "" && s.pass != "" {
		auth = smtp.PlainAuth("", s.user, s.pass, s.host)
	}

	// Pick a from address. MailHog doesn’t care, but some MTAs do.
	from := s.user
	if from == "" {
		from = "no-reply@" + s.host
	}

	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + strings.Join(to, ",") + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
			"\r\n" + body,
	)

	ch := make(chan error, 1)
	go func() {
		ch <- smtp.SendMail(server, auth, from, to, msg)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(s.timeout):
		return fmt.Errorf("send email timeout")
	case <-ctx.Done():
		return ctx.Err()
	}
}
