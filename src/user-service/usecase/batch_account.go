package usecase

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/repository"
)

type BatchUserDetailUseCase interface {
	Execute(ctx context.Context, IDs []string) ([]*entity.UserEntity, error)
}

type batchUserDetailUseCaseImpl struct {
	userRepo repository.UserRepo
}

func (u *batchUserDetailUseCaseImpl) Execute(ctx context.Context, IDs []string) ([]*entity.UserEntity, error) {
	if len(IDs) == 0 {
		return []*entity.UserEntity{}, nil
	}
	return u.userRepo.GetUsersByIDs(ctx, IDs)
}

func NewBatchUserDetailUseCase(userRepo repository.UserRepo) BatchUserDetailUseCase {
	return &batchUserDetailUseCaseImpl{
		userRepo: userRepo,
	}
}
