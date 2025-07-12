package req

type UserLoginReq struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
