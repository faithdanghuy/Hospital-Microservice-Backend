package mapper

import (
	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/model/req"
)

func TransformAppointmentEditReqToEntity(req *req.AppointmentEditReq) *entity.AppointmentEntity {
	return &entity.AppointmentEntity{
		ScheduledAt: *req.ScheduledAt,
		Status:      req.Status,
		Note:        req.Note,
	}
}
