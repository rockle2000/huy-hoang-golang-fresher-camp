package sanphambiz

import (
	"context"
	"test/common"
	"test/modules/sanpham/sanphammodel"
)

type DeleteFoodStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*sanphammodel.Food, error)
	SoftDeleteData(ctx context.Context, id int) error
}
type deleteFoodBiz struct {
	store DeleteFoodStore
}

func NewDeleteFoodBiz(store DeleteFoodStore) *deleteFoodBiz {
	return &deleteFoodBiz{store: store}
}

func (biz *deleteFoodBiz) DeleteFood(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(sanphammodel.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(sanphammodel.EntityName, nil)
	}
	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(sanphammodel.EntityName, err)
	}
	return nil
}
