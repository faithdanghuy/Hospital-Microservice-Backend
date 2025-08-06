package usecase

import (
	"context"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/repository"
	"github.com/Hospital-Microservice/hospital-core/log"
	"go.uber.org/zap"
)

type AppointmentCreateUseCase interface {
	Execute(ctx context.Context, appointment *entity.AppointmentEntity) error
}

type appointmentCreateUseCaseImpl struct {
	appointmentRepo repository.AppointmentRepo
}

func (r appointmentCreateUseCaseImpl) Execute(ctx context.Context, appointment *entity.AppointmentEntity) error {

	if err := r.appointmentRepo.InsertAppointment(ctx, appointment); err != nil {
		log.Error("failed to insert appointment", zap.Error(err))
		return err
	}
	return nil
}

func NewAppointmentCreateUseCase(
	AppointmentRepo repository.AppointmentRepo,
) AppointmentCreateUseCase {
	return &appointmentCreateUseCaseImpl{
		appointmentRepo: AppointmentRepo,
	}
}
