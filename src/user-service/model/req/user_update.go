package req

import "time"

type UserUpdateReq struct {
	FullName string    `json:"full_name" validate:"required"`
	Email    string    `json:"email" validate:"required"`
	Avatar   string    `json:"avatar" validate:"required"`
	Birthday time.Time `json:"birthday" validate:"required"`
	Gender   string    `json:"gender" validate:"required"`
	Address  string    `json:"address" validate:"required"`
}
