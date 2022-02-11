package ginsanpham

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/common"
	"test/component"
	"test/modules/sanpham/sanphambiz"
	"test/modules/sanpham/sanphamstorage"
)

func DeleteFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := sanphamstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := sanphambiz.NewDeleteFoodBiz(store)

		if err := biz.DeleteFood(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleResponseSuccess("delete successfully"))

	}
}
