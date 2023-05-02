package database

import "gorm.io/gorm"

type Database struct {
	Client *gorm.DB
}

func Setup() (db *Database, err error) {
	database, err := gorm.Open(postgres)
}
