package foodbiz

import (
	"context"
	"test/common"
	"test/modules/food/foodmodel"
)

type ListFoodStore interface {
	ListDataByCondition(ctx context.Context,
		condition map[string]interface{},
		filter *foodmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]foodmodel.Food, error)
}

type listFoodBiz struct {
	store ListFoodStore
}

func NewListFoodBiz(store ListFoodStore) *listFoodBiz {
	return &listFoodBiz{store: store}
}

func (biz *listFoodBiz) ListFood(ctx context.Context,
	filter *foodmodel.Filter,
	paging *common.Paging,
) ([]foodmodel.Food, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	return result, common.ErrCannotListEntity(foodmodel.EntityName, err)
}
