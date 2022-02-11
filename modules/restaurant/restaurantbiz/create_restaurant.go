package restaurantbiz

import (
	"context"
	"test/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStorage interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}
type createRestaurantBiz struct {
	store CreateRestaurantStorage
}

func NewRestaurantBiz(store CreateRestaurantStorage) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
