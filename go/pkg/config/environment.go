package config

import (
	"log"

	"github.com/spf13/viper"
)

const environmentPath = "go/.env"

type Config struct {
	Port             string `mapstructure:"PORT"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
}

func loadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile(environmentPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return Config{}, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return Config{}, err
	}

	return config, nil
}
