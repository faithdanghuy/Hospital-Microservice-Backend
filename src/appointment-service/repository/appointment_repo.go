package repository

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/db"

	"github.com/Hospital-Microservice/appointment-service/entity"
)

type AppointmentRepo interface {
	FindAppointmentByID(ctx context.Context, ID string) (*entity.AppointmentEntity, error)
	InsertAppointment(ctx context.Context, appointment *entity.AppointmentEntity) error
}

type appointmentRepoImpl struct {
	DB *db.Database
}

func NewAppointmentRepo(db *db.Database) AppointmentRepo {
	return &appointmentRepoImpl{
		DB: db,
	}
}
