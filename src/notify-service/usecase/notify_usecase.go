package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Hospital-Microservice/notify-service/entity"
	"github.com/Hospital-Microservice/notify-service/model/req"
)

type EmailSender interface {
	SendEmail(ctx context.Context, to []string, subject, body string) error
}

type SmsSender interface {
	SendSMS(ctx context.Context, to []string, body string) error
}

type NotificationRepo interface {
	Save(ctx context.Context, n *entity.NotificationEntity) error
	ListByUser(ctx context.Context, userID string, limit, offset int64) ([]entity.NotificationEntity, error)
	MarkRead(ctx context.Context, id string) error
}

type NotifyUseCase interface {
	Send(ctx context.Context, r req.NotificationReq) error
	HandleEvent(queue string, body []byte) error
	ListByUser(ctx context.Context, userID string, limit, offset int64) ([]entity.NotificationEntity, error)
	MarkAsRead(ctx context.Context, id string) error
}

type notifyUseCase struct {
	email EmailSender
	sms   SmsSender
	repo  NotificationRepo
}

func NewNotifyUseCase(e EmailSender, s SmsSender, r NotificationRepo) NotifyUseCase {
	return &notifyUseCase{email: e, sms: s, repo: r}
}

func (u *notifyUseCase) Send(ctx context.Context, r req.NotificationReq) error {
	fmt.Printf("[Send] Subject: %s, Body: %s, Meta: %+v\n", r.Subject, r.Body, r.Meta)
	if r.Meta != nil {
		if uid, ok := r.Meta["user_id"]; ok && uid != "" {
			n := &entity.NotificationEntity{
				UserID:    uid,
				Title:     r.Subject,
				Body:      r.Body,
				Channel:   "mixed",
				Payload:   map[string]interface{}(map[string]interface{}{}),
				IsRead:    false,
				CreatedAt: time.Now().UTC(),
			}
			_ = u.repo.Save(ctx, n)
		}
		if raw, ok := r.Meta["user_ids"]; ok && raw != "" {
			var ids []string
			if err := json.Unmarshal([]byte(raw), &ids); err != nil {
				for _, part := range splitAndTrim(raw) {
					if part != "" {
						ids = append(ids, part)
					}
				}
			}
			meta := make(map[string]any, len(r.Meta))
			for k, v := range r.Meta {
				meta[k] = v
			}
			for _, id := range ids {
				n := &entity.NotificationEntity{
					UserID:    id,
					Title:     r.Subject,
					Body:      r.Body,
					Channel:   "mixed",
					Meta:      meta,
					Payload:   map[string]interface{}(map[string]interface{}{}),
					IsRead:    false,
					CreatedAt: time.Now().UTC(),
				}
				_ = u.repo.Save(ctx, n)
			}
		}
	}

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
	var rawMap map[string]interface{}
	if err := json.Unmarshal(body, &r); err != nil {
		if err2 := json.Unmarshal(body, &rawMap); err2 != nil {
			return fmt.Errorf("invalid event payload: %w", err)
		}
		if b, ok := rawMap["body"].(string); ok {
			r.Body = b
		}
		if s, ok := rawMap["subject"].(string); ok {
			r.Subject = s
		}
		if to, ok := rawMap["to_emails"].([]interface{}); ok {
			for _, it := range to {
				if s, ok := it.(string); ok {
					r.ToEmails = append(r.ToEmails, s)
				}
			}
		}
		if to, ok := rawMap["to_phones"].([]interface{}); ok {
			for _, it := range to {
				if s, ok := it.(string); ok {
					r.ToPhones = append(r.ToPhones, s)
				}
			}
		}
		if m, ok := rawMap["meta"].(map[string]interface{}); ok {
			r.Meta = make(map[string]string)
			for k, v := range m {
				if s, ok := v.(string); ok {
					r.Meta[k] = s
				} else {
					b, _ := json.Marshal(v)
					r.Meta[k] = string(b)
				}
			}
		}
	}

	if r.Meta != nil {
		// single user
		if uid, ok := r.Meta["user_id"]; ok && uid != "" {
			n := &entity.NotificationEntity{
				UserID:    uid,
				Title:     r.Subject,
				Body:      r.Body,
				Channel:   "mixed",
				Payload:   map[string]interface{}(map[string]interface{}{}),
				IsRead:    false,
				CreatedAt: time.Now().UTC(),
			}
			_ = u.repo.Save(context.Background(), n)
		}
		if raw, ok := r.Meta["user_ids"]; ok && raw != "" {
			var ids []string
			if err := json.Unmarshal([]byte(raw), &ids); err != nil {
				for _, part := range splitAndTrim(raw) {
					ids = append(ids, part)
				}
			}
			for _, id := range ids {
				n := &entity.NotificationEntity{
					UserID:    id,
					Title:     r.Subject,
					Body:      r.Body,
					Channel:   "mixed",
					Payload:   map[string]interface{}(map[string]interface{}{}),
					IsRead:    false,
					CreatedAt: time.Now().UTC(),
				}
				_ = u.repo.Save(context.Background(), n)
			}
		}
	}

	return u.Send(context.Background(), r)
}

func splitAndTrim(s string) []string {
	out := []string{}
	for _, p := range strings.Split(s, ",") {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func (u *notifyUseCase) ListByUser(ctx context.Context, userID string, limit, offset int64) ([]entity.NotificationEntity, error) {
	return u.repo.ListByUser(ctx, userID, limit, offset)
}

func (u *notifyUseCase) MarkAsRead(ctx context.Context, id string) error {
	return u.repo.MarkRead(ctx, id)
}
