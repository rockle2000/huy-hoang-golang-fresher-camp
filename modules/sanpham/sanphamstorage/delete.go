package sanphamstorage

import (
	"context"
	"test/common"
	"test/modules/sanpham/sanphammodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db
	if err := db.Table(sanphammodel.Food{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
