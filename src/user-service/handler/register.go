package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/user-service/mapper"
	"github.com/Hospital-Microservice/user-service/model/req"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (u *userHandlerImpl) HandleRegister(c echo.Context) error {
	var user req.UserRegReq
	err := c.Bind(&user)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&user); err != nil {
		return response.Errors(c, http.StatusBadRequest, err)
	}

	err = u.registerUseCase.Execute(
		c.Request().Context(),
		mapper.TransformRegReqToEntity(user),
	)

	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.SimpleOK(c, http.StatusOK, nil)
}
