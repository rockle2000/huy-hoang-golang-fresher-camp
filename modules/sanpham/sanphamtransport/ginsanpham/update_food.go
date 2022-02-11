package ginsanpham

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/common"
	"test/component"
	"test/modules/sanpham/sanphambiz"
	"test/modules/sanpham/sanphammodel"
	"test/modules/sanpham/sanphamstorage"
)

func UpdateFood(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data sanphammodel.FoodUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := sanphamstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := sanphambiz.NewUpdateFoodBiz(store)
		if err := biz.UpdateFood(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleResponseSuccess("Update successfully"))
	}
}
