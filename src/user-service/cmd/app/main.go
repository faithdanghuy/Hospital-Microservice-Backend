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
func main() {

	server.Run("/Hospital-Microservice/src/user-service/deploy/local_conf.env")
	// e := echo.New()
	// g := e.Group("/user-service")

	// g.GET("/swagger/*", echoSwagger.WrapHandler)
	// port := os.Getenv("SERVICE_PORT")
	// if port == "" {
	// 	port = "3080"
	// }
	// e.Logger.Fatal(e.Start(":" + port))

}
