package usecase

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/repository"
)

type DetailMedicationUseCase interface {
	Execute(ctx context.Context, id string) (*entity.MedicationEntity, error)
}

type detailMedicationUseCaseImpl struct {
	repo repository.PrescriptionRepo
}

func (u *detailMedicationUseCaseImpl) Execute(ctx context.Context, id string) (*entity.MedicationEntity, error) {
	return u.repo.GetMedicationByID(ctx, id)
}

func NewDetailMedicationUseCase(repo repository.PrescriptionRepo) DetailMedicationUseCase {
	return &detailMedicationUseCaseImpl{repo: repo}
}
