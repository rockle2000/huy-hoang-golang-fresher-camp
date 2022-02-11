package sanphambiz

import (
	"context"
	"test/common"
	"test/modules/sanpham/sanphammodel"
)

type ListFoodStore interface {
	ListDataByCondition(ctx context.Context,
		condition map[string]interface{},
		filter *sanphammodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]sanphammodel.Food, error)
}

type listFoodBiz struct {
	store ListFoodStore
}

func NewListFoodBiz(store ListFoodStore) *listFoodBiz {
	return &listFoodBiz{store: store}
}

func (biz *listFoodBiz) ListFood(ctx context.Context,
	filter *sanphammodel.Filter,
	paging *common.Paging,
) ([]sanphammodel.Food, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	return result, common.ErrCannotListEntity(sanphammodel.EntityName, err)
}
