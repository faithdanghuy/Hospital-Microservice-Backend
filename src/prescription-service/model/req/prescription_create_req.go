package req

type PrescriptionMedReq struct {
	MedicationID string  `json:"medication_id" validate:"required"`
	Quantity     *int    `json:"quantity" validate:"required"`
	Dosage       *string `json:"dosage,omitempty"`
	Instruction  *string `json:"instruction,omitempty"`
}

type PrescriptionCreateReq struct {
	PatientID   string               `json:"patient_id" validate:"required"`
	DoctorID    string               `json:"doctor_id" validate:"required"`
	Status      string               `json:"status" validate:"required"`
	Medications []PrescriptionMedReq `json:"medications"`
}
