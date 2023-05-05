package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (svc *Svc) Token(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Success")
}
