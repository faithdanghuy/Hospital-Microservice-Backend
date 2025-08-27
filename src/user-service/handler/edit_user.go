package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/user-service/mapper"
	"github.com/Hospital-Microservice/user-service/model/req"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// EditUser godoc
// @Summary      Edit any user (Admin only)
// @Description  Update user data by admin
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        id  path      string             true  "User ID"
// @Param        body     body      req.UserUpdateReq  true  "Update Request"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      403  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /account/edit/{id} [patch]
func (u *userHandlerImpl) HandleEditUser(c echo.Context) error {
	ut := c.Get("user")
	if ut == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := ut.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Invalid Token")
	}
	if claims.AccountType != "admin" {
		return response.Error(c, http.StatusForbidden, "Forbidden")
	}

	var user req.UserUpdateReq
	if err := c.Bind(&user); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&user); err != nil {
		return response.Errors(c, http.StatusBadRequest, err)
	}

	userID := c.Param("id")

	err := u.editUserUseCase.Execute(
		c.Request().Context(),
		mapper.TransformUpdateReqToEntity(userID, user),
	)

	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.SimpleOK(c, http.StatusOK, nil)
}
