package repository

import (
	"context"

	"github.com/Hospital-Microservice/user-service/entity"
)

func (u *userRepoImpl) CreateEmptyDoctorProfile(
	ctx context.Context,
	userID *string,
) error {
	doctor := entity.DoctorProfileEntity{
		UserID:         userID,
		Specialization: nil,
	}

	return u.DB.Executor.WithContext(ctx).Create(&doctor).Error
}
