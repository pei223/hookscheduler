package web

import (
	"github.com/gin-gonic/gin"
)

type WebHandlerFunc func(c *gin.Context) (int, any, error)

func ToHandlerFunc(f WebHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		status, data, err := f(c)
		if err != nil {
			// TODO
			// exception handler的なやつをやりたい
			// エラーレスポンス型ならそのままそれを返す
			c.JSON(status, err.Error())
			return
		}
		c.JSON(status, data)
	}
}
