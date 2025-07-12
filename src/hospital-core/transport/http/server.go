package http

import (
	"fmt"
	"time"

	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/labstack/echo/v4"
)

type ServerOption func(*Server)

type Server struct {
	Name                    string
	Port                    int
	Routes                  []route.Route
	GroupRoutes             []route.GroupRoute
	EchoOptions             []ServerOption
	GracefulShutdownTimeout time.Duration
	Engine                  *echo.Echo
}

func NewHttpServer(options ...ServerOption) *Server {
	s := &Server{}

	// Thêm các options cho Server
	for _, option := range options {
		option(s)
	}

	for _, r := range s.Routes {
		r.AddRouteToEngine(s.Engine)
	}

	for _, gr := range s.GroupRoutes {
		gr.AddGroupRouteToEngine(s.Engine)
	}
	return s
}

func (s *Server) Run() {
	fmt.Printf("%v is running on port %d", s.Name, s.Port)
	s.Engine.Logger.Fatal(
		s.Engine.Start(fmt.Sprintf(":%d", s.Port)),
	)
}

func AddName(n string) ServerOption {
	return func(s *Server) {
		s.Name = n
	}
}

func AddPort(p int) ServerOption {
	return func(s *Server) {
		s.Port = p
	}
}

func AddEngine(e *echo.Echo) ServerOption {
	return func(s *Server) {
		s.Engine = e
	}
}

func AddGracefulShutdownTimeout(timeout int) ServerOption {
	return func(s *Server) {
		s.GracefulShutdownTimeout = time.Duration(timeout) * time.Second
	}
}

func AddGroupRoutes(gr []route.GroupRoute) ServerOption {
	return func(s *Server) {
		s.GroupRoutes = append(s.GroupRoutes, gr...)
	}
}

func AddRoutes(rs []route.Route) ServerOption {
	return func(s *Server) {
		s.Routes = append(s.Routes, rs...)
	}
}

func AddRoute(rs route.Route) ServerOption {
	return func(s *Server) {
		s.Routes = append(s.Routes, rs)
	}
}

func AddServerOptions(options ...ServerOption) ServerOption {
	return func(s *Server) {
		for _, o := range options {
			o(s)
		}
	}
}
