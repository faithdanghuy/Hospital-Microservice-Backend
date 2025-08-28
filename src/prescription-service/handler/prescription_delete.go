package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandlePrescriptionDelete godoc
// @Summary      Delete Prescription
// @Description  Delete a prescription record
// @Tags         prescription
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Prescription ID"
// @Success      200  {object}  response.ResOk
// @Failure      400  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Security     BearerAuth
// @Router       /prescription/delete/{id} [delete]
func (h *prescriptionHandlerImpl) HandlePrescriptionDelete(c echo.Context) error {
	id := c.Param("id")
	if err := h.prescriptionDeleteUseCase.Execute(c.Request().Context(), id); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.SimpleOK(c, http.StatusOK, nil)
}
