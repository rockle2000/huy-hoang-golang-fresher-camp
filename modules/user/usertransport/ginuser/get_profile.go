package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
)

func GetProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleResponseSuccess(data))
	}
}
