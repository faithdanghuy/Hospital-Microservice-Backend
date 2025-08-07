package mapper

import (
	"time"

	"github.com/Hospital-Microservice/hospital-core/pointer"
	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/model/req"
)

func TransformUpdateReqToEntity(id string, req req.UserUpdateReq) *entity.UserEntity {
	return &entity.UserEntity{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(id),
		},
		FullName: pointer.String(req.FullName),
		Email:    pointer.String(req.Email),
		Phone:    pointer.String(req.Phone),
		Password: pointer.String(req.Password),
		Avatar:   pointer.String(req.Avatar),
		Birthday: time.Now(),
	}
}
