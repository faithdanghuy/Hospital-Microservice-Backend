package req

type NotificationReq struct {
	ToEmails []string          `json:"to_emails,omitempty"`
	ToPhones []string          `json:"to_phones,omitempty"`
	Subject  string            `json:"subject,omitempty"`
	Body     string            `json:"body"`
	Meta     map[string]string `json:"meta,omitempty"`
}
