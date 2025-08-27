package handler

import (
	"github.com/Hospital-Microservice/appointment-service/provider"
	"github.com/Hospital-Microservice/appointment-service/usecase"
	"github.com/labstack/echo/v4"
)

type AppointmentHandler interface {
	HandleAppointmentDetail(c echo.Context) error
	HandleAppointmentCreate(c echo.Context) error
	HandleAppointmentChangeStatus(c echo.Context) error
	HandleAppointmentFilter(c echo.Context) error
	HandleAppointmentEdit(c echo.Context) error
}

type appointmentHandlerImpl struct {
	appointmentDetailUseCase       usecase.AppointmentDetailUseCase
	appointmentCreateUseCase       usecase.AppointmentCreateUseCase
	appointmentChangeStatusUseCase usecase.AppointmentChangeStatusUseCase
	appointmentFilterUseCase       usecase.AppointmentFilterUseCase
	appointmentEditUseCase         usecase.AppointmentEditUseCase
	UserService                    provider.UserService
}

type AppointmentHandlerInject struct {
	AppointmentDetailUseCase       usecase.AppointmentDetailUseCase
	AppointmentCreateUseCase       usecase.AppointmentCreateUseCase
	AppointmentChangeStatusUseCase usecase.AppointmentChangeStatusUseCase
	AppointmentFilterUseCase       usecase.AppointmentFilterUseCase
	AppointmentEditUseCase         usecase.AppointmentEditUseCase
	UserService                    provider.UserService
}

func NewAppointmentHandler(inject AppointmentHandlerInject) AppointmentHandler {
	return &appointmentHandlerImpl{
		appointmentDetailUseCase:       inject.AppointmentDetailUseCase,
		appointmentCreateUseCase:       inject.AppointmentCreateUseCase,
		appointmentChangeStatusUseCase: inject.AppointmentChangeStatusUseCase,
		appointmentFilterUseCase:       inject.AppointmentFilterUseCase,
		appointmentEditUseCase:         inject.AppointmentEditUseCase,
		UserService:                    inject.UserService,
	}
}
