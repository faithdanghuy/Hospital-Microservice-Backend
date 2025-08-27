package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleGetDoctors godoc
// @Summary      Get doctors with pagination
// @Description  Patients can see available doctors to make appointments
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        limit   query  int    false  "Limit"
// @Param        page    query  int    false  "Page"
// @Param        sort    query  string false  "Sort"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      403  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /account/doctors [get]
func (u *userHandlerImpl) HandleGetDoctors(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized")
	}
	_, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Invalid Token")
	}

	p := new(record.Pagination)
	if err := c.Bind(p); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	result, err := u.filterUsersUseCase.Execute(c.Request().Context(), p, "doctor")
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.OK(c, http.StatusOK, "Success", result)
}
