package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// TODO: Interface the binding and validation
func Register(ctx *gin.Context) {
	in := RegisterInput{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	ctx.JSON(http.StatusAccepted, &in)
}
