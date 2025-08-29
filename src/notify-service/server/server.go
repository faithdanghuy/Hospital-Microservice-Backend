package server

import (
	"context"
	"fmt"
	"time"

	"github.com/Hospital-Microservice/hospital-core/config"
	. "github.com/Hospital-Microservice/hospital-core/transport/http"
	"github.com/Hospital-Microservice/hospital-core/transport/http/engine"
	coreRoute "github.com/Hospital-Microservice/hospital-core/transport/http/route"
	"github.com/Hospital-Microservice/notify-service/handler"
	"github.com/Hospital-Microservice/notify-service/model"
	"github.com/Hospital-Microservice/notify-service/provider"
	"github.com/Hospital-Microservice/notify-service/repository"
	"github.com/Hospital-Microservice/notify-service/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewServer(serviceConf model.ServiceConfig, routes []coreRoute.GroupRoute) *Server {
	var e = engine.NewEcho()

	servicePrefix := "/" + serviceConf.ServiceName
	groupedRoutes := []coreRoute.GroupRoute{
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

	mongoURI := viper.GetString("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb+srv://hihuunguyen_db_user:neejQhziYbPaERmZ@notifyservice.pdptqkh.mongodb.net/?retryWrites=true&w=majority&appName=NotifyService"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	// ping
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	db := client.Database("NotifyService")
	coll := db.Collection("notifications")

	// providers
	rabbitURL := viper.GetString("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@rabbitmq:5672/"
	}
	smtpHost := viper.GetString("SMTP_HOST")
	smtpPort := viper.GetInt("SMTP_PORT")
	smtpUser := viper.GetString("SMTP_USER")
	smtpPass := viper.GetString("SMTP_PASS")
	smsProviderURL := viper.GetString("SMS_PROVIDER_URL")

	rabbitSub := provider.NewRabbitSubscriber(rabbitURL, 5*time.Second)
	emailClient := provider.NewSMTPEmailClient(smtpHost, smtpPort, smtpUser, smtpPass, 10*time.Second)
	smsClient := provider.NewHTTPSmsClient(smsProviderURL, 10*time.Second)

	notifRepo := repository.NewMongoNotificationRepo(coll)
	notifyUC := usecase.NewNotifyUseCase(emailClient, smsClient, notifRepo)
	notifyHandler := handler.NewNotifyHandler(handler.Inject{
		NotifyUC: notifyUC,
	})

	routes := Routes(notifyHandler)
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

	go func() {
		if err := rabbitSub.StartConsume(func(queue string, body []byte) error {
			return notifyUC.HandleEvent(queue, body)
		}); err != nil {
			fmt.Printf("Rabbit consumer error: %v\n", err)
		}
	}()

	server.Run()
}
