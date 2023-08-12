package utils

import "github.com/gin-gonic/gin"

const USER_ID = "USER_ID"

type session struct {
}

func (this *session) SetUserId(c *gin.Context, userId int) {
	c.Set(USER_ID, userId)
}

func (this *session) GetUserId(c *gin.Context) int {
	return c.GetInt(USER_ID)
}

var Session = session{}
