package repository

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/model/req"
)

func (u *prescriptionRepoImpl) FilterPrescriptions(ctx context.Context, pagination *record.Pagination, filter *req.PrescriptionFilterReq) (*record.Pagination, error) {
	var prescriptions []*entity.PrescriptionEntity
	query := u.DB.Executor.WithContext(ctx).Model(&entity.PrescriptionEntity{}).Preload("Medications")

	if filter.PatientID != "" {
		query = query.Where("patient_id = ?", filter.PatientID)
	}
	if filter.DoctorID != "" {
		query = query.Where("doctor_id = ?", filter.DoctorID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	var totalRows int64
	if err := query.Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if err := query.
		Order(pagination.GetSort()).
		Limit(pagination.GetLimit()).
		Offset(pagination.GetOffset()).
		Find(&prescriptions).Error; err != nil {
		return nil, err
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = int((totalRows + int64(pagination.GetLimit()) - 1) / int64(pagination.GetLimit()))
	pagination.Rows = prescriptions
	return pagination, nil
}
