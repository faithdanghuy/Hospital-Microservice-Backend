package migration

import (
	"github.com/Hospital-Microservice/prescription-service/entity"
	"gorm.io/gorm"
)

func Must(db *gorm.DB) {
	db.Exec("CREATE SCHEMA IF NOT EXISTS prescription_service")
	err := db.Debug().AutoMigrate(
		&entity.PrescriptionEntity{},
	)
	if err != nil {
		panic(err)
	}
}
