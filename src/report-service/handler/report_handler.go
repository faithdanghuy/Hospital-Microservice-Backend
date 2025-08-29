package handler

import (
	"net/http"
	"strconv"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/report-service/usecase"
	"github.com/labstack/echo/v4"
)

// ReportHandler holds usecases
type ReportHandler struct {
	PatientsUC      *usecase.PatientsReportUseCase
	AppointmentsUC  *usecase.AppointmentsReportUseCase
	PrescriptionsUC *usecase.PrescriptionsReportUseCase
}

type ReportHandlerInject struct {
	PatientsUC      *usecase.PatientsReportUseCase
	AppointmentsUC  *usecase.AppointmentsReportUseCase
	PrescriptionsUC *usecase.PrescriptionsReportUseCase
}

func NewReportHandler(in ReportHandlerInject) *ReportHandler {
	return &ReportHandler{
		PatientsUC:      in.PatientsUC,
		AppointmentsUC:  in.AppointmentsUC,
		PrescriptionsUC: in.PrescriptionsUC,
	}
}

// Get patients report
// @Summary Patients per day in month
// @Tags report
// @Produce json
// @Param month query int true "Month (1-12)"
// @Param year  query int true "Year"
// @Security BearerAuth
// @Success 200 {object} usecase.ChartData
// @Failure 400 {object} response.ResErr
// @Router /report/patients [get]
func (h *ReportHandler) HandlePatients(c echo.Context) error {
	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.Error(c, http.StatusBadRequest, "invalid month")
	}
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid year")
	}
	auth := c.Request().Header.Get("Authorization")
	data, err := h.PatientsUC.Execute(c.Request().Context(), month, year, auth)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.OK(c, http.StatusOK, "success", data)
}

// Get appointments report
// @Summary Appointments per day for doctor/month
// @Tags report
// @Produce json
// @Param doctor_id query string false "Doctor ID"
// @Param month     query int true "Month (1-12)"
// @Param year      query int true "Year"
// @Security BearerAuth
// @Success 200 {object} usecase.ChartData
// @Failure 400 {object} response.ResErr
// @Router /report/appointments [get]
func (h *ReportHandler) HandleAppointments(c echo.Context) error {
	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.Error(c, http.StatusBadRequest, "invalid month")
	}
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid year")
	}
	doctorID := c.QueryParam("doctor_id")
	auth := c.Request().Header.Get("Authorization")
	data, err := h.AppointmentsUC.Execute(c.Request().Context(), doctorID, month, year, auth)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.OK(c, http.StatusOK, "success", data)
}

// Get prescriptions report
// @Summary Prescriptions per day in month
// @Tags report
// @Produce json
// @Param month query int true "Month (1-12)"
// @Param year  query int true "Year"
// @Security BearerAuth
// @Success 200 {object} usecase.ChartData
// @Failure 400 {object} response.ResErr
// @Router /report/prescriptions [get]
func (h *ReportHandler) HandlePrescriptions(c echo.Context) error {
	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.Error(c, http.StatusBadRequest, "invalid month")
	}
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid year")
	}
	auth := c.Request().Header.Get("Authorization")
	data, err := h.PrescriptionsUC.Execute(c.Request().Context(), month, year, auth)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.OK(c, http.StatusOK, "success", data)
}
