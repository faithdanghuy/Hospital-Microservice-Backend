package repository

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
)

func (r *prescriptionRepoImpl) UpdateMedication(ctx context.Context, med *entity.MedicationEntity) error {
	return r.DB.Executor.WithContext(ctx).Save(med).Error
}
