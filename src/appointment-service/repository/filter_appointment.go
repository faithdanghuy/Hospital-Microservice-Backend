package repository

import (
	"context"
	"time"

	"github.com/Hospital-Microservice/appointment-service/entity"
)

func (u *appointmentRepoImpl) FilterAppointments(
	ctx context.Context,
	filter *entity.AppointmentEntity,
	fromDate *time.Time,
	toDate *time.Time,
) ([]*entity.AppointmentEntity, error) {
	var appointments []*entity.AppointmentEntity
	query := u.DB.Executor.WithContext(ctx).Model(&entity.AppointmentEntity{})

	if filter.PatientID != nil && *filter.PatientID != "" {
		query = query.Where("patient_id = ?", *filter.PatientID)
	}
	if filter.DoctorID != nil && *filter.DoctorID != "" {
		query = query.Where("doctor_id = ?", *filter.DoctorID)
	}
	if filter.Status != nil && *filter.Status != "" {
		query = query.Where("status = ?", *filter.Status)
	}
	if fromDate != nil {
		query = query.Where("scheduled_at >= ?", fromDate)
	}
	if toDate != nil {
		query = query.Where("scheduled_at <= ?", toDate)
	}

	if err := query.Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}
