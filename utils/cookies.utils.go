package utils

import (
	"github.com/gin-gonic/gin"
)

const COOKIE_KEY = "SESSION"

type cookie struct {
}

func (this *cookie) SetSessionCookie(c *gin.Context, cookie string) {
	c.SetCookie(COOKIE_KEY, cookie, 3600, "/", "localhost", false, true)
}

func (this *cookie) GetSessionCookie(c *gin.Context) string {
	cookie, error := c.Cookie(COOKIE_KEY)

	if error != nil {
		return ""
	}

	return cookie
}

var Cookies = cookie{}
