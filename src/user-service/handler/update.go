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

// UpdateUser godoc
// @Summary      Update current user
// @Description  Update profile data for the authenticated user
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        body  body      req.UserUpdateReq  true  "Update Request"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /account/update [patch]
func (u *userHandlerImpl) HandleUpdate(c echo.Context) error {
	var user req.UserUpdateReq
	err := c.Bind(&user)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&user); err != nil {
		return response.Errors(c, http.StatusBadRequest, err)
	}

	claims := c.Get("user").(token.JwtCustomClaims)
	userID := claims.ID

	err = u.updateUseCase.Execute(
		c.Request().Context(),
		mapper.TransformUpdateReqToEntity(userID, user),
	)

	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.SimpleOK(c, http.StatusOK, nil)
}
