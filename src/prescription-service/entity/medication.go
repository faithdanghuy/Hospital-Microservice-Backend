package entity

import (
	"github.com/Hospital-Microservice/hospital-core/record"
)

type MedicationEntity struct {
	record.BaseEntity
	DrugName    *string `gorm:"not null"`
	Stock       *int    `gorm:"not null"`
	Unit        *string `gorm:"not null;check:unit IN ('tablet','capsule','syrup','injection','ointment','drop','inhaler','patch','suppository','other');default:'tablet'"`
	Description *string `gorm:"type:text"`

	PrescMeds []*PrescMedEntity `gorm:"foreignKey:MedicationID;references:ID"`
}

func (MedicationEntity) TableName() string {
	return "prescription_service.medications"
}
