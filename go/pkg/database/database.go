package database

import (
	"log"
	"time"

	"github.com/Mikatech/sso-auth-server/go/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DBMaxOpenTries = 5

type Database struct {
	Client *gorm.DB
}

func dsnBuilder() string {
	dsn := "host=" + config.Config.PostgresHost + " user=" + config.Config.PostgresUser + " password=" + config.Config.PostgresPassword + " dbname=" + config.Config.PostgresDb + " port=" + config.Config.PostgresPort + " sslmode=disable TimeZone=Europe/Paris"
	return dsn
}

func Init() (*Database, error) {
	dsn := dsnBuilder()

	dbOpenTries := 0
	for dbOpenTries <= DBMaxOpenTries {
		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Error connecting to database, %s", err)
			dbOpenTries++
			time.Sleep(5 * time.Second)
			continue
		}
		err = migrate(database)
		if err != nil {
			log.Printf("Error migrating database, %s", err)
			dbOpenTries++
			return nil, err
		}
		return &Database{Client: database}, nil
	}
	return nil, nil
}
