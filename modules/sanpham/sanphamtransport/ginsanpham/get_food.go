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

func GetFood(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := sanphamstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := sanphambiz.NewGetFoodBiz(store)
		result, err := biz.GetFood(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, result)
	}
}
