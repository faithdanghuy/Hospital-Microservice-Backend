package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/mapper"
	"github.com/labstack/echo/v4"
)

// HandleGetPatients godoc
// @Summary      Get patients with pagination
// @Description  Doctors can see patients to make appointments
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
// @Router       /account/patients [get]
func (u *userHandlerImpl) HandleGetPatients(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized")
	}
	_, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Invalid Token")
	}
	p := new(record.Pagination)
	p.Limit = 999
	if err := c.Bind(p); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	result, err := u.filterUsersUseCase.Execute(c.Request().Context(), p, "patient")
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	var users []entity.UserEntity
	switch rows := result.Rows.(type) {
	case []entity.UserEntity:
		users = rows
	case []*entity.UserEntity:
		for _, u := range rows {
			if u != nil {
				users = append(users, *u)
			}
		}
	default:
		return response.Error(c, http.StatusInternalServerError, "Invalid result type")
	}

	resMapper := mapper.TransformUserEntitiesToRes(users)

	result.Rows = resMapper

	return response.OK(c, http.StatusOK, "Success", result)
}
