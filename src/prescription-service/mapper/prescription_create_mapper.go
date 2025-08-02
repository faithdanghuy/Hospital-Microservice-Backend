package mapper

import (
	"time"

	"github.com/Hospital-Microservice/hospital-core/pointer"
	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/model/req"
	"github.com/google/uuid"
)

func TransformPrescriptionCreateReqToEntity(req req.PrescriptionCreateReq) *entity.PrescriptionEntity {
	return &entity.PrescriptionEntity{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		PatientID:     pointer.String(req.PatientID),
		DoctorID:      pointer.String(req.DoctorID),
		AppointmentID: pointer.String(req.AppointmentID),
		DrugName:      pointer.String(req.DrugName),
		Dosage:        pointer.String(req.Dosage),
		Instruction:   pointer.String(req.Instruction),
		Status:        pointer.String(req.Status),
		IssuedAt:      time.Now(),
	}
}
