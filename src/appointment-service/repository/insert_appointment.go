package repository

import (
	"context"

	"github.com/Hospital-Microservice/appointment-service/entity"
)

func (u *appointmentRepoImpl) InsertAppointment(
	ctx context.Context,
	appointment *entity.AppointmentEntity,
) error {
	return u.DB.Executor.WithContext(ctx).Create(&appointment).Error
}
