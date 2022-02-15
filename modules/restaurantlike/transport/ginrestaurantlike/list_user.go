package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
	restaurantlikebiz "test/modules/restaurantlike/biz"
	restaurantlikemodel "test/modules/restaurantlike/model"
	restaurantlikestorage "test/modules/restaurantlike/storage"
)

func ListUser(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		//var filter restaurantlikemodel.Filter
		//
		//if err := c.ShouldBind(&filter); err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantlikebiz.NewListUserLikeRestaurantBiz(store)
		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewResponseSuccess(result, paging, filter))

	}
}
