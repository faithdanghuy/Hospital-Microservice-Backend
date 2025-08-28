package repository

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/db"
	"github.com/Hospital-Microservice/hospital-core/record"

	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/model/req"
)

type PrescriptionRepo interface {
	FindPrescriptionByID(ctx context.Context, ID string) (*entity.PrescriptionEntity, error)
	InsertPrescription(ctx context.Context, prescription *entity.PrescriptionEntity) error
	CreateMedication(ctx context.Context, med *entity.MedicationEntity) error
	GetMedicationByID(ctx context.Context, id string) (*entity.MedicationEntity, error)
	UpdateMedication(ctx context.Context, med *entity.MedicationEntity) error
	DeleteMedication(ctx context.Context, id string) error
	ListMedications(ctx context.Context, pagination *record.Pagination, filter *req.MedicationFilterReq) (*record.Pagination, error)
}

type prescriptionRepoImpl struct {
	DB *db.Database
}

func NewPrescriptionRepo(db *db.Database) PrescriptionRepo {
	return &prescriptionRepoImpl{
		DB: db,
	}
}
