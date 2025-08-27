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
// @Router       /appointment/change-status/{id} [patch]
func (u *appointmentHandlerImpl) HandleAppointmentChangeStatus(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "Missing Appointment ID")
	}

	var req req.AppointmentChangeStatusReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid Request Body")
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

	return response.OK(c, http.StatusOK, "Success", appointmentRes)
}
