package mapper

import (
	"time"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/hospital-core/pointer"
	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/google/uuid"
)

func TransformAppointmentCreateReqToEntity(req req.AppointmentCreateReq) *entity.AppointmentEntity {
	var scheduledAt time.Time
	if req.ScheduledAt != "" {
		if t, err := time.Parse(time.RFC3339, req.ScheduledAt); err == nil {
			scheduledAt = t
		}
	}
	return &entity.AppointmentEntity{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		PatientID:   pointer.String(req.PatientID),
		DoctorID:    pointer.String(req.DoctorID),
		ScheduledAt: scheduledAt,
		Note:        pointer.String(req.Note),
		ConfirmedAt: nil,
	}
}
