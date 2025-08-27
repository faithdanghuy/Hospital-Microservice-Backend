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
	return &entity.AppointmentEntity{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		PatientID:   pointer.String(req.PatientID),
		DoctorID:    pointer.String(req.DoctorID),
		ScheduledAt: time.Now(),
		Note:        pointer.String(req.Note),
		ConfirmedAt: nil,
	}
}
