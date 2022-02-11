package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
	"test/modules/restaurant/restaurantbiz"
	"test/modules/restaurant/restaurantmodel"
	"test/modules/restaurant/restaurantstorage"
	restaurantlikestorage "test/modules/restaurantlike/storage"
)

func ListRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		likeStore := restaurantlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store, likeStore)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}
		c.JSON(http.StatusOK, common.NewResponseSuccess(result, paging, filter))

	}
}
