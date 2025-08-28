package usecase

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/repository"
)

type UpdatePrescriptionUseCase interface {
	Execute(ctx context.Context, p *entity.PrescriptionEntity) error
}

type updatePrescriptionUseCaseImpl struct {
	repo repository.PrescriptionRepo
}

func NewUpdatePrescriptionUseCase(repo repository.PrescriptionRepo) UpdatePrescriptionUseCase {
	return &updatePrescriptionUseCaseImpl{repo: repo}
}
func (u *updatePrescriptionUseCaseImpl) Execute(ctx context.Context, p *entity.PrescriptionEntity) error {
	return u.repo.UpdatePrescription(ctx, p)
}
