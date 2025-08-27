package entity

import (
	"time"

	"github.com/Hospital-Microservice/hospital-core/record"
)

type AppointmentEntity struct {
	record.BaseEntity
	PatientID   *string    `gorm:"not null"`
	DoctorID    *string    `gorm:"not null"`
	ScheduledAt time.Time  `gorm:"not null"`
	Status      *string    `gorm:"type:text;check:status IN ('pending','confirmed','cancelled');default:'pending';index"`
	Note        *string    `gorm:"type:text"`
	ConfirmedAt *time.Time `gorm:""`
}

func (AppointmentEntity) TableName() string {
	return "appointment_service.appointments"
}
