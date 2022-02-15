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

func CreateRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newRestaurant restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&newRestaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		newRestaurant.OwnerId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &newRestaurant); err != nil {
			panic(err)
		}
		newRestaurant.GenUID(common.DBTypeRestaurant)
		c.JSON(http.StatusOK, common.SimpleResponseSuccess(newRestaurant.FakeId))
	}
}
