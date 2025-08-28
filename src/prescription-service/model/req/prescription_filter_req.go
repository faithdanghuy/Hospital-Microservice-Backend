package req

type PrescriptionFilterReq struct {
	PatientID string `query:"patient_id"`
	DoctorID  string `query:"doctor_id"`
	Status    string `query:"status"`
}
