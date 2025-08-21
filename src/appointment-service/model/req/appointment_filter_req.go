package req

type AppointmentFilterReq struct {
	PatientID string `json:"patient_id,omitempty" query:"patient_id"`
	DoctorID  string `json:"doctor_id,omitempty" query:"doctor_id"`
	Status    string `json:"status,omitempty" query:"status"`
	FromDate  string `json:"from_date,omitempty" query:"from_date"`
	ToDate    string `json:"to_date,omitempty" query:"to_date"`
}
