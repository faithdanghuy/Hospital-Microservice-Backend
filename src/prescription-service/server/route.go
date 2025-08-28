package server

import (
	"github.com/Hospital-Microservice/hospital-core/middleware"
	"github.com/Hospital-Microservice/hospital-core/transport/http/method"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/Hospital-Microservice/prescription-service/handler"
	"github.com/labstack/echo/v4"
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
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/update",
					Method:  method.PATCH,
					Handler: handler.HandlePrescriptionUpdate,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/detail/:id",
					Method:  method.GET,
					Handler: handler.HandlePrescriptionDetail,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/filter",
					Method:  method.GET,
					Handler: handler.HandlePrescriptionFilter,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/delete",
					Method:  method.DELETE,
					Handler: handler.HandlePrescriptionDelete,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
			},
		},
		{
			Prefix: "/medication",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWT(),
			},
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: handler.HandleCreateMedication,
				},
				{
					Path:    "/update",
					Method:  method.PATCH,
					Handler: handler.HandleUpdateMedication,
				},
				{
					Path:    "/filter",
					Method:  method.GET,
					Handler: handler.HandleListMedications,
				},
				{
					Path:    "/detail/:id",
					Method:  method.GET,
					Handler: handler.HandleDetailMedication,
				},
			},
		},
	}
}
