package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Hospital-Microservice/notify-service/model/req"
)

type EmailSender interface {
	SendEmail(ctx context.Context, to []string, subject, body string) error
}

type SmsSender interface {
	SendSMS(ctx context.Context, to []string, body string) error
}

type NotifyUseCase interface {
	Send(ctx context.Context, r req.NotificationReq) error
	HandleEvent(queue string, body []byte) error
}

type notifyUseCase struct {
	email EmailSender
	sms   SmsSender
}

func NewNotifyUseCase(e EmailSender, s SmsSender) NotifyUseCase {
	return &notifyUseCase{email: e, sms: s}
}

func (u *notifyUseCase) Send(ctx context.Context, r req.NotificationReq) error {
	if len(r.ToEmails) > 0 {
		if err := u.email.SendEmail(ctx, r.ToEmails, r.Subject, r.Body); err != nil {
			return err
		}
	}
	if len(r.ToPhones) > 0 {
		if err := u.sms.SendSMS(ctx, r.ToPhones, r.Body); err != nil {
			return err
		}
	}
	return nil
}

func (u *notifyUseCase) HandleEvent(queue string, body []byte) error {

	fmt.Printf("[x] Received from %s: %s\n", queue, string(body))
	var r req.NotificationReq
	if err := json.Unmarshal(body, &r); err != nil {
		var m map[string]interface{}
		if err2 := json.Unmarshal(body, &m); err2 != nil {
			return fmt.Errorf("invalid event payload: %w", err)
		}
		if b, ok := m["body"].(string); ok {
			r.Body = b
		}
		if s, ok := m["subject"].(string); ok {
			r.Subject = s
		}
		// try recipients
		if to, ok := m["to_emails"].([]interface{}); ok {
			for _, it := range to {
				if s, ok := it.(string); ok {
					r.ToEmails = append(r.ToEmails, s)
				}
			}
		}
		if to, ok := m["to_phones"].([]interface{}); ok {
			for _, it := range to {
				if s, ok := it.(string); ok {
					r.ToPhones = append(r.ToPhones, s)
				}
			}
		}
	}
	return u.Send(context.Background(), r)
}
