package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server ServerConfiguration
	Domain DomainConfiguration
}

type DomainConfiguration struct {
	Google   string
	GrpcUser string
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

// SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error al leer el archivo de configuración, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("No se puede decodificar en estructura, %v", err)
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
