package migration

import (
	"github.com/Hospital-Microservice/appointment-service/entity"
	"gorm.io/gorm"
)

func Must(db *gorm.DB) {
	db.Exec("CREATE SCHEMA IF NOT EXISTS appointment_service")
	err := db.Debug().AutoMigrate(
		&entity.AppointmentEntity{},
	)
	if err != nil {
		panic(err)
	}
}
