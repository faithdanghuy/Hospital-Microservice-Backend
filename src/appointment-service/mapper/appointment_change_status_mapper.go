package mapper

import (
	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/hospital-core/pointer"
)

func TransformAppointmentChangeStatusReqToEntity(r *req.AppointmentChangeStatusReq) *entity.AppointmentEntity {
	return &entity.AppointmentEntity{
		Status: pointer.String(r.Status),
	}
}
