package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
	"test/modules/restaurant/restaurantbiz"
	"test/modules/restaurant/restaurantstorage"
)

func GetRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)
		data, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleResponseSuccess(data))
	}
}
