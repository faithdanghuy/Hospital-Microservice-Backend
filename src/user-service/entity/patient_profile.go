package entity

import (
	"github.com/Hospital-Microservice/hospital-core/record"
)

type PatientProfileEntity struct {
	record.BaseEntity
	UserID         *string // FK → User.ID
	Gender         *string
	Address        *string
	MedicalHistory *string

	User UserEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`
}

func (u PatientProfileEntity) TableName() string {
	return "user_service.patient_profile"
}
