package repository

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/model/req"
)

func (r *prescriptionRepoImpl) ListMedications(ctx context.Context, pagination *record.Pagination, filter *req.MedicationFilterReq) (*record.Pagination, error) {
	var meds []*entity.MedicationEntity
	query := r.DB.Executor.WithContext(ctx).Model(&entity.MedicationEntity{})

	// Apply filters
	if filter.DrugName != "" {
		query = query.Where("drug_name ILIKE ?", "%"+filter.DrugName+"%")
	}
	if filter.Unit != "" {
		query = query.Where("unit = ?", filter.Unit)
	}

	// Count total rows
	var totalRows int64
	if err := query.Count(&totalRows).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	if err := query.
		Order(pagination.GetSort()).
		Limit(pagination.GetLimit()).
		Offset(pagination.GetOffset()).
		Find(&meds).Error; err != nil {
		return nil, err
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = int((totalRows + int64(pagination.GetLimit()) - 1) / int64(pagination.GetLimit()))
	pagination.Rows = meds

	return pagination, nil
}
