package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/appointment-service/mapper"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (u *appointmentHandlerImpl) HandleAppointmentDetail(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "missing appointment id in path")
	}
	appointmentEntity, err := u.appointmentDetailUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	appointmentDetailRes := mapper.TransformAppointmentEntityToRes(appointmentEntity)
	appointmentDetailRes.ID = *appointmentEntity.ID

	return response.OK(c, http.StatusOK, "OK", appointmentDetailRes)
}
