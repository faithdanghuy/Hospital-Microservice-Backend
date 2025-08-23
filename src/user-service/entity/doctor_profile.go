package entity

import (
	"github.com/Hospital-Microservice/hospital-core/record"
)

type DoctorProfileEntity struct {
	record.BaseEntity
	UserID         uint `gorm:"uniqueIndex"` // FK → User.ID
	Specialization string
	Phone          string
	Email          string

	User UserEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`
}

func (u DoctorProfileEntity) TableName() string {
	return "user_service.doctor_profile"
}
