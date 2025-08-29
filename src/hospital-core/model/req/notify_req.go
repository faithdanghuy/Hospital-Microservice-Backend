package req

type NotificationReq struct {
	ToEmails []string `json:"to_emails"`
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`
}
