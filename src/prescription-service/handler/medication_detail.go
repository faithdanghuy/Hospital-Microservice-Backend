package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/prescription-service/mapper"
	"github.com/labstack/echo/v4"
)

// DetailMedication godoc
// @Summary      Detail Medication
// @Description  Detail a medication record
// @Tags         medication
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Medication ID"
// @Success      200  {object}  res.MedicationRes
// @Failure      400  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Security     BearerAuth
// @Router       /medication/detail/{id} [get]
func (h *prescriptionHandlerImpl) HandleDetailMedication(c echo.Context) error {
	id := c.Param("id")
	med, err := h.detailMedicationUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		return response.Error(c, http.StatusNotFound, err.Error())
	}
	return response.OK(c, http.StatusOK, "Success", mapper.TransformMedicationEntityToRes(med))
}
