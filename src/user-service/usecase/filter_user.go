package usecase

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/user-service/repository"
)

type FilterUsersUseCase interface {
	Execute(ctx context.Context, pagination *record.Pagination, role string) (*record.Pagination, error)
}

type filterUsersUseCaseImpl struct {
	userRepo repository.UserRepo
}

func (u filterUsersUseCaseImpl) Execute(ctx context.Context, pagination *record.Pagination, role string) (*record.Pagination, error) {
	return u.userRepo.FilterUsers(ctx, pagination, role)
}

func NewFilterUsersUseCase(userRepo repository.UserRepo) FilterUsersUseCase {
	return &filterUsersUseCaseImpl{userRepo: userRepo}
}
