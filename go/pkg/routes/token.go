package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Token(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Success")
}
