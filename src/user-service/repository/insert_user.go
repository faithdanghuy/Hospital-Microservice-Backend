package repository

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
)

func (u *userRepoImpl) InsertUser(
	ctx context.Context,
	user *entity.UserEntity,
) error {
	return u.DB.Executor.WithContext(ctx).Create(&user).Error
}
