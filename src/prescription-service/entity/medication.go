package entity

import (
	"github.com/Hospital-Microservice/hospital-core/record"
)

type MedicationEntity struct {
	record.BaseEntity
	DrugName    *string `gorm:"not null"`
	Dosage      *string `gorm:"not null"`
	Description *string `gorm:"type:text"`

	PrescMeds []*PrescMedEntity `gorm:"foreignKey:MedicationID;references:ID"`
}

func (MedicationEntity) TableName() string {
	return "prescription_service.medications"
}
