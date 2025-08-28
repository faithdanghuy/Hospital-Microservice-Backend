package repository

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
)

func (u *prescriptionRepoImpl) GetPrescription(ctx context.Context, id string) (*entity.PrescriptionEntity, error) {
	var p entity.PrescriptionEntity
	if err := u.DB.Executor.WithContext(ctx).
		Preload("Medications").
		First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
