package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// DeleteUser godoc
// @Summary      Delete user (admin only)
// @Description  Soft delete a user by ID (flag as deleted, not permanent)
// @Tags         account
// @Produce      json
// @Param        user-id   path      string  true  "User ID"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      403  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /account/delete/{id} [delete]
func (u *userHandlerImpl) HandleDeleteUser(c echo.Context) error {
	claims := c.Get("user").(token.JwtCustomClaims)

	if claims.AccountType != "admin" {
		return response.Error(c, http.StatusForbidden, "Permission Denied")
	}

	userID := c.Param("id")
	if userID == "" {
		return response.Error(c, http.StatusBadRequest, "Missing User ID")
	}

	err := u.deleteUserUseCase.Execute(c.Request().Context(), userID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.SimpleOK(c, http.StatusOK, "User Deleted")
}
