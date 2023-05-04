package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Success")
}
