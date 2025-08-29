package server

import (
	"github.com/Hospital-Microservice/hospital-core/middleware"
	"github.com/Hospital-Microservice/hospital-core/transport/http/method"
	coreRoute "github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/Hospital-Microservice/notify-service/handler"
	"github.com/labstack/echo/v4"
)

// Routes returns group routes for notify service.
func Routes(h handler.NotifyHandler) []coreRoute.GroupRoute {
	return []coreRoute.GroupRoute{
		{
			Prefix: "/notify",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWT(), // optional, remove if you want public endpoint
			},
			Routes: []coreRoute.Route{
				{
					Path:    "/send",
					Method:  method.POST,
					Handler: h.HandleSend,
				},
				{
					Path:    "/health",
					Method:  method.GET,
					Handler: h.HandleHealth,
				},
			},
		},
	}
}
