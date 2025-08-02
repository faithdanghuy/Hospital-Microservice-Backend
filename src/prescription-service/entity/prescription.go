package entity

import (
	"time"

	"github.com/Hospital-Microservice/hospital-core/record"
)

type PrescriptionEntity struct {
	record.BaseEntity
	PatientID     *string   `gorm:"not null"`
	DoctorID      *string   `gorm:"not null"`
	AppointmentID *string   `gorm:"not null"`
	DrugName      *string   `gorm:"not null"`
	Dosage        *string   `gorm:"not null"`
	Instruction   *string   `gorm:"type:text"`
	Status        *string   `gorm:"default:'not_collected'"`
	IssuedAt      time.Time `gorm:"autoCreateTime"`
}

func (PrescriptionEntity) TableName() string {
	return "prescription_service.prescriptions"
}
