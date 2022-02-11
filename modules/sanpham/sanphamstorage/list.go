package sanphamstorage

import (
	"context"
	"test/common"
	"test/modules/sanpham/sanphammodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	condition map[string]interface{},
	filter *sanphammodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]sanphammodel.Food, error) {
	db := s.db
	var result []sanphammodel.Food

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	db = db.Table(sanphammodel.Food{}.TableName()).Where(condition).Where("status in (1)")

	if v := filter; v != nil {
		if v.Status > 0 {
			db = db.Where("status = ?", v.Status)
		}
	}

	if err := db.Table(sanphammodel.Food{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
