package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	c.String(http.StatusOK, "Success")
}
