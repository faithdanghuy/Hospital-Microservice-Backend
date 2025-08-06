package provider

import (
	"github.com/Hospital-Microservice/appointment-service/model"
	"github.com/Hospital-Microservice/hospital-core/db"
)

type AppProvider struct {
	Postgres *db.Database
}

func NewAppProvider(config model.ServiceConfig) *AppProvider {
	return &AppProvider{
		Postgres: ProvidePostgres(config),
	}
}
