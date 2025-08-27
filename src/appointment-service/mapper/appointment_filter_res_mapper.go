package mapper

import (
	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/model/res"
	"github.com/Hospital-Microservice/appointment-service/provider"
)

func TransformAppointmentEntitiesToRes(
	appointments []*entity.AppointmentEntity,
	users map[string]provider.UserRes,
) []*res.AppointmentRes {

	out := make([]*res.AppointmentRes, 0, len(appointments))
	for _, appt := range appointments {
		r := &res.AppointmentRes{
			ID:          *appt.ID,
			ScheduledAt: appt.ScheduledAt,
		}
		if appt.Status != nil {
			r.Status = *appt.Status
		}
		if appt.Note != nil {
			r.Note = *appt.Note
		}
		if appt.ConfirmedAt != nil {
			r.ConfirmedAt = appt.ConfirmedAt
		}
		if appt.PatientID != nil {
			r.PatientID = *appt.PatientID
			if u, ok := users[*appt.PatientID]; ok {
				r.Patient = &u
			}
		}
		if appt.DoctorID != nil {
			r.DoctorID = *appt.DoctorID
			if u, ok := users[*appt.DoctorID]; ok {
				r.Doctor = &u
			}
		}
		out = append(out, r)
	}
	return out
}
