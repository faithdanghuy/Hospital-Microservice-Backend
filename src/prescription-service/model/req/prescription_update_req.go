package req

type PrescriptionUpdateReq struct {
	ID          string               `json:"id" validate:"required"`
	PatientID   string               `json:"patient_id"`
	DoctorID    string               `json:"doctor_id"`
	Status      string               `json:"status"`
	Medications []PrescriptionMedReq `json:"medications"`
}
