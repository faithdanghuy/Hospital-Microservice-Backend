package server

import (
	"github.com/Hospital-Microservice/appointment-service/handler"
	"github.com/Hospital-Microservice/appointment-service/migration"
	"github.com/Hospital-Microservice/appointment-service/model"
	"github.com/Hospital-Microservice/appointment-service/provider"
	"github.com/Hospital-Microservice/appointment-service/repository"
	"github.com/Hospital-Microservice/appointment-service/usecase"
	"github.com/Hospital-Microservice/hospital-core/config"
	. "github.com/Hospital-Microservice/hospital-core/transport/http"
	"github.com/Hospital-Microservice/hospital-core/transport/http/engine"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
)

const allowMigration = true

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
		appProvider        = provider.NewAppProvider(serviceConf)
		appointmentRepo    = repository.NewAppointmentRepo(appProvider.Postgres)
		appointmentHandler = handler.NewAppointmentHandler(handler.AppointmentHandlerInject{
			AppointmentDetailUseCase: usecase.NewAppointmentDetailUseCase(appointmentRepo),
			AppointmentCreateUseCase: usecase.NewAppointmentCreateUseCase(appointmentRepo),
		})
		routes = Routes(appointmentHandler)
		server = NewServer(serviceConf, routes)
	)

	if allowMigration {
		migration.Must(appProvider.Postgres.Executor)
	}

	server.Run()
}
