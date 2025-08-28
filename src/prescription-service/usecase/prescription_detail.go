package usecase

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/repository"
)

type GetPrescriptionUseCase interface {
	Execute(ctx context.Context, id string) (*entity.PrescriptionEntity, error)
}

type getPrescriptionUseCaseImpl struct {
	repo repository.PrescriptionRepo
}

func NewGetPrescriptionUseCase(repo repository.PrescriptionRepo) GetPrescriptionUseCase {
	return &getPrescriptionUseCaseImpl{repo: repo}
}

func (u *getPrescriptionUseCaseImpl) Execute(ctx context.Context, id string) (*entity.PrescriptionEntity, error) {
	return u.repo.GetPrescription(ctx, id)
}
