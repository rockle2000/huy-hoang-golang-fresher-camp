package userlikebiz

import (
	"context"
	"test/common"
	userlikemodel "test/modules/userlike/model"
)

type ListRestaurantLikedByUserStore interface {
	GetRestaurantsLikedByUser(
		ctx context.Context,
		condition map[string]interface{},
		filter *userlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleRestaurant, error)
}

type listRestaurantLikedByUserBiz struct {
	store ListRestaurantLikedByUserStore
}

func NewListRestaurantLikedByUserBiz(store ListRestaurantLikedByUserStore) *listRestaurantLikedByUserBiz {
	return &listRestaurantLikedByUserBiz{store: store}
}
func (biz *listRestaurantLikedByUserBiz) ListRestaurant(
	ctx context.Context,
	filter *userlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleRestaurant, error) {

	restaurants, err := biz.store.GetRestaurantsLikedByUser(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(userlikemodel.EntityName, err)
	}
	return restaurants, nil
}
