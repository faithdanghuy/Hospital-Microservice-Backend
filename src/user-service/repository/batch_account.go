package repository

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
)

func (r *userRepoImpl) GetUsersByIDs(ctx context.Context, IDs []string) ([]*entity.UserEntity, error) {
	var users []*entity.UserEntity
	if err := r.DB.Executor.WithContext(ctx).
		Where("id IN ?", IDs).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
