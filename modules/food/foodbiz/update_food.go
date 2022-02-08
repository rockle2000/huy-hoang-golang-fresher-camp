package foodbiz

import (
	"context"
	"test/common"
	"test/modules/food/foodmodel"
)

type UpdateFoodStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*foodmodel.Food, error)
	UpdateData(ctx context.Context, id int, data *foodmodel.FoodUpdate) error
}

type updateFoodBiz struct {
	store UpdateFoodStore
}

func NewUpdateFoodBiz(store UpdateFoodStore) *updateFoodBiz {
	return &updateFoodBiz{store: store}
}
func (biz *updateFoodBiz) UpdateFood(ctx context.Context, id int, data *foodmodel.FoodUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(foodmodel.EntityName, nil)
	}
	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(foodmodel.EntityName, err)
	}
	return nil
}
