package usecase

import (
	"context"
	"encoding/json"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/repository"
	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/hospital-core/model/req"
	"github.com/Hospital-Microservice/hospital-core/pointer"
	"github.com/Hospital-Microservice/hospital-core/provider"
	"go.uber.org/zap"
)

type AppointmentCreateUseCase interface {
	Execute(ctx context.Context, appointment *entity.AppointmentEntity) error
}

type appointmentCreateUseCaseImpl struct {
	appointmentRepo repository.AppointmentRepo
	publisher       *provider.RabbitPublisher
}

func (r appointmentCreateUseCaseImpl) Execute(ctx context.Context, appointment *entity.AppointmentEntity) error {

	appointment.Status = pointer.String("pending")
	if err := r.appointmentRepo.InsertAppointment(ctx, appointment); err != nil {
		log.Error("failed to insert appointment", zap.Error(err))
		return err
	}

	notify := req.NotificationReq{
		ToEmails: []string{"hihuu456@gmail.com"},
		Subject:  "Appointment Created",
		Body:     "Your appointment has been scheduled successfully.",
	}
	body, _ := json.Marshal(notify)
	if err := r.publisher.Publish(ctx, body); err != nil {
		log.Error("failed to publish notify message", zap.Error(err))
	}
	log.Info("published notify message", zap.String("body", string(body)))
	return nil
}

func NewAppointmentCreateUseCase(
	AppointmentRepo repository.AppointmentRepo,
	publisher *provider.RabbitPublisher,
) AppointmentCreateUseCase {
	return &appointmentCreateUseCaseImpl{
		appointmentRepo: AppointmentRepo,
		publisher:       publisher,
	}
}
