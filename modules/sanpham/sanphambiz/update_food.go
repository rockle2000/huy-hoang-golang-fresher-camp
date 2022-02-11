package sanphambiz

import (
	"context"
	"test/common"
	"test/modules/sanpham/sanphammodel"
)

type UpdateFoodStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*sanphammodel.Food, error)
	UpdateData(ctx context.Context, id int, data *sanphammodel.FoodUpdate) error
}

type updateFoodBiz struct {
	store UpdateFoodStore
}

func NewUpdateFoodBiz(store UpdateFoodStore) *updateFoodBiz {
	return &updateFoodBiz{store: store}
}
func (biz *updateFoodBiz) UpdateFood(ctx context.Context, id int, data *sanphammodel.FoodUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(sanphammodel.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(sanphammodel.EntityName, nil)
	}
	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(sanphammodel.EntityName, err)
	}
	return nil
}
