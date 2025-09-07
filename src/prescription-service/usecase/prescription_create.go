package usecase

import (
	"context"
	"encoding/json"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/provider"
	"github.com/Hospital-Microservice/prescription-service/repository"
	"go.uber.org/zap"

	"github.com/Hospital-Microservice/hospital-core/model/req"
	rabbit "github.com/Hospital-Microservice/hospital-core/provider"
)

type PrescriptionCreateUseCase interface {
	Execute(ctx context.Context, prescription *entity.PrescriptionEntity) error
}

type prescriptionCreateUseCaseImpl struct {
	prescriptionRepo repository.PrescriptionRepo
	publisher        *rabbit.RabbitPublisher
	UserService      provider.UserService
}

func NewPrescriptionCreateUseCase(repo repository.PrescriptionRepo, publisher *rabbit.RabbitPublisher, userClient provider.UserService) PrescriptionCreateUseCase {
	return &prescriptionCreateUseCaseImpl{
		prescriptionRepo: repo,
		publisher:        publisher,
		UserService:      userClient,
	}
}

func (r *prescriptionCreateUseCaseImpl) Execute(ctx context.Context, prescription *entity.PrescriptionEntity) error {
	if err := r.prescriptionRepo.InsertPrescription(ctx, prescription); err != nil {
		log.Error("Failed To Insert Prescription", zap.Error(err))
		return err
	}

	var ids []string
	if prescription.PatientID != nil {
		ids = append(ids, *prescription.PatientID)
	}
	if prescription.DoctorID != nil {
		ids = append(ids, *prescription.DoctorID)
	}

	users, err := r.UserService.GetUsersByIDs(ctx, ids, "")
	if err != nil {
		log.Error("failed to fetch users for notification", zap.Error(err))
		return nil
	}

	var emails []string
	for _, u := range users {
		if u.Email != "" {
			emails = append(emails, u.Email)
		}
	}

	if len(emails) > 0 || len(ids) > 0 {
		idsBytes, _ := json.Marshal(ids)

		notify := req.NotificationReq{
			ToEmails: emails,
			Subject:  "Prescription Scheduled Successfully",
			Body:     "Your prescription has been scheduled successfully.",
			Meta: map[string]any{
				"user_ids": string(idsBytes),
			},
		}

		log.Error("failed to publish notify message", zap.Error(err))
		body, err := json.Marshal(notify)
		if err != nil {
			log.Error("failed to marshal notify message", zap.Error(err))
		} else if err := r.publisher.Publish(ctx, body); err != nil {
			log.Error("failed to publish notify message", zap.Error(err))
		}
	}

	return nil
}
