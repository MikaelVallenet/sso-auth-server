package routes

import "gorm.io/gorm"

type Svc struct {
	Db *gorm.DB
}
