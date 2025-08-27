package repository

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
	"gorm.io/gorm"
)

func (u *prescriptionRepoImpl) InsertPrescription(ctx context.Context, prescription *entity.PrescriptionEntity) error {
	return u.DB.Executor.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&prescription).Error; err != nil {
			return err
		}

		for _, med := range prescription.Medications {
			med.PrescriptionID = prescription.ID
			if err := tx.Create(med).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
