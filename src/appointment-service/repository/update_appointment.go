package repository

import (
	"context"

	"github.com/Hospital-Microservice/appointment-service/entity"
)

func (u *appointmentRepoImpl) UpdateAppointment(
	ctx context.Context,
	appointment *entity.AppointmentEntity,
) error {
	var existing entity.AppointmentEntity
	if err := u.DB.Executor.WithContext(ctx).
		Where("id = ?", appointment.ID).
		First(&existing).Error; err != nil {
		return err
	}

	if err := u.DB.Executor.WithContext(ctx).
		Model(&existing).
		Updates(appointment).Error; err != nil {
		return err
	}

	return nil
}
