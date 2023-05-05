package routes

import (
	"net/http"

	"github.com/Mikatech/sso-auth-server/go/pkg/database"
	"github.com/gin-gonic/gin"
)

// TODO: Create a pkg for custom error code
func (svc *Svc) Register(ctx *gin.Context) {
	in := RegisterInput{}
	if !bindingValidations(ctx, &in) {
		return
	}

	// Verify that a similar user does not exist
	var userCount int64
	err := svc.db.Model(&database.User{}).Where("username = ? OR email = ?", in.Username, in.Email).Count(&userCount).Error
	if err != nil || userCount != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username or email is already taken"})
		return
	}

	ctx.JSON(http.StatusAccepted, &in)
}
