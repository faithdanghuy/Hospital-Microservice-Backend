package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

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
	return response.SimpleOK(c, http.StatusOK, profile)
}
