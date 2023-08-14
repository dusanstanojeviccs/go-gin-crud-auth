package middleware

import (
	"go-gin-crud-auth/utils"
	"go-gin-crud-auth/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// extracts the UserId from the jwt and adds it as part of the context
func Auth(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth != "" {
		authFields := strings.Fields(auth)

		// Authorization header when present MUST be valid
		if len(authFields) != 2 || authFields[0] != "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		auth = authFields[1]
	} else {
		sessionCookie := utils.Cookies.GetSessionCookie(c)

		// cookie not being present or available is fine
		// it just means there is no currently logged in user
		// the request can procced but the per end point
		// security should block it
		if sessionCookie == "" {
			utils.Session.SetUserId(c, 0)
			utils.Cookies.SetSessionCookie(c, jwt.Jwt.GenerateSessionCookie(0))
			c.Next()
			return
		}

		auth = sessionCookie
	}

	// at this point it's guaranteed that the cookie or the authorization header was present
	// which means that it must be valid and contain a userId claim

	userId, valid := jwt.Jwt.ParseSessionCookie(auth)

	if valid {
		utils.Session.SetUserId(c, userId)
		c.Next()
	} else {
		utils.Session.SetUserId(c, 0)
		utils.Cookies.SetSessionCookie(c, jwt.Jwt.GenerateSessionCookie(0))
		// we clear the session and process the request
		c.Next()
	}
}
