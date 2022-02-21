package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
	"test/modules/restaurant/restaurantstorage"
	restaurantlikebiz "test/modules/restaurantlike/biz"
	restaurantlikemodel "test/modules/restaurantlike/model"
	restaurantlikestorage "test/modules/restaurantlike/storage"
)

func UserUnlikeRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		decStore := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantlikebiz.NewUserUnlikeRestaurantBiz(store, decStore)
		if err := biz.UserUnlikeRestaurant(c.Request.Context(), data.UserId, data.RestaurantId); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleResponseSuccess(true))
	}
}
