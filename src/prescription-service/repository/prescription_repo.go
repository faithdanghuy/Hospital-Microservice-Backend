package repository

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/db"

	"github.com/Hospital-Microservice/prescription-service/entity"
)

type PrescriptionRepo interface {
	FindPrescriptionByID(ctx context.Context, ID string) (*entity.PrescriptionEntity, error)
	InsertPrescription(ctx context.Context, prescription *entity.PrescriptionEntity) error
}

type prescriptionRepoImpl struct {
	DB *db.Database
}

func NewPrescriptionRepo(db *db.Database) PrescriptionRepo {
	return &prescriptionRepoImpl{
		DB: db,
	}
}
