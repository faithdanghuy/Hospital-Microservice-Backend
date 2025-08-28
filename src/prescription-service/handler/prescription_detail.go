package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/prescription-service/mapper"
	"github.com/Hospital-Microservice/prescription-service/provider"
	"github.com/labstack/echo/v4"
)

// HandlePrescriptionDetail godoc
// @Summary      Get Prescription Detail
// @Description  Get a prescription detail record
// @Tags         prescription
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Prescription ID"
// @Success      200  {object}  res.PrescriptionRes
// @Failure      400  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Security     BearerAuth
// @Router       /prescription/detail/{id} [get]
func (h *prescriptionHandlerImpl) HandlePrescriptionDetail(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "missing prescription id in path")
	}

	prescription, err := h.prescriptionDetailUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		return response.Error(c, http.StatusNotFound, err.Error())
	}

	var ids []string
	if prescription.PatientID != nil {
		ids = append(ids, *prescription.PatientID)
	}
	if prescription.DoctorID != nil {
		ids = append(ids, *prescription.DoctorID)
	}

	users, err := h.UserService.GetUsersByIDs(
		c.Request().Context(),
		ids,
		c.Request().Header.Get("Authorization"),
	)
	if err != nil {
		c.Logger().Errorf("user service error: %v", err)
	}
	convertedUsers := make(map[string]*provider.UserRes)
	for k, v := range users {
		vCopy := v
		convertedUsers[k] = &vCopy
	}

	res := mapper.TransformPrescriptionEntityToRes(prescription, convertedUsers)
	return response.OK(c, http.StatusOK, "Success", res)
}
