package repository

import (
	"context"

	"github.com/Hospital-Microservice/prescription-service/entity"
	"gorm.io/gorm"
)

func (u *prescriptionRepoImpl) UpdatePrescription(ctx context.Context, p *entity.PrescriptionEntity) error {
	return u.DB.Executor.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.PrescriptionEntity{}).
			Where("id = ?", p.ID).
			Updates(p).Error; err != nil {
			return err
		}
		if len(p.Medications) > 0 {
			if err := tx.Where("prescription_id = ?", p.ID).Delete(&entity.PrescMedEntity{}).Error; err != nil {
				return err
			}
			for _, med := range p.Medications {
				med.PrescriptionID = p.ID
				if err := tx.Create(med).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}
