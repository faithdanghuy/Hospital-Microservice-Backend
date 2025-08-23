package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"

	"github.com/Hospital-Microservice/appointment-service/mapper"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (u *appointmentHandlerImpl) HandleAppointmentFilter(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "missing token")
	}

	claims, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "invalid claims")
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
