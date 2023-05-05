package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {
	svc := &Svc{db: db}
	router := gin.New()
	router.POST("/register", svc.Register)
	router.POST("/login", svc.Login)
	router.POST("/token", svc.Token)
	return router
}
