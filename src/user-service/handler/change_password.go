package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/user-service/model/req"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ChangePassword godoc
// @Summary      Change password
// @Description  Update password for the authenticated user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      req.UserChangePasswordReq  true  "Change Password Request"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /auth/change-password [patch]
func (u *userHandlerImpl) HandleChangePassword(c echo.Context) error {
	var reqBody req.UserChangePasswordReq
	if err := c.Bind(&reqBody); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&reqBody); err != nil {
		return response.Errors(c, http.StatusBadRequest, err)
	}

	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Invalid Token")
	}

	userID := claims.ID

	if err := u.changePwdUseCase.Execute(
		c.Request().Context(),
		userID,
		reqBody.OldPassword,
		reqBody.NewPassword,
	); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	return response.SimpleOK(c, http.StatusOK, nil)
}
