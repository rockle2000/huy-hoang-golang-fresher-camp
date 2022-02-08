package foodstorage

import (
	"context"
	"test/common"
	"test/modules/food/foodmodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db
	if err := db.Table(foodmodel.Food{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
