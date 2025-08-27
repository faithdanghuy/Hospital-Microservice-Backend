package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/appointment-service/mapper"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CreateAppointment godoc
// @Summary      Create appointment
// @Description  Create a new appointment
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Param        body  body      req.AppointmentCreateReq  true  "Appointment Create Request"
// @Success      201   {object}  response.ResOk
// @Failure      400   {object}  response.ResErr
// @Security     BearerAuth
// @Router       /appointment/create [post]
func (u *appointmentHandlerImpl) HandleAppointmentCreate(c echo.Context) error {
	var appointment req.AppointmentCreateReq
	err := c.Bind(&appointment)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&appointment); err != nil {
		return response.Errors(c, http.StatusBadRequest, err)
	}

	err = u.appointmentCreateUseCase.Execute(
		c.Request().Context(),
		mapper.TransformAppointmentCreateReqToEntity(appointment),
	)

	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.SimpleOK(c, http.StatusOK, nil)
}
