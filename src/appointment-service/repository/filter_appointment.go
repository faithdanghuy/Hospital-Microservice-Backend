package repository

import (
	"context"
	"time"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/hospital-core/record"
)

func (u *appointmentRepoImpl) FilterAppointments(
	ctx context.Context,
	pagination *record.Pagination,
	filter req.AppointmentFilterReq,
	fromDate *time.Time,
	toDate *time.Time,
) (*record.Pagination, error) {
	var appointments []entity.AppointmentEntity
	var totalRows int64

	query := u.DB.Executor.WithContext(ctx).Model(&entity.AppointmentEntity{})

	if filter.PatientID != "" {
		query = query.Where("patient_id = ?", filter.PatientID)
	}
	if filter.DoctorID != "" {
		query = query.Where("doctor_id = ?", filter.DoctorID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if fromDate != nil {
		query = query.Where("scheduled_at >= ?", fromDate)
	}
	if toDate != nil {
		query = query.Where("scheduled_at <= ?", toDate)
	}

	if err := query.Count(&totalRows).Error; err != nil {
		return nil, err
	}

	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	sort := pagination.GetSort()

	if err := query.
		Order(sort).
		Limit(limit).
		Offset(offset).
		Find(&appointments).Error; err != nil {
		return nil, err
	}

	pagination.Rows = appointments
	pagination.TotalRows = totalRows
	pagination.TotalPages = int((totalRows + int64(limit) - 1) / int64(limit))

	return pagination, nil
}
