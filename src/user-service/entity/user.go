package entity

import (
	"time"

	"github.com/Hospital-Microservice/hospital-core/record"
)

type UserEntity struct {
	record.BaseEntity
	FullName *string
	Email    *string
	Phone    *string `gorm:"unique"`
	Password *string
	Avatar   *string
	Birthday time.Time
}

func (u UserEntity) TableName() string {
	return "user_service.users"
}
