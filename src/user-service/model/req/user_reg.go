package req

import "time"

type UserRegReq struct {
	FullName string    `json:"full_name" validate:"required"`
	Email    string    `json:"email" validate:"required"`
	Phone    string    `json:"phone" validate:"required"`
	Avatar   string    `json:"avatar" validate:"omitempty,url"`
	Gender   string    `json:"gender" validate:"required,oneof=male female"`
	Address  string    `json:"address" validate:"required"`
	Birthday time.Time `json:"birthday" validate:"required"`
	Role     string    `json:"role" validate:"required,oneof=patient doctor admin"`
}
