package req

type AppointmentCreateReq struct {
	PatientID   string `json:"patient_id" validate:"required"`
	DoctorID    string `json:"doctor_id" validate:"required"`
	ScheduledAt string `json:"scheduled_at" validate:"required"`
	Note        string `json:"note,omitempty"`
}
