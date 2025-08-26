package usecase

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/hospital-core/security"
	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/repository"
	"go.uber.org/zap"
)

type EditUserUseCase interface {
	Execute(ctx context.Context, user *entity.UserEntity) error
}

type editUserUseCaseImpl struct {
	userRepo repository.UserRepo
}

func (u editUserUseCaseImpl) Execute(ctx context.Context, user *entity.UserEntity) error {
	// hash password if provided
	if user.Password != nil {
		hashPwd, err := security.HashPassword(*user.Password)
		if err != nil {
			log.Error("failed to hash password", zap.Error(err))
			return err
		}
		user.Password = &hashPwd
	}

	if err := u.userRepo.UpdateUser(ctx, *user); err != nil {
		log.Error("failed to update user", zap.Error(err))
		return err
	}
	return nil
}

func NewEditUserUseCase(UserRepo repository.UserRepo) EditUserUseCase {
	return &editUserUseCaseImpl{
		userRepo: UserRepo,
	}
}
