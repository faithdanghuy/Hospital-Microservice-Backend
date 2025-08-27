package req

import "time"

type AppointmentEditReq struct {
	PatientID   *string    `json:"patient_id,omitempty"`
	DoctorID    *string    `json:"doctor_id,omitempty"`
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	Status      *string    `json:"status,omitempty"`
	Note        *string    `json:"note,omitempty"`
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`
}
