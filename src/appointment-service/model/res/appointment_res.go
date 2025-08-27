package res

import (
	"time"

	"github.com/Hospital-Microservice/appointment-service/provider"
)

type AppointmentRes struct {
	ID          string            `json:"id"`
	ScheduledAt time.Time         `json:"scheduled_at"`
	Status      string            `json:"status"`
	Note        string            `json:"note,omitempty"`
	PatientID   string            `json:"patient_id"`
	Patient     *provider.UserRes `json:"patient,omitempty"`
	DoctorID    string            `json:"doctor_id"`
	Doctor      *provider.UserRes `json:"doctor,omitempty"`
	ConfirmedAt *time.Time        `json:"confirmed_at,omitempty"`
}
