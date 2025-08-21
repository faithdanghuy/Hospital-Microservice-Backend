package mapper

import (
	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/hospital-core/pointer"
)

func TransformAppointmentFilterReqToEntity(r *req.AppointmentFilterReq) *entity.AppointmentEntity {
	return &entity.AppointmentEntity{
		PatientID: pointer.String(r.PatientID),
		DoctorID:  pointer.String(r.DoctorID),
		Status:    pointer.String(r.Status),
	}
}
