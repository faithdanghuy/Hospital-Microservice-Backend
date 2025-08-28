package usecase

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/prescription-service/model/req"
	"github.com/Hospital-Microservice/prescription-service/repository"
)

type FilterPrescriptionUseCase interface {
	Execute(ctx context.Context, pagination *record.Pagination, filter *req.PrescriptionFilterReq) (*record.Pagination, error)
}

type filterPrescriptionUseCaseImpl struct {
	repo repository.PrescriptionRepo
}

func NewFilterPrescriptionUseCase(repo repository.PrescriptionRepo) FilterPrescriptionUseCase {
	return &filterPrescriptionUseCaseImpl{repo: repo}
}
func (u *filterPrescriptionUseCaseImpl) Execute(ctx context.Context, pagination *record.Pagination, filter *req.PrescriptionFilterReq) (*record.Pagination, error) {
	return u.repo.FilterPrescriptions(ctx, pagination, filter)
}
