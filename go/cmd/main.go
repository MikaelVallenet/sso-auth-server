package main

import (
	"github.com/Mikatech/sso-auth-server/go/pkg/config"
	"github.com/Mikatech/sso-auth-server/go/pkg/database"
	"github.com/Mikatech/sso-auth-server/go/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		return
	}
	_, err = database.Init()
	if err != nil {
		return
	}
	router := gin.Default()
	router.POST("/register", routes.Register)
	router.Run(":" + config.Config.Port)
}
