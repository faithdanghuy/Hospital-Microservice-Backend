package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/prescription-service/mapper"
	"github.com/Hospital-Microservice/prescription-service/model/req"
	"github.com/labstack/echo/v4"
)

// UpdateMedication godoc
// @Summary      Update Medication
// @Description  Update a medication record
// @Tags         medication
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Medication ID"
// @Param        body  body      req.MedicationUpdateReq  true  "Medication Update Request"
// @Success      200  {object}  res.MedicationRes
// @Failure      400  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Security     BearerAuth
// @Router       /medication/update/{id} [patch]
func (h *prescriptionHandlerImpl) HandleUpdateMedication(c echo.Context) error {
	id := c.Param("id")
	var req req.MedicationUpdateReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	med, err := h.detailMedicationUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		return response.Error(c, http.StatusNotFound, err.Error())
	}
	if req.DrugName != nil {
		med.DrugName = req.DrugName
	}
	if req.Stock != nil {
		med.Stock = req.Stock
	}
	if req.Unit != nil {
		med.Unit = req.Unit
	}
	if req.Description != nil {
		med.Description = req.Description
	}
	if err := h.updateMedicationUseCase.Execute(c.Request().Context(), med); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.SimpleOK(c, http.StatusOK, mapper.TransformMedicationEntityToRes(med))
}
