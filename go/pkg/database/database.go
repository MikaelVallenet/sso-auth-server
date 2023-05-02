package database

import (
	"fmt"
	"log"

	"github.com/Mikatech/sso-auth-server/go/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func dsnBuilder() string {
	dsn := "host=" + config.Config.PostgresHost + " user=" + config.Config.PostgresUser + " password=" + config.Config.PostgresPassword + " dbname=" + config.Config.PostgresDb + " port=" + config.Config.PostgresPort + " sslmode=disable TimeZone=Europe/Paris"
	return dsn
}

func Init() (*Database, error) {
	dsn := dsnBuilder()
	fmt.Println(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database, %s", err)
		return nil, err
	}
	return &Database{Client: database}, nil
}
