package handler

import (
	"net/http"
	"time"

	"github.com/Hospital-Microservice/hospital-core/security"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/user-service/mapper"
	"github.com/Hospital-Microservice/user-service/model/req"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// LoginUser godoc
// @Summary      Login user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      req.UserLoginReq  true  "Login Request"
// @Success 200 {object} res.LoginRes
// @Failure      400 {object} response.ResErr
// @Router       /auth/login [post]
func (u *userHandlerImpl) HandleLogin(c echo.Context) error {
	var user req.UserLoginReq
	err := c.Bind(&user)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&user); err != nil {
		return response.Errors(c, http.StatusBadRequest, err)
	}

	userEntity, err := u.loginUseCase.Execute(c.Request().Context(), user.Phone, user.Password)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	accessToken, err := security.GenToken(*userEntity.ID, *userEntity.Role, time.Hour*24)
	if err != nil {
		return err
	}
	loginRes := mapper.TransformUserEntityToRes(userEntity)
	loginRes.ID = *userEntity.ID
	loginRes.AccessToken = *accessToken

	return response.OK(
		c, http.StatusOK,
		"OK",
		loginRes,
	)
}
