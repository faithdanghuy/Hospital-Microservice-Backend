package usecase

import (
	"context"
	"time"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/repository"
	"github.com/Hospital-Microservice/hospital-core/log"
	"go.uber.org/zap"
)

type AppointmentChangeStatusUseCase interface {
	Execute(ctx context.Context, id string, status string) (*entity.AppointmentEntity, error)
}

type appointmentChangeStatusUseCaseImpl struct {
	appointmentRepo repository.AppointmentRepo
}

func (u appointmentChangeStatusUseCaseImpl) Execute(
	ctx context.Context,
	id string,
	status string,
) (*entity.AppointmentEntity, error) {
	var confirmedAt *time.Time
	if status == "confirmed" {
		now := time.Now()
		confirmedAt = &now
	}

	appointment, err := u.appointmentRepo.ChangeAppointmentStatus(ctx, id, status)
	if err != nil {
		log.Error("failed to change appointment status", zap.Error(err))
		return nil, err
	}

	if confirmedAt != nil {
		appointment.ConfirmedAt = confirmedAt
		if err := u.appointmentRepo.UpdateAppointment(ctx, appointment); err != nil {
			log.Error("failed to update confirmedAt", zap.Error(err))
			return nil, err
		}
	}

	return appointment, nil
}

func NewAppointmentChangeStatusUseCase(
	appointmentRepo repository.AppointmentRepo,
) AppointmentChangeStatusUseCase {
	return &appointmentChangeStatusUseCaseImpl{
		appointmentRepo: appointmentRepo,
	}
}
