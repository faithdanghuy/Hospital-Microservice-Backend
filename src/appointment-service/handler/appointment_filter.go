package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/record"

	"github.com/Hospital-Microservice/appointment-service/entity"
	"github.com/Hospital-Microservice/appointment-service/model/req"
	"github.com/Hospital-Microservice/appointment-service/model/res"
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
// @Router       /appointment/filter [get]
func (u *appointmentHandlerImpl) HandleAppointmentFilter(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Invalid token")
	}

	p := new(record.Pagination)
	if err := c.Bind(p); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var filterReq req.AppointmentFilterReq
	if err := c.Bind(&filterReq); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid Query Params")
	}

	switch claims.AccountType {
	case "patient":
		filterReq.PatientID = claims.ID
	case "doctor":
		filterReq.DoctorID = claims.ID
	case "admin":
	default:
		return response.Error(c, http.StatusForbidden, "Unauthorized Role")
	}

	result, err := u.appointmentFilterUseCase.Execute(c.Request().Context(), p, filterReq)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	var appointments []*entity.AppointmentEntity
	switch rows := result.Rows.(type) {
	case []entity.AppointmentEntity:
		for i := range rows {
			appointments = append(appointments, &rows[i])
		}
	case []*entity.AppointmentEntity:
		appointments = rows
	default:
		return response.Error(c, http.StatusInternalServerError, "invalid result type")
	}

	var ids []string
	for _, appt := range appointments {
		if appt.PatientID != nil {
			ids = append(ids, *appt.PatientID)
		}
		if appt.DoctorID != nil {
			ids = append(ids, *appt.DoctorID)
		}
	}

	users, err := u.UserService.GetUsersByIDs(c.Request().Context(), ids, c.Request().Header.Get("Authorization"))
	if err != nil {
		c.Logger().Errorf("user service error: %v", err)
	}

	var enriched []res.AppointmentRes
	for _, appt := range appointments {
		r := res.AppointmentRes{
			ID:          *appt.ID,
			ScheduledAt: appt.ScheduledAt,
		}
		if appt.Status != nil {
			r.Status = *appt.Status
		}
		if appt.Note != nil {
			r.Note = *appt.Note
		}
		if appt.PatientID != nil {
			if u, ok := users[*appt.PatientID]; ok {
				r.Patient = &u
			}
		}
		if appt.DoctorID != nil {
			if u, ok := users[*appt.DoctorID]; ok {
				r.Doctor = &u
			}
		}
		enriched = append(enriched, r)
	}

	result.Rows = enriched
	return response.OK(c, http.StatusOK, "success", result)
}
