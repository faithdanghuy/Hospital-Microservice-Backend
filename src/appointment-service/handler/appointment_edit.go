package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/appointment-service/mapper"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// EditAppointment godoc
// @Summary      Edit appointment
// @Description  Edit an appointment by ID
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Param        id   path      string                 true  "Appointment ID"
// @Param        body body      req.AppointmentEditReq true  "Edit Appointment Request"
// @Success      200  {object} response.ResOk
// @Failure      400  {object} response.ResErr
// @Failure      404  {object} response.ResErr
// @Security     BearerAuth
// @Router       /appointment/edit/{id} [patch]
func (u *appointmentHandlerImpl) HandleAppointmentEdit(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "Missing Appointment ID")
	}

	var req req.AppointmentEditReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body")
	}

	appointmentEntity := mapper.TransformAppointmentEditReqToEntity(&req)

	updatedEntity, err := u.appointmentEditUseCase.Execute(c.Request().Context(), id, appointmentEntity)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var ids []string
	if updatedEntity.PatientID != nil {
		ids = append(ids, *updatedEntity.PatientID)
	}
	if updatedEntity.DoctorID != nil {
		ids = append(ids, *updatedEntity.DoctorID)
	}

	users, err := u.UserService.GetUsersByIDs(c.Request().Context(), ids, c.Request().Header.Get("Authorization"))
	if err != nil {
		c.Logger().Errorf("user service error: %v", err)
	}

	appointmentRes := mapper.TransformAppointmentEntityToDetailRes(updatedEntity, users)

	return response.OK(c, http.StatusOK, "Success", appointmentRes)
}
