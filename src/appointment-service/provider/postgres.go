package provider

import (
	"fmt"

	"github.com/Hospital-Microservice/appointment-service/model"
	"github.com/Hospital-Microservice/hospital-core/db"
)

func ProvidePostgres(config model.ServiceConfig) *db.Database {
	fmt.Println("Connecting to PostgreSQL")

	return db.New(db.Connection{
		Host:                        config.DBHost,
		Port:                        config.DBPort,
		Database:                    config.DBName,
		User:                        config.DBUser,
		Password:                    config.DBPwd,
		SSLMode:                     db.VerifyCA,
		SSLCertAuthorityCertificate: "E:/Hospital-Microservice/src/appointment-service/deploy/ca.pem",
	})
}
