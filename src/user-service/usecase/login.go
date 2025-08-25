package usecase

import (
	"context"
	"errors"

	"github.com/Hospital-Microservice/hospital-core/security"
	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/provider"
	"github.com/Hospital-Microservice/user-service/repository"
)

type LoginUseCase interface {
	Execute(ctx context.Context, phone, pwd string) (*entity.UserEntity, error)
}

type loginUseCaseImpl struct {
	provider *provider.AppProvider
	userRepo repository.UserRepo
}

func (l *loginUseCaseImpl) Execute(
	ctx context.Context,
	phone string,
	pwd string,
) (*entity.UserEntity, error) {
	user, err := l.userRepo.FindUserByPhone(
		ctx,
		entity.UserEntity{Phone: &phone},
	)
	if err != nil {
		return nil, err
	}

	if user == nil || user.Password == nil {
		return nil, errors.New("user not found")
	}

	if !security.VerifyPassword(pwd, *user.Password) {
		return nil, errors.New("password does not match")
	}

	return user, nil
}

func NewLoginUseCase(
	provider *provider.AppProvider,
	UserRepo repository.UserRepo,
) LoginUseCase {
	return &loginUseCaseImpl{
		provider: provider,
		userRepo: UserRepo,
	}
}
