package repository

import (
	"context"

	"github.com/Hospital-Microservice/appointment-service/entity"
)

func (u *appointmentRepoImpl) ChangeAppointmentStatus(
	ctx context.Context,
	id string,
	status string,
) (*entity.AppointmentEntity, error) {
	var appointmentEntity entity.AppointmentEntity

	if err := u.DB.Executor.WithContext(ctx).
		Where("id = ?", id).
		First(&appointmentEntity).Error; err != nil {
		return nil, err
	}

	appointmentEntity.Status = &status

	if err := u.DB.Executor.WithContext(ctx).Save(&appointmentEntity).Error; err != nil {
		return nil, err
	}

	return &appointmentEntity, nil
}
