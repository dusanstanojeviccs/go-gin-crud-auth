package middleware

import (
	"go-gin-crud-auth/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AppClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

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
		sessionCookie, error := c.Cookie("session")

		// cookie not being present or available is fine
		// it just means there is no currently logged in user
		// the request can procced but the per end point
		// security should block it
		if error == nil || sessionCookie == "" {
			c.Next()
			return
		}

		auth = sessionCookie
	}

	// at this point it's guaranteed that the cookie or the authorization header was present
	// which means that it must be valid and contain a userId claim

	token, error := jwt.ParseWithClaims(auth, &AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("THIS_WILL_BE_OUR_ENV_VAR_SOON"), nil
	})

	if error != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(*AppClaims); ok {
		utils.Session.SetUserId(c, claims.UserId)
		c.Next()
	}
}
