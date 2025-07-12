package config

import (
	"github.com/Hospital-Microservice/hospital-core/io"
	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func MustLoadConfig(configPath string, config interface{}) {
	dirPath := io.GetDirectoryPath(configPath)
	fileName, err := io.GetFileName(configPath)
	if err != nil {
		log.Error("failed to get file name", zap.String("error", err.Error()))
		panic(err)
	}
	viper.AddConfigPath(dirPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Error("failed to read config", zap.String("error", err.Error()))
		panic(err)
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Error("failed to decode config", zap.String("error", err.Error()))
		panic(err)
	}
}
