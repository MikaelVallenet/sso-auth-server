package routes

import "github.com/go-playground/validator/v10"

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
