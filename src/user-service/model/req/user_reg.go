package req

import "time"

type UserRegReq struct {
	FullName string    `json:"full_name" validate:"required"`
	Email    string    `json:"email" validate:"required"`
	Phone    string    `json:"phone" validate:"required"`
	Avatar   string    `json:"avatar" validate:"omitempty,url"`
	Birthday time.Time `json:"birthday" validate:"required"`
	Role     string    `json:"role" validate:"required,oneof=patient doctor admin"`
}
