package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/user-service/mapper"
	"github.com/labstack/echo/v4"
)

// Profile godoc
// @Summary      Get current user profile
// @Description  Retrieve profile for the authenticated user (from JWT claims)
// @Tags         account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /account/profile [get]
func (u *userHandlerImpl) HandleProfile(c echo.Context) error {
	claims := c.Get("user").(token.JwtCustomClaims)
	if claims.ID == "" {
		return response.Error(c, http.StatusBadRequest)
	}

	profile, err := u.profileUseCase.Execute(c.Request().Context(), claims.ID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	profile.Password = nil

	resUsers := mapper.TransformUserEntityToFilterRes(profile)

	return response.SimpleOK(c, http.StatusOK, resUsers)
}
