package main

import (
	"fmt"

	"github.com/Mikatech/sso-auth-server/go/pkg/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		return
	}
	fmt.Println(config.Config.PostgresDb)
}
