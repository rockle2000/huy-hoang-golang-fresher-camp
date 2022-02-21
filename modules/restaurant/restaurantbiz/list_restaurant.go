package restaurantbiz

import (
	"context"
	"test/common"
	"test/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		condition map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBiz struct {
	store    ListRestaurantStore
	likStore LikeStore
}

func NewListRestaurantBiz(store ListRestaurantStore, likeStore LikeStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likStore: likeStore}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	//ids := make([]int, len(result))
	//
	//for i := range result {
	//	ids[i] = result[i].Id
	//}
	//
	//mapResLike, err := biz.likStore.GetRestaurantLike(ctx, ids)
	//if err != nil {
	//	log.Println("Cannot get restaurant likes: ", err)
	//}
	//
	//if v := mapResLike; v != nil {
	//	for i, item := range result {
	//		result[i].LikeCount = mapResLike[item.Id]
	//	}
	//}

	return result, nil
}
