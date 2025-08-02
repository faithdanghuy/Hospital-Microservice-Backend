package req

type PrescriptionCreateReq struct {
	PatientID     string `json:"patient_id" validate:"required"`
	DoctorID      string `json:"doctor_id" validate:"required"`
	AppointmentID string `json:"appointment_id" validate:"required"`
	DrugName      string `json:"drug_name" validate:"required"`
	Dosage        string `json:"dosage" validate:"required"`
	Instruction   string `json:"instruction" validate:"omitempty"`
	Status        string `json:"status" validate:"omitempty,oneof=not_collected collected"`
}
