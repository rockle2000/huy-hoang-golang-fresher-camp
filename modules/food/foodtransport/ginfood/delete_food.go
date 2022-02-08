package ginfood

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/common"
	"test/component"
	"test/modules/food/foodbiz"
	"test/modules/food/foodstorage"
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

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewDeleteFoodBiz(store)

		if err := biz.DeleteFood(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleResponseSuccess("delete successfully"))

	}
}
