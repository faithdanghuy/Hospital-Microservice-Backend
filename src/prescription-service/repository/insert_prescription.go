package repository

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
)

func (u *prescriptionRepoImpl) InsertPrescription(
	ctx context.Context,
	prescription *entity.PrescriptionEntity,
) error {
	return u.DB.Executor.WithContext(ctx).Create(&prescription).Error
}
