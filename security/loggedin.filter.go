package security

import (
	"go-gin-crud-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggedInFilter(c *gin.Context) {
	if utils.Session.GetUserId(c) > 0 {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
