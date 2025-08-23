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
	Birthday time.Time
	Role     *string `gorm:"type:enum('patient','doctor','admin');default:'patient'"`

	PatientProfile *PatientProfileEntity `gorm:"foreignKey:UserID;references:ID"`
	DoctorProfile  *DoctorProfileEntity  `gorm:"foreignKey:UserID;references:ID"`
}

func (u UserEntity) TableName() string {
	return "user_service.users"
}
