package model

type ServiceConfig struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	ServicePort int    `mapstructure:"SERVICE_PORT"`
}
