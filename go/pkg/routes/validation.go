package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "This field must be a valid email address"
	case "min":
		return "This field must be at least " + fe.Param() + " characters long"
	case "max":
		return "This field must be at most " + fe.Param() + " characters long"
	}
	return "Unknown error"
}

func bindingValidations(ctx *gin.Context, in interface{}) bool {
	if err := ctx.ShouldBindJSON(&in); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return false
	}
	return true
}
