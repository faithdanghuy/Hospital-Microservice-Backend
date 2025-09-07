package handler

import (
	"net/http"
	"strconv"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"

	token "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/notify-service/model/req"
	"github.com/Hospital-Microservice/notify-service/usecase"
	"github.com/labstack/echo/v4"
)

type NotifyHandler interface {
	HandleSend(c echo.Context) error
	HandleHealth(c echo.Context) error
	HandleListByUser(c echo.Context) error
	HandleMarkRead(c echo.Context) error
}

type Inject struct {
	NotifyUC usecase.NotifyUseCase
}

type notifyHandlerImpl struct {
	notif usecase.NotifyUseCase
}

func NewNotifyHandler(in Inject) NotifyHandler {
	return &notifyHandlerImpl{
		notif: in.NotifyUC,
	}
}

// HandleSend godoc
// @Summary Send notification (email/sms)
// @Param body body req.NotificationReq true "Notification"
// @Tags notify
// @Success 200 {object} response.ResOk
// @Failure 400 {object} response.ResErr
// @Router /notify/send [post]
func (h *notifyHandlerImpl) HandleSend(c echo.Context) error {
	var r req.NotificationReq
	if err := c.Bind(&r); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	if err := h.notif.Send(c.Request().Context(), r); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.SimpleOK(c, http.StatusOK, nil)
}

func (h *notifyHandlerImpl) HandleHealth(c echo.Context) error {
	return response.OK(c, http.StatusOK, "ok", map[string]string{"service": "notify-service"})
}

// HandleListByUser godoc
// @Summary Get notifications for a user
// @Tags notify
// @Param user_id path string true "User ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} response.ResOk
// @Failure 400 {object} response.ResErr
// @Router /notify/notification [get]
func (h *notifyHandlerImpl) HandleListByUser(c echo.Context) error {
	// Use token user id as authoritative identity
	user := c.Get("user")
	if user == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := user.(token.JwtCustomClaims)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Invalid token")
	}
	if claims.ID == "" {
		return response.Error(c, http.StatusUnauthorized, "Invalid token: missing user ID")
	}

	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	offset, _ := strconv.ParseInt(c.QueryParam("offset"), 10, 64)

	list, err := h.notif.ListByUser(c.Request().Context(), claims.ID, limit, offset)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	out := make([]interface{}, 0, len(list))
	for _, n := range list {
		out = append(out, map[string]interface{}{
			"id":         n.ID,
			"user_id":    n.UserID,
			"title":      n.Title,
			"body":       n.Body,
			"channel":    n.Channel,
			"is_read":    n.IsRead,
			"created_at": n.CreatedAt,
			"read_at":    n.ReadAt,
		})
	}
	return response.OK(c, http.StatusOK, "success", out)
}

// HandleMarkRead godoc
// @Summary Mark a notification as read
// @Tags notify
// @Param id path string true "Notification ID"
// @Success 200 {object} response.ResOk
// @Failure 400 {object} response.ResErr
// @Failure 500 {object} response.ResErr
// @Router /notify/mark-read/{id} [post]
func (h *notifyHandlerImpl) HandleMarkRead(c echo.Context) error {
	id := c.Param("id")

	if err := h.notif.MarkAsRead(c.Request().Context(), id); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.SimpleOK(c, http.StatusOK, nil)
}
