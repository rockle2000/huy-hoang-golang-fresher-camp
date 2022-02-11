package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
	"test/modules/restaurant/restaurantbiz"
	"test/modules/restaurant/restaurantmodel"
	"test/modules/restaurant/restaurantstorage"
)

func UpdateRestaurant(ctx component.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		if err := biz.UpdateRestaurant(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleResponseSuccess(true))
	}
}
