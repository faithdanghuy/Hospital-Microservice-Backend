package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/user-service/mapper"
	"github.com/labstack/echo/v4"
)

// BatchRequest represents the request body for /account/batch
type BatchRequest struct {
	IDs []string `json:"ids"`
}

// HandleAccountBatch godoc
// @Summary      Get multiple user details by ids
// @Description  Return list of users for given ids (batch)
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        body  body      BatchRequest  true  "IDs"
// @Success      200  {object}  []interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      401  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /account/batch [post]
// HandleAccountBatch godoc
func (u *userHandlerImpl) HandleAccountBatch(c echo.Context) error {
	var req BatchRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid body")
	}

	if len(req.IDs) == 0 {
		return response.OK(c, http.StatusOK, "Success", []interface{}{})
	}

	users, err := u.getBatchUseCase.Execute(c.Request().Context(), req.IDs)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "failed to fetch users")
	}

	out := make([]interface{}, 0, len(users))
	for _, user := range users {
		out = append(out, mapper.TransformUserEntityToRes(user))
	}

	return response.OK(c, http.StatusOK, "Success", out)
}
