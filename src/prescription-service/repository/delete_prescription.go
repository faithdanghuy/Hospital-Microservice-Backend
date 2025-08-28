package repository

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
)

func (u *prescriptionRepoImpl) DeletePrescription(ctx context.Context, id string) error {
	return u.DB.Executor.WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.PrescriptionEntity{}).Error
}
