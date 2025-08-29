package usecase

import (
	"context"
	"encoding/json"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/provider"
	"github.com/Hospital-Microservice/appointment-service/repository"
	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/hospital-core/model/req"
	"github.com/Hospital-Microservice/hospital-core/pointer"
	rabbit "github.com/Hospital-Microservice/hospital-core/provider"
	"go.uber.org/zap"
)

type AppointmentCreateUseCase interface {
	Execute(ctx context.Context, appointment *entity.AppointmentEntity) error
}

type appointmentCreateUseCaseImpl struct {
	appointmentRepo repository.AppointmentRepo
	publisher       *rabbit.RabbitPublisher
	UserService     provider.UserService
}

func (r appointmentCreateUseCaseImpl) Execute(ctx context.Context, appointment *entity.AppointmentEntity) error {
	appointment.Status = pointer.String("pending")
	if err := r.appointmentRepo.InsertAppointment(ctx, appointment); err != nil {
		log.Error("failed to insert appointment", zap.Error(err))
		return err
	}

	var ids []string
	if appointment.PatientID != nil {
		ids = append(ids, *appointment.PatientID)
	}
	if appointment.DoctorID != nil {
		ids = append(ids, *appointment.DoctorID)
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
		// Convert user IDs slice to JSON string for NotifyService
		idsBytes, _ := json.Marshal(ids)

		notify := req.NotificationReq{
			ToEmails: emails,
			Subject:  "Appointment Created",
			Body:     "Your appointment has been scheduled successfully.",
			Meta: map[string]any{
				"user_ids": string(idsBytes), // <-- fix here
			},
		}
		body, err := json.Marshal(notify)
		if err != nil {
			log.Error("failed to marshal notify message", zap.Error(err))
		} else if err := r.publisher.Publish(ctx, body); err != nil {
			log.Error("failed to publish notify message", zap.Error(err))
		}
	}

	return nil
}

func NewAppointmentCreateUseCase(
	AppointmentRepo repository.AppointmentRepo,
	publisher *rabbit.RabbitPublisher,
	UserService provider.UserService,
) AppointmentCreateUseCase {
	return &appointmentCreateUseCaseImpl{
		appointmentRepo: AppointmentRepo,
		publisher:       publisher,
		UserService:     UserService,
	}
}
