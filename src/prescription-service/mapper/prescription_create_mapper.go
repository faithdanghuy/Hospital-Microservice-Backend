package mapper

import (
	"github.com/Hospital-Microservice/hospital-core/pointer"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/model/req"
)

func TransformPrescriptionCreateReqToEntity(req req.PrescriptionCreateReq) *entity.PrescriptionEntity {
	p := &entity.PrescriptionEntity{
		PatientID: pointer.String(req.PatientID),
		DoctorID:  pointer.String(req.DoctorID),
		Status:    pointer.String(req.Status),
	}

	meds := make([]*entity.PrescMedEntity, 0, len(req.Medications))
	for _, m := range req.Medications {
		meds = append(meds, &entity.PrescMedEntity{
			MedicationID: pointer.String(m.MedicationID),
			Quantity:     m.Quantity,
			Dosage:       m.Dosage,
			Instruction:  m.Instruction,
			IssuedAt:     m.IssuedAt,
		})
	}
	p.Medications = meds
	return p
}
