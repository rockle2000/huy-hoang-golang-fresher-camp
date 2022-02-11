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

func DeleteRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(common.ErrEntityDeleted(restaurantmodel.EntityName, err))
		}
		c.JSON(http.StatusOK, common.SimpleResponseSuccess(true))

	}
}
