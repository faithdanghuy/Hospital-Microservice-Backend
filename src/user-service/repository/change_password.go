package repository

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
)

func (u *userRepoImpl) ChangePassword(ctx context.Context, userID, newPwd string) error {
	return u.DB.Executor.WithContext(ctx).
		Model(&entity.UserEntity{}).
		Where("id = ?", userID).
		Update("password", newPwd).Error
}
