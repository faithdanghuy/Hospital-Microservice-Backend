package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/mapper"
	"github.com/Hospital-Microservice/prescription-service/model/req"
	"github.com/Hospital-Microservice/prescription-service/model/res"
	"github.com/Hospital-Microservice/prescription-service/provider"
	"github.com/labstack/echo/v4"
)

// HandlePrescriptionFilter godoc
// @Summary      Filter Prescriptions with pagination & filters
// @Description  Filter all prescriptions records with pagination & filters
// @Tags         prescription
// @Accept       json
// @Produce      json
// @Param        limit  query int    false "Limit per page"
// @Param        page   query int    false "Page number"
// @Param        sort   query string false "Sort (e.g. id asc)"
// @Param        patient_id query string false "Filter by patient ID"
// @Param        doctor_id query string false "Filter by doctor ID"
// @Param        status query string false "Filter by status"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /prescription/filter [get]
func (h *prescriptionHandlerImpl) HandlePrescriptionFilter(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Invalid token")
	}

	p := new(record.Pagination)
	if err := c.Bind(p); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	filter := new(req.PrescriptionFilterReq)
	if err := c.Bind(filter); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	switch claims.AccountType {
	case "patient":
		filter.PatientID = claims.ID

	case "doctor":
		filter.DoctorID = claims.ID
	case "admin":
	default:
		return response.Error(c, http.StatusForbidden, "Unauthorized Role")
	}

	result, err := h.prescriptionFilterUseCase.Execute(c.Request().Context(), p, filter)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	var prescriptions []*entity.PrescriptionEntity
	switch rows := result.Rows.(type) {
	case []entity.PrescriptionEntity:
		for i := range rows {
			prescriptions = append(prescriptions, &rows[i])
		}
	case []*entity.PrescriptionEntity:
		prescriptions = rows
	default:
		return response.Error(c, http.StatusInternalServerError, "invalid result type")
	}

	var ids []string
	for _, pr := range prescriptions {
		if pr.PatientID != nil {
			ids = append(ids, *pr.PatientID)
		}
		if pr.DoctorID != nil {
			ids = append(ids, *pr.DoctorID)
		}
	}

	users, err := h.UserService.GetUsersByIDs(c.Request().Context(), ids, c.Request().Header.Get("Authorization"))
	if err != nil {
		c.Logger().Errorf("user service error: %v", err)
	}

	convertedUsers := make(map[string]*provider.UserRes)
	for k, v := range users {
		vCopy := v
		convertedUsers[k] = &vCopy
	}

	var resList []*res.PrescriptionRes
	for _, pr := range prescriptions {
		resList = append(resList, mapper.TransformPrescriptionEntityToRes(pr, convertedUsers))
	}
	result.Rows = resList

	return response.OK(c, http.StatusOK, "Success", result)
}
