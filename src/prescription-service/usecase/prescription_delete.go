package usecase

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/repository"
)

type DeletePrescriptionUseCase interface {
	Execute(ctx context.Context, id string) error
}

type deletePrescriptionUseCaseImpl struct {
	repo repository.PrescriptionRepo
}

func NewDeletePrescriptionUseCase(repo repository.PrescriptionRepo) DeletePrescriptionUseCase {
	return &deletePrescriptionUseCaseImpl{repo: repo}
}

func (u *deletePrescriptionUseCaseImpl) Execute(ctx context.Context, id string) error {
	return u.repo.DeletePrescription(ctx, id)
}
