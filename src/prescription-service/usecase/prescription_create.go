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

	if len(emails) > 0 {
		notify := req.NotificationReq{
			ToEmails: emails,
			Subject:  "Medication Prescribed",
			Body:     "Your medication has been prescribed successfully.",
		}
		body, _ := json.Marshal(notify)
		if err := r.publisher.Publish(ctx, body); err != nil {
			log.Error("failed to publish notify message", zap.Error(err))
		}
	}

	return nil
}
