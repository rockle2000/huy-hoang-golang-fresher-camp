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

func UserLikeRestaurant(ctx component.AppContext) gin.HandlerFunc {
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
		incStore := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantlikebiz.NewUserLikeRestaurantStore(store, incStore)
		if err := biz.UserLikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleResponseSuccess(true))
	}
}
