package engine

import (
	"net/http"
	"time"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoOption func(echo *echo.Echo)

func NewEcho(options ...EchoOption) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Secure())
	//e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "connection timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			log.Infof("timeout url: %s", c.Path())
		},
		Timeout: 30 * time.Second,
	}))
	e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 100, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}))

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	for _, option := range options {
		option(e)
	}
	return e
}

func AddMiddlewares(ms ...echo.MiddlewareFunc) EchoOption {
	return func(e *echo.Echo) {
		for _, m := range ms {
			e.Use(m)
		}
	}
}

func AddListGroupRoutes(lgr ...[]route.GroupRoute) EchoOption {
	return func(e *echo.Echo) {
		for _, gr := range lgr {
			for _, r := range gr {
				r.AddGroupRouteToEngine(e)
			}
		}
	}
}

func AddGroupRoutes(gr []route.GroupRoute) EchoOption {
	return func(e *echo.Echo) {
		for _, r := range gr {
			r.AddGroupRouteToEngine(e)
		}
	}
}

func AddRoutes(rs []route.Route) EchoOption {
	return func(e *echo.Echo) {
		for _, r := range rs {
			r.AddRouteToEngine(e)
		}
	}
}

func AddOptions(options ...EchoOption) EchoOption {
	return func(e *echo.Echo) {
		for _, o := range options {
			o(e)
		}
	}
}
