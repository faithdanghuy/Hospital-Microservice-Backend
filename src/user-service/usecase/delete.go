package usecase

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/user-service/repository"
	"go.uber.org/zap"
)

type DeleteUserUseCase interface {
	Execute(ctx context.Context, userID string) error
}

type deleteUserUseCaseImpl struct {
	userRepo repository.UserRepo
}

func (d deleteUserUseCaseImpl) Execute(ctx context.Context, userID string) error {
	if err := d.userRepo.SoftDeleteUser(ctx, userID); err != nil {
		log.Error("Failed To Delete User", zap.Error(err))
		return err
	}
	return nil
}

func NewDeleteUserUseCase(userRepo repository.UserRepo) DeleteUserUseCase {
	return &deleteUserUseCaseImpl{
		userRepo: userRepo,
	}
}
