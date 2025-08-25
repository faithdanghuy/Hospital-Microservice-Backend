package usecase

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/repository"
)

type UserDetailUseCase interface {
	Execute(ctx context.Context, ID string) (*entity.UserEntity, error)
}

type userDetailUseCaseImpl struct {
	userRepo repository.UserRepo
}

func (u *userDetailUseCaseImpl) Execute(ctx context.Context, ID string) (*entity.UserEntity, error) {
	return u.userRepo.GetUserDetail(ctx, ID)
}

func NewUserDetailUseCase(userRepo repository.UserRepo) UserDetailUseCase {
	return &userDetailUseCaseImpl{
		userRepo: userRepo,
	}
}
