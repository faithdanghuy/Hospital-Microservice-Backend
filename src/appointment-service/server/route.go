package server

import (
	"github.com/Hospital-Microservice/appointment-service/handler"
	"github.com/Hospital-Microservice/hospital-core/middleware"
	"github.com/Hospital-Microservice/hospital-core/transport/http/method"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/labstack/echo/v4"
)

func Routes(handler handler.AppointmentHandler) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/appointment",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWT(),
			},
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: handler.HandleAppointmentCreate,
				},
				{
					Path:    "/detail/:id",
					Method:  method.GET,
					Handler: handler.HandleAppointmentDetail,
				},
				{
					Path:    "/change-status/:id",
					Method:  method.PATCH,
					Handler: handler.HandleAppointmentChangeStatus,
				},
				{
					Path:    "/filter",
					Method:  method.GET,
					Handler: handler.HandleAppointmentFilter,
				},
				{
					Path:    "/edit/:id",
					Method:  method.PATCH,
					Handler: handler.HandleAppointmentEdit,
				},
				// {
				// 	Path:    "/delete",
				// 	Method:  method.DELETE,
				// 	Handler: handler.HandleProfile,
				// },
			},
		},
	}
}
