package provider

import (
	"github.com/Hospital-Microservice/hospital-core/db"
	"github.com/Hospital-Microservice/user-service/model"
)

type AppProvider struct {
	Postgres *db.Database
}

func NewAppProvider(config model.ServiceConfig) *AppProvider {
	return &AppProvider{
		Postgres: ProvidePostgres(config),
	}
}
