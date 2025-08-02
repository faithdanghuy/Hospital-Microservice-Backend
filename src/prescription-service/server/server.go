package server

import (
	"github.com/Hospital-Microservice/hospital-core/config"
	. "github.com/Hospital-Microservice/hospital-core/transport/http"
	"github.com/Hospital-Microservice/hospital-core/transport/http/engine"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/Hospital-Microservice/prescription-service/handler"
	"github.com/Hospital-Microservice/prescription-service/migration"
	"github.com/Hospital-Microservice/prescription-service/model"
	"github.com/Hospital-Microservice/prescription-service/provider"
	"github.com/Hospital-Microservice/prescription-service/repository"
	"github.com/Hospital-Microservice/prescription-service/usecase"
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
		appProvider         = provider.NewAppProvider(serviceConf)
		prescriptionRepo    = repository.NewPrescriptionRepo(appProvider.Postgres)
		prescriptionHandler = handler.NewPrescriptionHandler(handler.PrescriptionHandlerInject{
			PrescriptionDetailUseCase: usecase.NewPrescriptionDetailUseCase(prescriptionRepo),
		})
		routes = Routes(prescriptionHandler)
		server = NewServer(serviceConf, routes)
	)

	if allowMigration {
		migration.Must(appProvider.Postgres.Executor)
	}

	server.Run()
}
