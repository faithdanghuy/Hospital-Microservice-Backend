package usecase

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/repository"
)

type CreateMedicationUseCase interface {
	Execute(ctx context.Context, med *entity.MedicationEntity) error
}

type createMedicationUseCaseImpl struct {
	repo repository.PrescriptionRepo
}

func NewCreateMedicationUseCase(repo repository.PrescriptionRepo) CreateMedicationUseCase {
	return &createMedicationUseCaseImpl{repo: repo}
}

func (u *createMedicationUseCaseImpl) Execute(ctx context.Context, med *entity.MedicationEntity) error {
	return u.repo.CreateMedication(ctx, med)
}
