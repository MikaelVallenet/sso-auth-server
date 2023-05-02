package main

import (
	"fmt"

	"github.com/Mikatech/sso-auth-server/go/pkg/config"
	"github.com/Mikatech/sso-auth-server/go/pkg/database"
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
	fmt.Println(config.Config.PostgresDb)
}
