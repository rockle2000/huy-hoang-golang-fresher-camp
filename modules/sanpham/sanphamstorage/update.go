package sanphamstorage

import (
	"context"
	"test/common"
	"test/modules/sanpham/sanphammodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *sanphammodel.FoodUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
