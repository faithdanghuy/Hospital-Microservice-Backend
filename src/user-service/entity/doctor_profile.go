package entity

import (
	"github.com/Hospital-Microservice/hospital-core/record"
)

type DoctorProfileEntity struct {
	record.BaseEntity
	UserID         *string // FK → User.ID
	Specialization *string

	User UserEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`
}

func (u DoctorProfileEntity) TableName() string {
	return "user_service.doctor_profile"
}
