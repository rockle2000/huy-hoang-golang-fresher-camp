package foodstorage

import (
	"context"
	"test/common"
	"test/modules/food/foodmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	condition map[string]interface{},
	filter *foodmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]foodmodel.Food, error) {
	db := s.db
	var result []foodmodel.Food

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	db = db.Table(foodmodel.Food{}.TableName()).Where(condition)

	if v := filter; v != nil {
		if v.Status > 0 {
			db = db.Where("status = ?", v.Status)
		}
	}

	if err := db.Table(foodmodel.Food{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
