package foodstorage

import (
	"context"
	"test/modules/food/foodmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *foodmodel.FoodCreate) error {
	db := s.db

	//Create(data) kp Create(&data) vì data đầu vào đã là con trỏ
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
