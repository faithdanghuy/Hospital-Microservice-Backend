package req

type AppointmentCreateReq struct {
	PatientID   string  `json:"patient_id" validate:"required"`
	DoctorID    string  `json:"doctor_id" validate:"required"`
	ScheduledAt string  `json:"scheduled_at" validate:"required"`
	Status      string  `json:"status" validate:"required,oneof=pending confirmed cancelled"`
	Note        string  `json:"note,omitempty"`
	ConfirmedAt *string `json:"confirmed_at,omitempty"`
}
