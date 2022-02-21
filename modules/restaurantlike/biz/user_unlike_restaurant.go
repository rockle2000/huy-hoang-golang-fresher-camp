package restaurantlikebiz

import (
	"context"
	"errors"
	"test/common"
	restaurantlikemodel "test/modules/restaurantlike/model"
)

type UserUnlikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
	Find(ctx context.Context, condition map[string]interface{}) (*restaurantlikemodel.Like, error)
}

type userUnlikeRestaurantBiz struct {
	store    UserUnlikeRestaurantStore
	decStore DecreaseLikeCountStore
}

type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore, decStore DecreaseLikeCountStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store: store, decStore: decStore}
}

func (biz *userUnlikeRestaurantBiz) UserUnlikeRestaurant(
	ctx context.Context,
	userId, restaurantId int,
) error {

	if data, _ := biz.store.Find(ctx, map[string]interface{}{"restaurant_id": restaurantId, "user_id": userId}); data == nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(errors.New("you did not like this restaurant"))
	}

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		_ = biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	}()
	return nil
}
