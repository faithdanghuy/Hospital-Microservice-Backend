package res

import (
	"time"

	"github.com/Hospital-Microservice/appointment-service/provider"
)

type AppointmentDetailRes struct {
	ID          string            `json:"id"`
	PatientID   string            `json:"patient_id"`
	Patient     *provider.UserRes `json:"patient,omitempty"`
	DoctorID    string            `json:"doctor_id"`
	Doctor      *provider.UserRes `json:"doctor,omitempty"`
	ScheduledAt time.Time         `json:"scheduled_at"`
	Status      string            `json:"status"`
	Note        string            `json:"note,omitempty"`
	ConfirmedAt *time.Time        `json:"confirmed_at,omitempty"`
}
