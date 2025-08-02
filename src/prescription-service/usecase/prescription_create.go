package usecase

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/repository"
	"go.uber.org/zap"
)

type PrescriptionCreateUseCase interface {
	Execute(ctx context.Context, prescription *entity.PrescriptionEntity) error
}

type prescriptionCreateUseCaseImpl struct {
	prescriptionRepo repository.PrescriptionRepo
}

func (r prescriptionCreateUseCaseImpl) Execute(ctx context.Context, prescription *entity.PrescriptionEntity) error {

	if err := r.prescriptionRepo.InsertPrescription(ctx, prescription); err != nil {
		log.Error("failed to insert prescription", zap.Error(err))
		return err
	}
	return nil
}

func NewPrescriptionCreateUseCase(
	PrescriptionRepo repository.PrescriptionRepo,
) PrescriptionCreateUseCase {
	return &prescriptionCreateUseCaseImpl{
		prescriptionRepo: PrescriptionRepo,
	}
}
