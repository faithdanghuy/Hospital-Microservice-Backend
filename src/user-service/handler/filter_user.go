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

// HandleFilterUsers godoc
// @Summary      Filter users with pagination
// @Description  Filter users with pagination (only admin can access)
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        limit   query  int    false  "Limit"
// @Param        page    query  int    false  "Page"
// @Param        sort    query  string false  "Sort"
// @Param        role    query  string false  "Role filter (admin|doctor|patient)"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      403  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /account/filter [get]
func (u *userHandlerImpl) HandleFilterUsers(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Invalid Token")
	}
	if claims.AccountType != "admin" {
		return response.Error(c, http.StatusForbidden, "Forbidden")
	}

	p := new(record.Pagination)
	if err := c.Bind(p); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	role := c.QueryParam("role")

	result, err := u.filterUsersUseCase.Execute(c.Request().Context(), p, role)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	users, ok := result.Rows.([]entity.UserEntity)
	if !ok {
		return response.Error(c, http.StatusInternalServerError, "invalid result type")
	}

	resUsers := mapper.TransformUserEntitiesToRes(users)
	result.Rows = resUsers
	return response.OK(c, http.StatusOK, "Success", result)
}
