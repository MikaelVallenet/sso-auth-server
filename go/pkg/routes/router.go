package routes

import (
	"github.com/gin-gonic/gin"
)

func (svc *Svc) Init() *gin.Engine {
	router := gin.New()
	router.POST("/register", svc.Register)
	router.POST("/login", svc.Login)
	router.POST("/token", svc.Token)
	return router
}
