package mapper

import (
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
		Avatar:   pointer.String(req.Avatar),
		Birthday: req.Birthday,
		Gender:   pointer.String(req.Gender),
		Address:  pointer.String(req.Address),
	}
}
