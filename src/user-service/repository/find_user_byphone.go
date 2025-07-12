package repository

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
)

func (u *userRepoImpl) FindUserByPhone(
	ctx context.Context,
	user entity.UserEntity,
) (*entity.UserEntity, error) {
	var userEntity entity.UserEntity

	var findQuery = u.DB.Executor.WithContext(ctx).
		Where("phone = ?", user.Phone)

	if err := findQuery.Find(&userEntity).Error; err != nil {
		return nil, err
	}

	return &userEntity, nil
}
