package handler

import (
	"github.com/Hospital-Microservice/appointment-service/usecase"
	"github.com/labstack/echo/v4"
)

type AppointmentHandler interface {
	HandleAppointmentDetail(c echo.Context) error
	HandleAppointmentCreate(c echo.Context) error
}

type appointmentHandlerImpl struct {
	appointmentDetailUseCase usecase.AppointmentDetailUseCase
	appointmentCreateUseCase usecase.AppointmentCreateUseCase
}

type AppointmentHandlerInject struct {
	AppointmentDetailUseCase usecase.AppointmentDetailUseCase
	AppointmentCreateUseCase usecase.AppointmentCreateUseCase
}

func NewAppointmentHandler(inject AppointmentHandlerInject) AppointmentHandler {
	return &appointmentHandlerImpl{
		appointmentDetailUseCase: inject.AppointmentDetailUseCase,
		appointmentCreateUseCase: inject.AppointmentCreateUseCase,
	}
}
