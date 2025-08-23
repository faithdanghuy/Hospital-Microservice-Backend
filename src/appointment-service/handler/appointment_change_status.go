package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/appointment-service/mapper"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// ChangeAppointmentStatus godoc
// @Summary      Change appointment status
// @Description  Change the status of an appointment by ID
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Appointment ID"
// @Param        body  body      req.AppointmentChangeStatusReq  true  "Appointment Change Status Request"
// @Success      200  {object}  response.ResOk
// @Failure      404  {object}  response.ResErr
// @Security     BearerAuth
// @Router       /appointment-service/appointments/{id}/status [patch]
func (u *appointmentHandlerImpl) HandleAppointmentChangeStatus(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "missing appointment id in path")
	}

	var req req.AppointmentChangeStatusReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid request body")
	}

	appointmentEntity := mapper.TransformAppointmentChangeStatusReqToEntity(&req)
	appointmentEntity.ID = &id

	updatedEntity, err := u.appointmentChangeStatusUseCase.Execute(
		c.Request().Context(),
		id,
		*appointmentEntity.Status,
	)

	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	appointmentRes := mapper.TransformAppointmentEntityToRes(updatedEntity)
	appointmentRes.ID = *updatedEntity.ID

	return response.OK(c, http.StatusOK, "appointment status changed successfully", appointmentRes)
}
