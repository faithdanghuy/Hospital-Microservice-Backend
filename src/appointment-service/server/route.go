package server

import (
	"github.com/Hospital-Microservice/appointment-service/handler"
	"github.com/Hospital-Microservice/hospital-core/transport/http/method"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
)

func Routes(handler handler.AppointmentHandler) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/appointment",
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: handler.HandleAppointmentCreate,
				},
				// {
				// 	Path:    "/update",
				// 	Method:  method.PATCH,
				// 	Handler: handler.HandleRegister,
				// },
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
				// {
				// 	Path:    "/delete",
				// 	Method:  method.DELETE,
				// 	Handler: handler.HandleProfile,
				// },
			},
		},
	}
}
