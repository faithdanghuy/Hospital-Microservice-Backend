package repository

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
)

func (r *prescriptionRepoImpl) GetMedicationByID(ctx context.Context, id string) (*entity.MedicationEntity, error) {
	var med entity.MedicationEntity
	if err := r.DB.Executor.WithContext(ctx).First(&med, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &med, nil
}
