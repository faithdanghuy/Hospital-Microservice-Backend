package main

import (
	_ "github.com/Hospital-Microservice/appointment-service/docs"
	"github.com/Hospital-Microservice/appointment-service/server"
)

// @title Hospital Microservice API
// @version 1.0
// @description API documentation for Hospital Microservice
// @host localhost:3080
// @BasePath /appointment-service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer {your-token}"
func main() {
	server.Run("/Hospital-Microservice/src/appointment-service/deploy/local_conf.env")

}
