package usecase

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/repository"
)

type ProfileUseCase interface {
	Execute(ctx context.Context, ID string) (*entity.UserEntity, error)
}

type profileUseCaseImpl struct {
	userRepo repository.UserRepo
}

func (l *profileUseCaseImpl) Execute(
	ctx context.Context,
	ID string,
) (*entity.UserEntity, error) {
	user, err := l.userRepo.FindUserByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewProfileUseCase(UserRepo repository.UserRepo) ProfileUseCase {
	return &profileUseCaseImpl{
		userRepo: UserRepo,
	}
}
