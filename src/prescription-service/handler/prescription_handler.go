package handler

import (
	"github.com/Hospital-Microservice/prescription-service/provider"
	"github.com/Hospital-Microservice/prescription-service/usecase"
	"github.com/labstack/echo/v4"
)

type PrescriptionHandler interface {
	HandlePrescriptionDetail(c echo.Context) error
	HandlePrescriptionCreate(c echo.Context) error
	HandlePrescriptionDelete(c echo.Context) error
	HandlePrescriptionUpdate(c echo.Context) error
	HandlePrescriptionFilter(c echo.Context) error

	HandleCreateMedication(c echo.Context) error
	HandleUpdateMedication(c echo.Context) error
	HandleListMedications(c echo.Context) error
	HandleDetailMedication(c echo.Context) error
	HandleDeleteMedication(c echo.Context) error
}

type prescriptionHandlerImpl struct {
	prescriptionFilterUseCase usecase.FilterPrescriptionUseCase
	prescriptionDetailUseCase usecase.GetPrescriptionUseCase
	prescriptionDeleteUseCase usecase.DeletePrescriptionUseCase
	prescriptionUpdateUseCase usecase.UpdatePrescriptionUseCase
	prescriptionCreateUseCase usecase.PrescriptionCreateUseCase
	deleteMedicationUseCase   usecase.DeleteMedicationUseCase
	createMedicationUseCase   usecase.CreateMedicationUseCase
	updateMedicationUseCase   usecase.UpdateMedicationUseCase
	listMedicationUseCase     usecase.ListMedicationUseCase
	detailMedicationUseCase   usecase.DetailMedicationUseCase
	UserService               provider.UserService
}

type PrescriptionHandlerInject struct {
	FilterPrescriptionUseCase usecase.FilterPrescriptionUseCase
	PrescriptionDetailUseCase usecase.GetPrescriptionUseCase
	DeletePrescriptionUseCase usecase.DeletePrescriptionUseCase
	UpdatePrescriptionUseCase usecase.UpdatePrescriptionUseCase
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
		prescriptionFilterUseCase: inject.FilterPrescriptionUseCase,
		prescriptionDetailUseCase: inject.PrescriptionDetailUseCase,
		prescriptionDeleteUseCase: inject.DeletePrescriptionUseCase,
		prescriptionUpdateUseCase: inject.UpdatePrescriptionUseCase,
		prescriptionCreateUseCase: inject.PrescriptionCreateUseCase,
		deleteMedicationUseCase:   inject.DeleteMedicationUseCase,
		createMedicationUseCase:   inject.CreateMedicationUseCase,
		updateMedicationUseCase:   inject.UpdateMedicationUseCase,
		listMedicationUseCase:     inject.ListMedicationUseCase,
		detailMedicationUseCase:   inject.DetailMedicationUseCase,
		UserService:               inject.UserService,
	}
}
