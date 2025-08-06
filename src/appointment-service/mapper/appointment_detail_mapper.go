package mapper

import (
	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/model/res"
)

func TransformAppointmentEntityToRes(appointmentEntity *entity.AppointmentEntity) *res.AppointmentDetailRes {
	return &res.AppointmentDetailRes{
		ID:          *appointmentEntity.ID,
		PatientID:   *appointmentEntity.PatientID,
		DoctorID:    *appointmentEntity.DoctorID,
		ScheduledAt: appointmentEntity.ScheduledAt,
		Status:      *appointmentEntity.Status,
		Note:        *appointmentEntity.Note,
		ConfirmedAt: appointmentEntity.ConfirmedAt,
	}
}
