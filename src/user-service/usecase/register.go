package usecase

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/hospital-core/security"
	"go.uber.org/zap"

	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/repository"
)

type RegisterUseCase interface {
	Execute(ctx context.Context, user *entity.UserEntity) error
}

type registerUseCaseImpl struct {
	UserRepo repository.UserRepo
}

func (r registerUseCaseImpl) Execute(ctx context.Context, user *entity.UserEntity) error {
	hashPwd, err := security.HashPassword(*user.Password)
	if err != nil {
		log.Error("failed to hash password", zap.Error(err))
	}
	user.Password = &hashPwd

	if err := r.UserRepo.InsertUser(ctx, user); err != nil {
		log.Error("failed to register user", zap.Error(err))
		return err
	}
	return nil
}

func NewRegisterUseCase(UserRepo repository.UserRepo) RegisterUseCase {
	return &registerUseCaseImpl{
		UserRepo: UserRepo,
	}
}
