package entity

import (
	"github.com/Hospital-Microservice/hospital-core/record"
)

type PrescriptionEntity struct {
	record.BaseEntity
	PatientID *string `gorm:"not null"`
	DoctorID  *string `gorm:"not null"`
	Status    *string `gorm:"default:'not_collected'; check:status IN ('not_collected','collected');index"`

	Medications []*PrescMedEntity `gorm:"foreignKey:PrescriptionID;references:ID"`
}

func (PrescriptionEntity) TableName() string {
	return "prescription_service.prescriptions"
}
