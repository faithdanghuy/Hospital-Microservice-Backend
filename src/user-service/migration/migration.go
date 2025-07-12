package migration

import (
	"github.com/Hospital-Microservice/user-service/entity"
	"gorm.io/gorm"
)

func Must(db *gorm.DB) {
	db.Exec("CREATE SCHEMA IF NOT EXISTS user_service")
	err := db.Debug().AutoMigrate(
		&entity.UserEntity{},
	)
	if err != nil {
		panic(err)
	}
}
