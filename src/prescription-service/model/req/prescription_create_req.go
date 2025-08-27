package req

import "time"

type PrescriptionMedReq struct {
	MedicationID string    `json:"medication_id" validate:"required"`
	Quantity     *int      `json:"quantity" validate:"required"`
	Unit         string    `json:"unit" validate:"required"`
	Instruction  *string   `json:"instruction,omitempty"`
	IssuedAt     time.Time `json:"issued_at" validate:"required"`
}

type PrescriptionCreateReq struct {
	PatientID   string               `json:"patient_id" validate:"required"`
	DoctorID    string               `json:"doctor_id" validate:"required"`
	Status      string               `json:"status" validate:"required"`
	Medications []PrescriptionMedReq `json:"medications"`
}
