package repository

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
)

func (u *userRepoImpl) CreateEmptyPatientProfile(
	ctx context.Context,
	userID *string,
) error {
	patient := entity.PatientProfileEntity{
		UserID:         userID,
		Gender:         nil,
		Address:        nil,
		MedicalHistory: nil,
	}

	return u.DB.Executor.WithContext(ctx).Create(&patient).Error
}
