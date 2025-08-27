package entity

import (
	"time"

	"github.com/Hospital-Microservice/hospital-core/record"
)

type PrescMedEntity struct {
	record.BaseEntity
	PrescriptionID *string   `gorm:"not null"`
	MedicationID   *string   `gorm:"not null"`
	Quantity       *int      `gorm:"not null"`
	Unit           *string   `gorm:"not null"`
	Instruction    *string   `gorm:"type:text"`
	IssuedAt       time.Time `gorm:"not null"`

	Prescription *PrescriptionEntity `gorm:"foreignKey:PrescriptionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Medication   *MedicationEntity   `gorm:"foreignKey:MedicationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (PrescMedEntity) TableName() string {
	return "prescription_service.prescription_medications"
}
