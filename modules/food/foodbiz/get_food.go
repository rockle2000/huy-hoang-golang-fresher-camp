package foodbiz

import (
	"context"
	"test/common"
	"test/modules/food/foodmodel"
)

type GetFoodStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*foodmodel.Food, error)
}

type getFoodBiz struct {
	store GetFoodStore
}

func NewGetFoodBiz(store GetFoodStore) *getFoodBiz {
	return &getFoodBiz{store: store}
}

func (biz *getFoodBiz) GetFood(ctx context.Context, id int) (*foodmodel.Food, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(foodmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(foodmodel.EntityName, nil)
	}
	return data, nil
}
