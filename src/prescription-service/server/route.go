package server

import (
	"github.com/Hospital-Microservice/hospital-core/transport/http/method"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/Hospital-Microservice/prescription-service/handler"
)

func Routes(handler handler.PrescriptionHandler) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/prescription",
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: handler.HandlePrescriptionCreate,
				},
				// {
				// 	Path:    "/update",
				// 	Method:  method.PATCH,
				// 	Handler: handler.HandleRegister,
				// },
				{
					Path:    "/detail/:id",
					Method:  method.GET,
					Handler: handler.HandlePrescriptionDetail,
				},
				// {
				// 	Path:    "/filter",
				// 	Method:  method.GET,
				// 	Handler: handler.HandleRegister,
				// },
				// {
				// 	Path:    "/delete",
				// 	Method:  method.DELETE,
				// 	Handler: handler.HandleProfile,
				// },
			},
		},
	}
}
