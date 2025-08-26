package handler

import (
	"net/http"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// GetUserDetail godoc
// @Summary      Get user detail
// @Description  Retrieve detailed information for a specific user
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      403  {object}  response.ResErr
// @Failure      404  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /account/detail/{id} [get]
func (u *userHandlerImpl) HandleUserDetail(c echo.Context) error {
	// Get JWT claims
	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "unauthorized")
	}

	claims, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "invalid token")
	}

	// Get target user ID from path
	targetUserID := c.Param("id")
	if targetUserID == "" {
		return response.Error(c, http.StatusBadRequest, "user id is required")
	}

	// Get target user profile
	profile, err := u.profileUseCase.Execute(c.Request().Context(), targetUserID)
	if err != nil {
		return response.Error(c, http.StatusNotFound, "User Not Found")
	}

	switch claims.AccountType {
	case "patient":
		if claims.ID != targetUserID {
			return response.Error(c, http.StatusForbidden, "Unauthorized Role")
		}
	case "doctor":
		if profile.Role == nil || *profile.Role != "patient" {
			return response.Error(c, http.StatusForbidden, "Unauthorized Role")
		}
	case "admin":
	default:
		return response.Error(c, http.StatusForbidden, "Unauthorized Role")
	}

	profile.Password = nil

	return response.SimpleOK(c, http.StatusOK, profile)
}
