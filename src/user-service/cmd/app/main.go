package main

import (
	_ "github.com/Hospital-Microservice/user-service/docs"
	"github.com/Hospital-Microservice/user-service/server"
)

// @title Hospital Microservice API
// @version 1.0
// @description API documentation for Hospital Microservice
// @host localhost:3080
// @BasePath /user-service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer {your-token}"
func main() {

	server.Run("/Hospital-Microservice/src/user-service/deploy/local_conf.env")

}
