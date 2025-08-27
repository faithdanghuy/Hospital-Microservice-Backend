package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/appointment-service/mapper"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// GetAppointment godoc
// @Summary      Get appointment details
// @Description  Retrieve details of an appointment by ID
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Appointment ID"
// @Success      200  {object}  res.AppointmentDetailRes
// @Failure      404  {object}  response.ResErr
// @Security     BearerAuth
// @Router       /appointment/detail/{id} [get]
func (u *appointmentHandlerImpl) HandleAppointmentDetail(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "missing appointment id in path")
	}

	appointmentEntity, err := u.appointmentDetailUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var ids []string
	if appointmentEntity.PatientID != nil {
		ids = append(ids, *appointmentEntity.PatientID)
	}
	if appointmentEntity.DoctorID != nil {
		ids = append(ids, *appointmentEntity.DoctorID)
	}

	users, err := u.UserService.GetUsersByIDs(c.Request().Context(), ids, c.Request().Header.Get("Authorization"))
	if err != nil {
		c.Logger().Errorf("user service error: %v", err)
	}

	appointmentDetailRes := mapper.TransformAppointmentEntityToDetailRes(appointmentEntity, users)

	return response.OK(c, http.StatusOK, "OK", appointmentDetailRes)
}
