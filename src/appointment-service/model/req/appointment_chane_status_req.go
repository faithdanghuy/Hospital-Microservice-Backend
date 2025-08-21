package req

type AppointmentChangeStatusReq struct {
	Status string `json:"status" validate:"required,oneof=pending confirmed cancelled"`
}
