package main

import (
	"github.com/Mikatech/sso-auth-server/go/pkg/config"
	"github.com/Mikatech/sso-auth-server/go/pkg/database"
	"github.com/Mikatech/sso-auth-server/go/pkg/routes"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		return
	}
	db, err := database.Init()
	if err != nil {
		return
	}
	router := routes.Init(db)
	router.Run(":" + config.Config.Port)
}
