package model

type ServiceConfig struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	ServicePort int    `mapstructure:"SERVICE_PORT"`

	DBHost string `mapstructure:"DB_HOST"`
	DBPort int    `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPwd  string `mapstructure:"DB_PWD"`
	DBName string `mapstructure:"DB_DBNAME"`

	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}
