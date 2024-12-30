package web

import (
	"github.com/gin-gonic/gin"
)

type WebHandlerFunc func(c *gin.Context) (int, any, error)

func ToHandlerFunc(f WebHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		status, data, err := f(c)
		if err != nil {
			HandleError(c, err)
			return
		}
		c.JSON(status, data)
	}
}

func HandleError(c *gin.Context, err error) {
	status, res := ErrorResFrom(c, err)
	c.JSON(status, res)
}
