package restaurantstorage

import (
	"context"
	"gorm.io/gorm"
	"test/common"
	"test/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	//
	db := s.db
	var restaurant restaurantmodel.Restaurant

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&restaurant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &restaurant, nil
}
