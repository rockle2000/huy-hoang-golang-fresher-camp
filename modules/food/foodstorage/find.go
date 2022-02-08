package foodstorage

import (
	"context"
	"gorm.io/gorm"
	"test/common"
	"test/modules/food/foodmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*foodmodel.Food, error) {
	db := s.db
	var item foodmodel.Food

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}
	if err := db.Where(condition).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &item, nil
}
