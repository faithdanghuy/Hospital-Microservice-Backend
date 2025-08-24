package mapper

import (
	"github.com/Hospital-Microservice/hospital-core/pointer"
	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/model/req"
	"github.com/google/uuid"
)

func TransformRegReqToEntity(req req.UserRegReq) *entity.UserEntity {
	return &entity.UserEntity{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		FullName: pointer.String(req.FullName),
		Email:    pointer.String(req.Email),
		Phone:    pointer.String(req.Phone),
		Avatar:   pointer.String(req.Avatar),
		Role:     pointer.String(req.Role),
		Birthday: req.Birthday,
	}
}
