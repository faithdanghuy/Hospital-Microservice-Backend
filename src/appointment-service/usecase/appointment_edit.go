package usecase

import (
	"context"
	"time"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/repository"
	"github.com/Hospital-Microservice/hospital-core/log"
	"go.uber.org/zap"
)

type AppointmentEditUseCase interface {
	Execute(ctx context.Context, id string, appointment *entity.AppointmentEntity) (*entity.AppointmentEntity, error)
}

type appointmentEditUseCaseImpl struct {
	appointmentRepo repository.AppointmentRepo
}

func (u *appointmentEditUseCaseImpl) Execute(ctx context.Context, id string, appointment *entity.AppointmentEntity) (*entity.AppointmentEntity, error) {
	appointment.ID = &id
	if appointment.Status != nil && *appointment.Status == "confirmed" && appointment.ConfirmedAt == nil {
		now := time.Now()
		appointment.ConfirmedAt = &now
	}
	if err := u.appointmentRepo.UpdateAppointment(ctx, appointment); err != nil {
		log.Error("failed to update appointment", zap.Error(err))
		return nil, err
	}

	updatedAppointment, err := u.appointmentRepo.FindAppointmentByID(ctx, id)
	if err != nil {
		log.Error("failed to fetch updated appointment", zap.Error(err))
		return nil, err
	}

	return updatedAppointment, nil
}

func NewAppointmentEditUseCase(repo repository.AppointmentRepo) AppointmentEditUseCase {
	return &appointmentEditUseCaseImpl{
		appointmentRepo: repo,
	}
}
