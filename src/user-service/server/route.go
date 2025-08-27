package server

import (
	"github.com/Hospital-Microservice/hospital-core/middleware"
	"github.com/Hospital-Microservice/hospital-core/transport/http/method"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/Hospital-Microservice/user-service/handler"
	"github.com/labstack/echo/v4"
)

func Routes(handler handler.UserHandler) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/auth",
			Routes: []route.Route{
				{
					Path:    "/login",
					Method:  method.POST,
					Handler: handler.HandleLogin,
				},
				{
					Path:    "/register",
					Method:  method.POST,
					Handler: handler.HandleRegister,
				},
				{
					Path:    "/change-password",
					Method:  method.PATCH,
					Handler: handler.HandleChangePassword,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
			},
		},
		{
			Prefix: "/account",
			Routes: []route.Route{
				{
					Path:    "/profile",
					Method:  method.GET,
					Handler: handler.HandleProfile,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/update",
					Method:  method.PATCH,
					Handler: handler.HandleUpdate,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/filter",
					Method:  method.GET,
					Handler: handler.HandleFilterUsers,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/detail/:id",
					Method:  method.GET,
					Handler: handler.HandleUserDetail,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/edit/:id",
					Method:  method.PATCH,
					Handler: handler.HandleEditUser,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/delete/:id",
					Method:  method.DELETE,
					Handler: handler.HandleDeleteUser,
					Middlewares: []echo.MiddlewareFunc{
						middleware.JWT(),
					},
				},
				{
					Path:    "/doctors",
					Method:  method.GET,
					Handler: handler.HandleGetDoctors,
				},
				{
					Path:    "/patients",
					Method:  method.GET,
					Handler: handler.HandleGetPatients,
				},
				{
					Path:    "/batch",
					Method:  method.POST,
					Handler: handler.HandleAccountBatch,
				},
			},
		},
	}
}
