package foodstorage

import (
	"context"
	"test/common"
	"test/modules/food/foodmodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *foodmodel.FoodUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
