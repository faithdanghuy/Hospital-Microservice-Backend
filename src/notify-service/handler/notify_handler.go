package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/notify-service/model/req"
	"github.com/Hospital-Microservice/notify-service/usecase"
	"github.com/labstack/echo/v4"
)

type NotifyHandler interface {
	HandleSend(c echo.Context) error
	HandleHealth(c echo.Context) error
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
