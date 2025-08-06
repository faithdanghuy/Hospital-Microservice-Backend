package repository

import (
	"context"

	"github.com/Hospital-Microservice/appointment-service/entity"
)

func (u *appointmentRepoImpl) FindAppointmentByID(
	ctx context.Context,
	ID string,
) (*entity.AppointmentEntity, error) {
	var appointmentEntity entity.AppointmentEntity

	var findQuery = u.DB.Executor.WithContext(ctx).
		Where("id = ?", ID)

	if err := findQuery.Find(&appointmentEntity).Error; err != nil {
		return nil, err
	}

	return &appointmentEntity, nil
}
