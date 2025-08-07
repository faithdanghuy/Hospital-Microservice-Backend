package req

import "time"

type UserUpdateReq struct {
	FullName string    `json:"full_name" validate:"required"`
	Email    string    `json:"email" validate:"required"`
	Phone    string    `json:"phone" validate:"required"`
	Password string    `json:"password" validate:"required"`
	Avatar   string    `json:"avatar" validate:"required"`
	Birthday time.Time `json:"birthday" validate:"required"`
}
