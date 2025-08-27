package req

import "time"

type AppointmentEditReq struct {
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	Status      *string    `json:"status,omitempty"`
	Note        *string    `json:"note,omitempty"`
}
