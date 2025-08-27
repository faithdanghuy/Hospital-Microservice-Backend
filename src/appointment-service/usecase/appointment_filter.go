package usecase

import (
	"context"
	"time"

	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/appointment-service/repository"
	"github.com/Hospital-Microservice/hospital-core/record"
)

type AppointmentFilterUseCase interface {
	Execute(ctx context.Context, pagination *record.Pagination, filter req.AppointmentFilterReq) (*record.Pagination, error)
}

type appointmentFilterUseCaseImpl struct {
	appointmentRepo repository.AppointmentRepo
}

func (u *appointmentFilterUseCaseImpl) Execute(
	ctx context.Context,
	pagination *record.Pagination,
	filter req.AppointmentFilterReq,
) (*record.Pagination, error) {
	var fromDate, toDate *time.Time

	if filter.FromDate != "" {
		if t, err := time.Parse("2006-01-02", filter.FromDate); err == nil {
			fromDate = &t
		}
	}
	if filter.ToDate != "" {
		if t, err := time.Parse("2006-01-02", filter.ToDate); err == nil {
			toDate = &t
		}
	}

	return u.appointmentRepo.FilterAppointments(ctx, pagination, filter, fromDate, toDate)
}

func NewAppointmentFilterUseCase(repo repository.AppointmentRepo) AppointmentFilterUseCase {
	return &appointmentFilterUseCaseImpl{appointmentRepo: repo}
}
