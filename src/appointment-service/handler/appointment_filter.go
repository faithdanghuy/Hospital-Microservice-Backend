package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"

	"github.com/Hospital-Microservice/appointment-service/mapper"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// GetAppointments godoc
// @Summary      Get appointments with filters
// @Description  Retrieve appointments with filters
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Param        filter  body      req.AppointmentFilterReq  true  "Appointment Filter Request"
// @Success      200  {object}  []interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Security     BearerAuth
// @Router       /appointment-service/appointments [get]
func (u *appointmentHandlerImpl) HandleAppointmentFilter(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "unauthorized")
	}

	claims, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "invalid token")
	}

	var filterReq req.AppointmentFilterReq
	if err := c.Bind(&filterReq); err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid query params")
	}

	filterReq.PatientID = claims.ID

	filterEntity := mapper.TransformAppointmentFilterReqToEntity(&filterReq)

	appointments, err := u.appointmentFilterUseCase.Execute(
		c.Request().Context(),
		filterEntity,
		filterReq.FromDate,
		filterReq.ToDate,
	)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	res := make([]interface{}, 0)
	for _, appt := range appointments {
		r := mapper.TransformAppointmentEntityToRes(appt)
		r.ID = *appt.ID
		res = append(res, r)
	}

	return response.OK(c, http.StatusOK, "success", res)
}
