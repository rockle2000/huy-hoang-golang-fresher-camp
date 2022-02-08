package ginfood

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/common"
	"test/component"
	"test/modules/food/foodbiz"
	"test/modules/food/foodmodel"
	"test/modules/food/foodstorage"
)

func UpdateFood(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data foodmodel.FoodUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := foodbiz.NewUpdateFoodBiz(store)
		if err := biz.UpdateFood(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleResponseSuccess("Update successfully"))
	}
}
