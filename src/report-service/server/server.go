package server

import (
	"time"

	"github.com/Hospital-Microservice/hospital-core/config"
	. "github.com/Hospital-Microservice/hospital-core/transport/http"
	"github.com/Hospital-Microservice/hospital-core/transport/http/engine"
	"github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/Hospital-Microservice/report-service/handler"
	"github.com/Hospital-Microservice/report-service/model"
	"github.com/Hospital-Microservice/report-service/provider"
	"github.com/Hospital-Microservice/report-service/usecase"
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
	apptServiceURL := viper.GetString("APPOINTMENT_SERVICE_URL")
	if apptServiceURL == "" {
		apptServiceURL = "http://appointment-service:3001"
	}
	presServiceURL := viper.GetString("PRESCRIPTION_SERVICE_URL")
	if presServiceURL == "" {
		presServiceURL = "http://prescription-service:3002"
	}

	userClient := provider.NewHttpUserClient(userServiceURL, 10*time.Second)
	apptClient := provider.NewHttpAppointmentClient(apptServiceURL, 10*time.Second)
	presClient := provider.NewHttpPrescriptionClient(presServiceURL, 10*time.Second)

	patientsUC := usecase.NewPatientsReportUseCase(userClient)
	apptUC := usecase.NewAppointmentsReportUseCase(apptClient)
	presUC := usecase.NewPrescriptionsReportUseCase(presClient)

	reportHandler := handler.NewReportHandler(handler.ReportHandlerInject{
		PatientsUC:      patientsUC,
		AppointmentsUC:  apptUC,
		PrescriptionsUC: presUC,
	})

	routes := Routes(*reportHandler)
	server := NewServer(serviceConf, routes)

	e := server.Engine
	g := e.Group("/" + serviceConf.ServiceName)
	g.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PATCH, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	server.Run()
}
