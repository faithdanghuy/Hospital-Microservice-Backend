package usecase

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/repository"
	"go.uber.org/zap"
)

type UpdateUseCase interface {
	Execute(ctx context.Context, user *entity.UserEntity) error
}

type updateUseCaseImpl struct {
	userRepo repository.UserRepo
}

func (u updateUseCaseImpl) Execute(ctx context.Context, user *entity.UserEntity) error {
	if err := u.userRepo.UpdateUser(ctx, *user); err != nil {
		log.Error("failed to update user", zap.Error(err))
		return err
	}
	return nil
}

func NewUpdateUseCase(UserRepo repository.UserRepo) UpdateUseCase {
	return &updateUseCaseImpl{
		userRepo: UserRepo,
	}
}
