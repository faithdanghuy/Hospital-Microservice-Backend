package repository

import (
	"context"
	"time"

	"github.com/Hospital-Microservice/hospital-core/db"

	"github.com/Hospital-Microservice/appointment-service/entity"
)

type AppointmentRepo interface {
	FindAppointmentByID(ctx context.Context, ID string) (*entity.AppointmentEntity, error)
	InsertAppointment(ctx context.Context, appointment *entity.AppointmentEntity) error
	ChangeAppointmentStatus(ctx context.Context, id string, status string) (*entity.AppointmentEntity, error)
	UpdateAppointment(ctx context.Context, appointment *entity.AppointmentEntity) error
	FilterAppointments(ctx context.Context, filter *entity.AppointmentEntity, fromDate *time.Time, toDate *time.Time) ([]*entity.AppointmentEntity, error)
}

type appointmentRepoImpl struct {
	DB *db.Database
}

func NewAppointmentRepo(db *db.Database) AppointmentRepo {
	return &appointmentRepoImpl{
		DB: db,
	}
}
