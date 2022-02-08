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

func GetFood(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := foodbiz.NewGetFoodBiz(store)
		result, err := biz.GetFood(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, result)
	}
}
