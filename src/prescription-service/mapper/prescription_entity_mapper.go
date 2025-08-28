package mapper

import (
	"github.com/Hospital-Microservice/hospital-core/pointer"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/model/res"
	"github.com/Hospital-Microservice/prescription-service/provider"
)

func TransformPrescriptionEntityToRes(
	p *entity.PrescriptionEntity,
	users map[string]*provider.UserRes,
) *res.PrescriptionRes {
	resp := &res.PrescriptionRes{
		ID:        *p.ID,
		PatientID: *p.PatientID,
		DoctorID:  *p.DoctorID,
		Status:    *p.Status,
		CreatedAt: pointer.Time(p.CreatedAt),
		UpdatedAt: pointer.Time(p.UpdatedAt),
	}

	if u, ok := users[*p.PatientID]; ok {
		resp.Patient = u
	}
	if u, ok := users[*p.DoctorID]; ok {
		resp.Doctor = u
	}

	meds := make([]res.PrescriptionMedRes, 0, len(p.Medications))
	for _, m := range p.Medications {
		meds = append(meds, res.PrescriptionMedRes{
			ID:           *m.ID,
			MedicationID: *m.MedicationID,
			Quantity:     *m.Quantity,
			Dosage:       derefString(m.Dosage),
			Instruction:  derefString(m.Instruction),
			IssuedAt:     m.IssuedAt,
		})
	}
	resp.Medications = meds
	return resp
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
