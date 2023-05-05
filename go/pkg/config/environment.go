package config

import (
	"log"

	"github.com/spf13/viper"
)

const environmentPath = "go/.env"

var Config *config

type config struct {
	Port             string `mapstructure:"PORT"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	Issuer           string `mapstructure:"ISSUER"`
	ExpireIn         int    `mapstructure:"EXPIRE_IN"`
	PrivateKeyPath   string `mapstructure:"PRIVATE_KEY_PATH"`
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
