package server

import (
	"time"

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
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

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

	userServiceURL := viper.GetString("USER_SERVICE_URL")
	if userServiceURL == "" {
		userServiceURL = "http://user-service:3000"
	}
	userClient := provider.NewHttpUserService(userServiceURL, 5*time.Second)
	var (
		appProvider         = provider.NewAppProvider(serviceConf)
		prescriptionRepo    = repository.NewPrescriptionRepo(appProvider.Postgres)
		prescriptionHandler = handler.NewPrescriptionHandler(handler.PrescriptionHandlerInject{
			FilterPrescriptionUseCase: usecase.NewFilterPrescriptionUseCase(prescriptionRepo),
			PrescriptionDetailUseCase: usecase.NewGetPrescriptionUseCase(prescriptionRepo),
			PrescriptionCreateUseCase: usecase.NewPrescriptionCreateUseCase(prescriptionRepo),
			DeleteMedicationUseCase:   usecase.NewDeleteMedicationUseCase(prescriptionRepo),
			CreateMedicationUseCase:   usecase.NewCreateMedicationUseCase(prescriptionRepo),
			UpdateMedicationUseCase:   usecase.NewUpdateMedicationUseCase(prescriptionRepo),
			ListMedicationUseCase:     usecase.NewListMedicationUseCase(prescriptionRepo),
			DetailMedicationUseCase:   usecase.NewDetailMedicationUseCase(prescriptionRepo),

			UserService: userClient,
		})
		routes = Routes(prescriptionHandler)
		server = NewServer(serviceConf, routes)
	)

	e := server.Engine
	g := e.Group("/" + serviceConf.ServiceName)
	g.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	if allowMigration {
		migration.Must(appProvider.Postgres.Executor)
	}

	server.Run()
}
