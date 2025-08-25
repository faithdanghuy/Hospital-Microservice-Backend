package repository

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
)

func (u *userRepoImpl) GetUserDetail(ctx context.Context, ID string) (*entity.UserEntity, error) {
	return u.FindUserByID(ctx, ID)
}
