package usecase

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/repository"
)

type UpdateMedicationUseCase interface {
	Execute(ctx context.Context, med *entity.MedicationEntity) error
}

type updateMedicationUseCaseImpl struct {
	repo repository.PrescriptionRepo
}

func (u *updateMedicationUseCaseImpl) Execute(ctx context.Context, med *entity.MedicationEntity) error {
	return u.repo.UpdateMedication(ctx, med)
}

func NewUpdateMedicationUseCase(repo repository.PrescriptionRepo) UpdateMedicationUseCase {
	return &updateMedicationUseCaseImpl{repo: repo}
}
