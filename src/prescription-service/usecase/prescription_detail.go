package usecase

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/repository"
)

type PrescriptionDetailUseCase interface {
	Execute(ctx context.Context, id string) (*entity.PrescriptionEntity, error)
}

type prescriptionDetailUseCaseImpl struct {
	prescriptionRepo repository.PrescriptionRepo
}

func (l *prescriptionDetailUseCaseImpl) Execute(
	ctx context.Context,
	id string,
) (*entity.PrescriptionEntity, error) {
	prescription, err := l.prescriptionRepo.FindPrescriptionByID(
		ctx,
		id,
	)
	if err != nil {
		return nil, err
	}

	return prescription, nil
}

func NewPrescriptionDetailUseCase(
	PrescriptionRepo repository.PrescriptionRepo,
) PrescriptionDetailUseCase {
	return &prescriptionDetailUseCaseImpl{
		prescriptionRepo: PrescriptionRepo,
	}
}
