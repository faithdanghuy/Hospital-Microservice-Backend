package res

import "time"

type PrescriptionDetailRes struct {
	ID            string    `json:"id"`
	PatientID     string    `json:"patient_id"`
	DoctorID      string    `json:"doctor_id"`
	AppointmentID string    `json:"appointment_id"`
	DrugName      string    `json:"drug_name"`
	Dosage        string    `json:"dosage"`
	Instruction   string    `json:"instruction,omitempty"`
	Status        string    `json:"status,omitempty"`
	IssuedAt      time.Time `json:"created_at"`
}
