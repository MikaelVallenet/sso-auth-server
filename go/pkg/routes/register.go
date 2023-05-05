package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	in := RegisterInput{}
	if !bindingValidations(ctx, &in) {
		return
	}

	fmt.Println(in.Username)
	ctx.JSON(http.StatusAccepted, &in)
}
