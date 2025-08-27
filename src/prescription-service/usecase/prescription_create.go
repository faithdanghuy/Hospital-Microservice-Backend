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

func NewPrescriptionCreateUseCase(repo repository.PrescriptionRepo) PrescriptionCreateUseCase {
	return &prescriptionCreateUseCaseImpl{
		prescriptionRepo: repo,
	}
}

func (r *prescriptionCreateUseCaseImpl) Execute(ctx context.Context, prescription *entity.PrescriptionEntity) error {
	if err := r.prescriptionRepo.InsertPrescription(ctx, prescription); err != nil {
		log.Error("Failed To Insert Prescription", zap.Error(err))
		return err
	}
	return nil
}
