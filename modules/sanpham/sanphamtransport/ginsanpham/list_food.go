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

func ListFood(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter sanphammodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Fulfill()

		store := sanphamstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := sanphambiz.NewListFoodBiz(store)
		result, err := biz.ListFood(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewResponseSuccess(result, paging, filter))
	}
}
