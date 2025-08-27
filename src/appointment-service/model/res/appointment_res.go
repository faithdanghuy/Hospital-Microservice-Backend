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
	Patient     *provider.UserRes `json:"patient,omitempty"`
	Doctor      *provider.UserRes `json:"doctor,omitempty"`
}
