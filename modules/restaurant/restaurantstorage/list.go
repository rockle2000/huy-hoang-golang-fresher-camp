package restaurantstorage

import (
	"context"
	"test/common"
	"test/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {

	db := s.db
	var result []restaurantmodel.Restaurant

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}
	db = db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where(condition).
		Where("status in (1)")

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", filter.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
