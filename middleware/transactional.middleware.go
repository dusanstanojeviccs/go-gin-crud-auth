package middleware

import (
	"go-gin-crud-auth/utils"
	"go-gin-crud-auth/utils/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Transactional(c *gin.Context) {
	txHandle, err := db.DB.Begin()

	if err != nil {
		utils.Error.Report(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			txHandle.Rollback()
		}
	}()

	db.SetTransaction(c, txHandle)

	c.Next()

	if c.Writer.Status() == http.StatusOK || c.Writer.Status() == http.StatusCreated {
		if err := txHandle.Commit(); err != nil {
			utils.Error.Report(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	} else {
		txHandle.Rollback()
	}
}
