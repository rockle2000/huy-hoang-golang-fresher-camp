package ginsanpham

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
	"test/modules/sanpham/sanphambiz"
	"test/modules/sanpham/sanphammodel"
	"test/modules/sanpham/sanphamstorage"
)

func CreateFood(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newFood sanphammodel.FoodCreate

		if err := c.ShouldBind(&newFood); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := sanphamstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := sanphambiz.NewCreateFoodBiz(store)
		if err := biz.CreateFood(c.Request.Context(), &newFood); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleResponseSuccess(newFood))
	}
}
