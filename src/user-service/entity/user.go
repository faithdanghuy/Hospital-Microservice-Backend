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
	Password *string `gorm:"not null"`
	Avatar   *string
	Gender   *string
	Address  *string
	Birthday time.Time
	Role     *string `gorm:"type:text;check:role IN ('patient','doctor','admin');default:'patient';index"`
}

func (u UserEntity) TableName() string {
	return "user_service.users"
}
