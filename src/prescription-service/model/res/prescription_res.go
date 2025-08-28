package res

import (
	"time"

	"github.com/Hospital-Microservice/prescription-service/provider"
)

type PrescriptionRes struct {
	ID        string            `json:"id"`
	PatientID string            `json:"patient_id"`
	Patient   *provider.UserRes `json:"patient,omitempty"`
	DoctorID  string            `json:"doctor_id"`
	Doctor    *provider.UserRes `json:"doctor,omitempty"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`

	Medications []PrescriptionMedRes `json:"medications"`
}

type PrescriptionMedRes struct {
	ID           string    `json:"id"`
	MedicationID string    `json:"medication_id"`
	Quantity     int       `json:"quantity"`
	Dosage       string    `json:"dosage"`
	Instruction  string    `json:"instruction,omitempty"`
	IssuedAt     time.Time `json:"issued_at"`
}
