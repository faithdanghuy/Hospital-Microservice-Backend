package usecase

import (
	"context"
	"time"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/repository"
)

type AppointmentFilterUseCase interface {
	Execute(ctx context.Context, filter *entity.AppointmentEntity, fromDateStr, toDateStr string) ([]*entity.AppointmentEntity, error)
}

type appointmentFilterUseCaseImpl struct {
	appointmentRepo repository.AppointmentRepo
}

func (u *appointmentFilterUseCaseImpl) Execute(
	ctx context.Context,
	filter *entity.AppointmentEntity,
	fromDateStr, toDateStr string,
) ([]*entity.AppointmentEntity, error) {
	var fromDate, toDate *time.Time

	if fromDateStr != "" {
		if t, err := time.Parse("2006-01-02", fromDateStr); err == nil {
			fromDate = &t
		}
	}
	if toDateStr != "" {
		if t, err := time.Parse("2006-01-02", toDateStr); err == nil {
			toDate = &t
		}
	}

	return u.appointmentRepo.FilterAppointments(ctx, filter, fromDate, toDate)
}

func NewAppointmentFilterUseCase(repo repository.AppointmentRepo) AppointmentFilterUseCase {
	return &appointmentFilterUseCaseImpl{appointmentRepo: repo}
}
