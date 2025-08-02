package handler

import (
	"github.com/Hospital-Microservice/prescription-service/usecase"
	"github.com/labstack/echo/v4"
)

type PrescriptionHandler interface {
	HandlePrescriptionDetail(c echo.Context) error
	HandlePrescriptionCreate(c echo.Context) error
}

type prescriptionHandlerImpl struct {
	prescriptionDetailUseCase usecase.PrescriptionDetailUseCase
	prescriptionCreateUseCase usecase.PrescriptionCreateUseCase
}

type PrescriptionHandlerInject struct {
	PrescriptionDetailUseCase usecase.PrescriptionDetailUseCase
	PrescriptionCreateUseCase usecase.PrescriptionCreateUseCase
}

func NewPrescriptionHandler(inject PrescriptionHandlerInject) PrescriptionHandler {
	return &prescriptionHandlerImpl{
		prescriptionDetailUseCase: inject.PrescriptionDetailUseCase,
		prescriptionCreateUseCase: inject.PrescriptionCreateUseCase,
	}
}
