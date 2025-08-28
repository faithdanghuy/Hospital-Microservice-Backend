package repository

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
)

func (r *prescriptionRepoImpl) DeleteMedication(ctx context.Context, id string) error {
	return r.DB.Executor.WithContext(ctx).Delete(&entity.MedicationEntity{}, "id = ?", id).Error
}
