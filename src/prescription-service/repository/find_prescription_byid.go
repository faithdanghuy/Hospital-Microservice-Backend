package repository

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
)

func (u *prescriptionRepoImpl) FindPrescriptionByID(
	ctx context.Context,
	ID string,
) (*entity.PrescriptionEntity, error) {
	var prescriptionEntity entity.PrescriptionEntity

	var findQuery = u.DB.Executor.WithContext(ctx).
		Where("id = ?", ID)

	if err := findQuery.Find(&prescriptionEntity).Error; err != nil {
		return nil, err
	}

	return &prescriptionEntity, nil
}
