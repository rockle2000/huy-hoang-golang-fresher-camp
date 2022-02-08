package middleware

import (
	"github.com/gin-gonic/gin"
	"test/common"
	"test/component"
)

func Recover(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(appErr)
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(appErr)
				return
			}
		}()
		c.Next()
	}
}
