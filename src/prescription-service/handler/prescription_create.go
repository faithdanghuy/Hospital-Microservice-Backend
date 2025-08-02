package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/prescription-service/mapper"
	"github.com/Hospital-Microservice/prescription-service/model/req"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (u *prescriptionHandlerImpl) HandlePrescriptionCreate(c echo.Context) error {
	var prescription req.PrescriptionCreateReq
	err := c.Bind(&prescription)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&prescription); err != nil {
		return response.Errors(c, http.StatusBadRequest, err)
	}

	err = u.prescriptionCreateUseCase.Execute(
		c.Request().Context(),
		mapper.TransformPrescriptionCreateReqToEntity(prescription),
	)

	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.SimpleOK(c, http.StatusOK, nil)
}
