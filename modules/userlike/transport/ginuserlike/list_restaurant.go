package ginuserlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
	userlikebiz "test/modules/userlike/biz"
	userlikemodel "test/modules/userlike/model"
	userlikestorage "test/modules/userlike/storage"
)

func ListRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter := userlikemodel.Filter{
			UserId: int(uid.GetLocalID()),
		}
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := userlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userlikebiz.NewListRestaurantLikedByUserBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewResponseSuccess(result, paging, filter))

	}
}
