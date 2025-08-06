package res

import "time"

type AppointmentDetailRes struct {
	ID          string     `json:"id"`
	PatientID   string     `json:"patient_id"`
	DoctorID    string     `json:"doctor_id"`
	ScheduledAt time.Time  `json:"scheduled_at"`
	Status      string     `json:"status"`
	Note        string     `json:"note,omitempty"`
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`
}
