package usecase

import (
	"context"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/repository"
)

type AppointmentDetailUseCase interface {
	Execute(ctx context.Context, id string) (*entity.AppointmentEntity, error)
}

type appointmentDetailUseCaseImpl struct {
	appointmentRepo repository.AppointmentRepo
}

func (l *appointmentDetailUseCaseImpl) Execute(
	ctx context.Context,
	id string,
) (*entity.AppointmentEntity, error) {
	appointment, err := l.appointmentRepo.FindAppointmentByID(
		ctx,
		id,
	)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func NewAppointmentDetailUseCase(
	AppointmentRepo repository.AppointmentRepo,
) AppointmentDetailUseCase {
	return &appointmentDetailUseCaseImpl{
		appointmentRepo: AppointmentRepo,
	}
}
