package mapper

import (
	"time"

	"github.com/Hospital-Microservice/hospital-core/pointer"
	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/model/req"
)

func TransformPrescriptionUpdateReqToEntity(req req.PrescriptionUpdateReq) *entity.PrescriptionEntity {
	p := &entity.PrescriptionEntity{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(req.ID),
		},
		PatientID: pointer.String(req.PatientID),
		DoctorID:  pointer.String(req.DoctorID),
		Status:    pointer.String(req.Status),
	}
	p.ID = pointer.String(req.ID)

	meds := make([]*entity.PrescMedEntity, 0, len(req.Medications))
	for _, m := range req.Medications {
		meds = append(meds, &entity.PrescMedEntity{
			MedicationID: pointer.String(m.MedicationID),
			Quantity:     m.Quantity,
			Dosage:       m.Dosage,
			Instruction:  m.Instruction,
			IssuedAt:     time.Now(),
		})
	}
	p.Medications = meds
	return p
}
