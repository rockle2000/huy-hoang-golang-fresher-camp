package ginfood

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
	"test/modules/food/foodbiz"
	"test/modules/food/foodmodel"
	"test/modules/food/foodstorage"
)

func CreateFood(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newFood foodmodel.FoodCreate

		if err := c.ShouldBind(&newFood); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := foodstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := foodbiz.NewCreateFoodBiz(store)
		if err := biz.CreateFood(c.Request.Context(), &newFood); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleResponseSuccess(newFood))
	}
}
