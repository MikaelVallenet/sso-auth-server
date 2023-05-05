package main

import (
	"github.com/Mikatech/sso-auth-server/go/pkg/config"
	"github.com/Mikatech/sso-auth-server/go/pkg/database"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		return
	}
	svc, err := database.Init()
	if err != nil {
		return
	}
	router := svc.Init()
	router.Run(":" + config.Config.Port)
}
