package mapper

import (
	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/model/res"
)

func TransformUserEntityToFilterRes(userEntity *entity.UserEntity) *res.FilterRes {
	return &res.FilterRes{
		ID:        *userEntity.ID,
		UpdatedAt: *userEntity.UpdatedAt,
		CreatedAt: *userEntity.CreatedAt,
		FullName:  *userEntity.FullName,
		Email:     *userEntity.Email,
		Phone:     *userEntity.Phone,
		Avatar:    *userEntity.Avatar,
		Birthday:  userEntity.Birthday,
		Role:      *userEntity.Role,
		Gender:    *userEntity.Gender,
		Address:   *userEntity.Address,
	}
}
