package sanphambiz

import (
	"context"
	"test/modules/sanpham/sanphammodel"
)

type CreateFoodStore interface {
	Create(ctx context.Context, data *sanphammodel.FoodCreate) error
}

type createFoodBiz struct {
	store CreateFoodStore
}

func NewCreateFoodBiz(store CreateFoodStore) *createFoodBiz {
	return &createFoodBiz{store: store}
}

func (biz *createFoodBiz) CreateFood(ctx context.Context, data *sanphammodel.FoodCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	err := biz.store.Create(ctx, data)
	return err
}
