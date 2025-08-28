package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/mapper"
	"github.com/Hospital-Microservice/prescription-service/model/req"
	"github.com/labstack/echo/v4"
)

// CreateMedication godoc
// @Summary      Create Medication
// @Description  Create a new medication record
// @Tags         medication
// @Accept       json
// @Produce      json
// @Param        body  body      req.MedicationCreateReq  true  "Medication Create Request"
// @Success      201   {object}  response.ResOk
// @Failure      400   {object}  response.ResErr
// @Security     BearerAuth
// @Router       /medication/create [post]
func (h *prescriptionHandlerImpl) HandleCreateMedication(c echo.Context) error {
	var req req.MedicationCreateReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	med := &entity.MedicationEntity{
		DrugName:    req.DrugName,
		Stock:       req.Stock,
		Unit:        req.Unit,
		Description: req.Description,
	}
	if err := h.createMedicationUseCase.Execute(c.Request().Context(), med); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.SimpleOK(c, http.StatusCreated, mapper.TransformMedicationEntityToRes(med))
}
