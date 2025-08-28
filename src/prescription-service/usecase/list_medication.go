package usecase

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/prescription-service/model/req"
	"github.com/Hospital-Microservice/prescription-service/repository"
)

type ListMedicationUseCase interface {
	Execute(ctx context.Context, pagination *record.Pagination, filter *req.MedicationFilterReq) (*record.Pagination, error)
}

type listMedicationUseCaseImpl struct {
	repo repository.PrescriptionRepo
}

func NewListMedicationUseCase(repo repository.PrescriptionRepo) ListMedicationUseCase {
	return &listMedicationUseCaseImpl{repo: repo}
}

func (u *listMedicationUseCaseImpl) Execute(ctx context.Context, pagination *record.Pagination, filter *req.MedicationFilterReq) (*record.Pagination, error) {
	return u.repo.ListMedications(ctx, pagination, filter)
}
