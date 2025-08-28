package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// DeleteMedication godoc
// @Summary      Delete Medication
// @Description  Delete a medication record
// @Tags         medication
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Medication ID"
// @Success      200  {object}  response.ResOk
// @Failure      400  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Security     BearerAuth
// @Router       /medication/delete/{id} [delete]
func (h *prescriptionHandlerImpl) HandleDeleteMedication(c echo.Context) error {
	id := c.Param("id")
	if err := h.deleteMedicationUseCase.Execute(c.Request().Context(), id); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.SimpleOK(c, http.StatusOK, nil)
}
