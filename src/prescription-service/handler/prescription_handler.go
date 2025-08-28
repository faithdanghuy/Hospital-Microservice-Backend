package handler

import (
	"github.com/Hospital-Microservice/prescription-service/provider"
	"github.com/Hospital-Microservice/prescription-service/usecase"
	"github.com/labstack/echo/v4"
)

type PrescriptionHandler interface {
	HandlePrescriptionDetail(c echo.Context) error
	HandlePrescriptionCreate(c echo.Context) error
	HandleCreateMedication(c echo.Context) error
	HandleUpdateMedication(c echo.Context) error
	HandleListMedications(c echo.Context) error
	HandleDetailMedication(c echo.Context) error
	HandleDeleteMedication(c echo.Context) error
}

type prescriptionHandlerImpl struct {
	prescriptionDetailUseCase usecase.PrescriptionDetailUseCase
	prescriptionCreateUseCase usecase.PrescriptionCreateUseCase
	deleteMedicationUseCase   usecase.DeleteMedicationUseCase
	createMedicationUseCase   usecase.CreateMedicationUseCase
	updateMedicationUseCase   usecase.UpdateMedicationUseCase
	listMedicationUseCase     usecase.ListMedicationUseCase
	detailMedicationUseCase   usecase.DetailMedicationUseCase
	UserService               provider.UserService
}

type PrescriptionHandlerInject struct {
	PrescriptionDetailUseCase usecase.PrescriptionDetailUseCase
	PrescriptionCreateUseCase usecase.PrescriptionCreateUseCase
	DeleteMedicationUseCase   usecase.DeleteMedicationUseCase
	CreateMedicationUseCase   usecase.CreateMedicationUseCase
	UpdateMedicationUseCase   usecase.UpdateMedicationUseCase
	ListMedicationUseCase     usecase.ListMedicationUseCase
	DetailMedicationUseCase   usecase.DetailMedicationUseCase
	UserService               provider.UserService
}

func NewPrescriptionHandler(inject PrescriptionHandlerInject) PrescriptionHandler {
	return &prescriptionHandlerImpl{
		prescriptionDetailUseCase: inject.PrescriptionDetailUseCase,
		prescriptionCreateUseCase: inject.PrescriptionCreateUseCase,
		deleteMedicationUseCase:   inject.DeleteMedicationUseCase,
		createMedicationUseCase:   inject.CreateMedicationUseCase,
		updateMedicationUseCase:   inject.UpdateMedicationUseCase,
		listMedicationUseCase:     inject.ListMedicationUseCase,
		detailMedicationUseCase:   inject.DetailMedicationUseCase,
		UserService:               inject.UserService,
	}
}
