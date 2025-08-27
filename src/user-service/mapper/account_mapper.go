package mapper

import (
	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/model/res"
)

func TransformUserEntityToRes(userEntity *entity.UserEntity) *res.LoginRes {
	return &res.LoginRes{
		ID:       *userEntity.ID,
		FullName: *userEntity.FullName,
		Email:    *userEntity.Email,
		Phone:    *userEntity.Phone,
		Avatar:   *userEntity.Avatar,
		Birthday: userEntity.Birthday,
		Role:     *userEntity.Role,
		Gender:   *userEntity.Gender,
		Address:  *userEntity.Address,
	}
}
