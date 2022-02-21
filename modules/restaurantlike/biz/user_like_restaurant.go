package restaurantlikebiz

import (
	"context"
	"errors"
	"test/common"
	restaurantlikemodel "test/modules/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
	Find(ctx context.Context, condition map[string]interface{}) (*restaurantlikemodel.Like, error)
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncreaseLikeCountStore
}

type IncreaseLikeCountStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

func NewUserLikeRestaurantStore(store UserLikeRestaurantStore, incStore IncreaseLikeCountStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, incStore: incStore}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {

	if data, _ := biz.store.Find(ctx, map[string]interface{}{"restaurant_id": data.RestaurantId, "user_id": data.UserId}); data != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(errors.New("you had liked this restaurant"))
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		_ = biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	}()

	return nil
}
