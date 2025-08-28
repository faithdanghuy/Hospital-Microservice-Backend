package usecase

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/repository"
)

type DeleteMedicationUseCase interface {
	Execute(ctx context.Context, id string) error
}

type deleteMedicationUseCaseImpl struct {
	repo repository.PrescriptionRepo
}

func (u *deleteMedicationUseCaseImpl) Execute(ctx context.Context, id string) error {
	return u.repo.DeleteMedication(ctx, id)
}

func NewDeleteMedicationUseCase(repo repository.PrescriptionRepo) DeleteMedicationUseCase {
	return &deleteMedicationUseCaseImpl{repo: repo}
}
