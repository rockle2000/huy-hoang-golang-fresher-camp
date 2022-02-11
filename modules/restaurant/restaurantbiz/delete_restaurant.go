package restaurantbiz

import (
	"context"
	"test/common"
	"test/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	SoftDeleteData(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, err)
	}
	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
