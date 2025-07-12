package route

import (
	"github.com/Hospital-Microservice/hospital-core/transport/http/method"
	"github.com/labstack/echo/v4"
)

type GroupRoute struct {
	Prefix      string
	Middlewares []echo.MiddlewareFunc
	Routes      []Route
}

type Route struct {
	Path        string
	Method      method.Method
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

func (r Route) AddRouteToEngine(e *echo.Echo) {
	switch r.Method {
	case method.GET:
		e.GET(r.Path, r.Handler, r.Middlewares...)
	case method.POST:
		e.POST(r.Path, r.Handler, r.Middlewares...)
	case method.PUT:
		e.PUT(r.Path, r.Handler, r.Middlewares...)
	case method.PATCH:
		e.PATCH(r.Path, r.Handler, r.Middlewares...)
	case method.DELETE:
		e.DELETE(r.Path, r.Handler, r.Middlewares...)
	case method.HEAD:
		e.HEAD(r.Path, r.Handler, r.Middlewares...)
	case method.OPTIONS:
		e.OPTIONS(r.Path, r.Handler, r.Middlewares...)
	}
}

func (r GroupRoute) AddGroupRouteToEngine(e *echo.Echo) {
	gr := e.Group(r.Prefix)
	gr.Use(r.Middlewares...)
	for _, r := range r.Routes {
		switch r.Method {
		case method.GET:
			gr.GET(r.Path, r.Handler, r.Middlewares...)
		case method.POST:
			gr.POST(r.Path, r.Handler, r.Middlewares...)
		case method.PUT:
			gr.PUT(r.Path, r.Handler, r.Middlewares...)
		case method.PATCH:
			gr.PATCH(r.Path, r.Handler, r.Middlewares...)
		case method.DELETE:
			gr.DELETE(r.Path, r.Handler, r.Middlewares...)
		case method.HEAD:
			gr.HEAD(r.Path, r.Handler, r.Middlewares...)
		case method.OPTIONS:
			gr.OPTIONS(r.Path, r.Handler, r.Middlewares...)
		}
	}
}
