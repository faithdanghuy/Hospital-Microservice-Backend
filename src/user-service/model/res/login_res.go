package res

import "time"

type LoginRes struct {
	ID          string    `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Password    string    `json:"password,omitempty"`
	Avatar      string    `json:"avatar"`
	Birthday    time.Time `json:"birthday"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessToken string    `json:"access_token"`
}
