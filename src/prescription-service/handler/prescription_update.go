package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/prescription-service/mapper"
	"github.com/Hospital-Microservice/prescription-service/model/req"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// UpdatePrescription godoc
// @Summary      Update Prescription
// @Description  Update a prescription record
// @Tags         prescription
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Prescription ID"
// @Param        body  body      req.PrescriptionUpdateReq  true  "Prescription Update Request"
// @Success      200  {object}  response.ResOk
// @Failure      400  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Security     BearerAuth
// @Router       /prescription/update/{id} [patch]
func (h *prescriptionHandlerImpl) HandlePrescriptionUpdate(c echo.Context) error {
	var body req.PrescriptionUpdateReq
	if err := c.Bind(&body); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	if err := validator.New().Struct(&body); err != nil {
		return response.Errors(c, http.StatusBadRequest, err)
	}
	entity := mapper.TransformPrescriptionUpdateReqToEntity(body)
	if err := h.prescriptionUpdateUseCase.Execute(c.Request().Context(), entity); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.SimpleOK(c, http.StatusOK, nil)
}
