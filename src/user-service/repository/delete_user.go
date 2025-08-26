package repository

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
)

func (u *userRepoImpl) SoftDeleteUser(ctx context.Context, ID string) error {
	return u.DB.Executor.WithContext(ctx).
		Where("id = ?", ID).
		Delete(&entity.UserEntity{}).Error
}
