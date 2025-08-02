package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/prescription-service/mapper"
	"github.com/labstack/echo/v4"
)

func (u *prescriptionHandlerImpl) HandlePrescriptionDetail(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "missing prescription id in path")
	}
	prescriptionEntity, err := u.prescriptionDetailUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	prescriptionDetailRes := mapper.TransformPrescriptionEntityToRes(prescriptionEntity)
	prescriptionDetailRes.ID = *prescriptionEntity.ID

	return response.OK(c, http.StatusOK, "OK", prescriptionDetailRes)
}
