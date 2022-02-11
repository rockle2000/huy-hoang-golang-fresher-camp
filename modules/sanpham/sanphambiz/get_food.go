package sanphambiz

import (
	"context"
	"test/common"
	"test/modules/sanpham/sanphammodel"
)

type GetFoodStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*sanphammodel.Food, error)
}

type getFoodBiz struct {
	store GetFoodStore
}

func NewGetFoodBiz(store GetFoodStore) *getFoodBiz {
	return &getFoodBiz{store: store}
}

func (biz *getFoodBiz) GetFood(ctx context.Context, id int) (*sanphammodel.Food, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(sanphammodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(sanphammodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(sanphammodel.EntityName, nil)
	}
	return data, nil
}
