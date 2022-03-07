package restaurantlikebiz

import (
	"context"
	"errors"
	"go.opencensus.io/trace"
	"test/common"
	"test/component/asyncjob"
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
	_, span := trace.StartSpan(ctx, "restaurant.biz.like")
	span.AddAttributes(
		trace.Int64Attribute("restaurant_id", int64(data.RestaurantId)),
		trace.Int64Attribute("user_id", int64(data.UserId)),
	)

	defer span.End()
	if data, _ := biz.store.Find(ctx, map[string]interface{}{"restaurant_id": data.RestaurantId, "user_id": data.UserId}); data != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(errors.New("you had liked this restaurant"))
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
