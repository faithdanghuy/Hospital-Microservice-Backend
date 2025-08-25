package server

import (
	"github.com/Hospital-Microservice/hospital-core/config"
	. "github.com/Hospital-Microservice/hospital-core/transport/http"
	"github.com/Hospital-Microservice/hospital-core/transport/http/engine"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/Hospital-Microservice/user-service/handler"
	"github.com/Hospital-Microservice/user-service/migration"
	"github.com/Hospital-Microservice/user-service/model"
	"github.com/Hospital-Microservice/user-service/provider"
	"github.com/Hospital-Microservice/user-service/repository"
	"github.com/Hospital-Microservice/user-service/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

const allowMigration = false

func NewServer(serviceConf model.ServiceConfig, routes []route.GroupRoute) *Server {
	var e = engine.NewEcho()

	servicePrefix := "/" + serviceConf.ServiceName
	groupedRoutes := []route.GroupRoute{
		{
			Prefix: servicePrefix,
			Routes: nil,
		},
	}
	for _, gr := range routes {
		gr.Prefix = servicePrefix + gr.Prefix
		groupedRoutes = append(groupedRoutes, gr)
	}

	return NewHttpServer(
		AddName(serviceConf.ServiceName),
		AddPort(serviceConf.ServicePort),
		AddEngine(e),
		AddGroupRoutes(groupedRoutes[1:]),
	)
}

func Run(confPath string) {
	var serviceConf model.ServiceConfig
	config.MustLoadConfig(confPath, &serviceConf)

	var (
		appProvider = provider.NewAppProvider(serviceConf)
		userRepo    = repository.NewUserRepo(appProvider.Postgres)
		userHandler = handler.NewUserHandler(handler.UserHandlerInject{
			LoginUseCase:       usecase.NewLoginUseCase(appProvider, userRepo),
			RegisterUseCase:    usecase.NewRegisterUseCase(userRepo),
			ProfileUseCase:     usecase.NewProfileUseCase(userRepo),
			UpdateUseCase:      usecase.NewUpdateUseCase(userRepo),
			ChangePwdUseCase:   usecase.NewChangePasswordUseCase(userRepo),
			FilterUsersUseCase: usecase.NewFilterUsersUseCase(userRepo),
			UserDetailUseCase:  usecase.NewUserDetailUseCase(userRepo),
		})
		routes = Routes(userHandler)
		server = NewServer(serviceConf, routes)
	)

	e := server.Engine
	g := e.Group("/" + serviceConf.ServiceName)
	g.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PATCH, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	if allowMigration {
		migration.Must(appProvider.Postgres.Executor)
	}

	server.Run()
}
