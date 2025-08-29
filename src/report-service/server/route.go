package server

import (
	"github.com/Hospital-Microservice/hospital-core/middleware"
	"github.com/Hospital-Microservice/hospital-core/transport/http/method"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/Hospital-Microservice/report-service/handler"
	"github.com/labstack/echo/v4"
)

func Routes(h handler.ReportHandler) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/report",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWT(),
			},
			Routes: []route.Route{
				{
					Path:    "/patients",
					Method:  method.GET,
					Handler: h.HandlePatients,
				},
				{
					Path:    "/appointments",
					Method:  method.GET,
					Handler: h.HandleAppointments,
				},
				{
					Path:    "/prescriptions",
					Method:  method.GET,
					Handler: h.HandlePrescriptions,
				},
			},
		},
	}
}
