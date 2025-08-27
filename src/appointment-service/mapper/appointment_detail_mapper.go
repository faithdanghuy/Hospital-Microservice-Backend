package mapper

import (
	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/model/res"
	"github.com/Hospital-Microservice/appointment-service/provider"
)

func TransformAppointmentEntityToDetailRes(
	appointmentEntity *entity.AppointmentEntity,
	users map[string]provider.UserRes,
) *res.AppointmentDetailRes {

	r := &res.AppointmentDetailRes{
		ID:          *appointmentEntity.ID,
		ScheduledAt: appointmentEntity.ScheduledAt,
		ConfirmedAt: appointmentEntity.ConfirmedAt,
	}

	if appointmentEntity.Status != nil {
		r.Status = *appointmentEntity.Status
	}
	if appointmentEntity.Note != nil {
		r.Note = *appointmentEntity.Note
	}
	if appointmentEntity.PatientID != nil {
		r.PatientID = *appointmentEntity.PatientID
		if u, ok := users[*appointmentEntity.PatientID]; ok {
			r.Patient = &u
		}
	}
	if appointmentEntity.DoctorID != nil {
		r.DoctorID = *appointmentEntity.DoctorID
		if u, ok := users[*appointmentEntity.DoctorID]; ok {
			r.Doctor = &u
		}
	}
	return r
}
