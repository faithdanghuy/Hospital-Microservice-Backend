package mapper

import (
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/model/res"
)

func TransformPrescriptionEntityToRes(prescriptionEntity *entity.PrescriptionEntity) *res.PrescriptionDetailRes {
	return &res.PrescriptionDetailRes{
		ID:            *prescriptionEntity.ID,
		PatientID:     *prescriptionEntity.PatientID,
		DoctorID:      *prescriptionEntity.DoctorID,
		AppointmentID: *prescriptionEntity.AppointmentID,
		DrugName:      *prescriptionEntity.DrugName,
		Dosage:        *prescriptionEntity.Dosage,
		Instruction:   *prescriptionEntity.Instruction,
		Status:        *prescriptionEntity.Status,
		IssuedAt:      prescriptionEntity.IssuedAt,
	}
}
