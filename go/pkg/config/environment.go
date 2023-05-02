package config

import (
	"log"

	"github.com/spf13/viper"
)

const environmentPath = "go/.env"

var Config *config

type config struct {
	Port             string `mapstructure:"PORT"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
}

func LoadConfig() error {

	viper.SetConfigFile(environmentPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return err
	}

	return nil
}
