package provider

import (
	"fmt"

	"github.com/Hospital-Microservice/hospital-core/db"
	"github.com/Hospital-Microservice/user-service/model"
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
		SSLCertAuthorityCertificate: config.DBSSLRootCert,
	})
}
