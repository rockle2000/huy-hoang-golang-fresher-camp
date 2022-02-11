package sanphamstorage

import (
	"context"
	"test/common"
	"test/modules/sanpham/sanphammodel"
)

func (s *sqlStore) Create(ctx context.Context, data *sanphammodel.FoodCreate) error {
	db := s.db

	//Create(data) kp Create(&data) vì data đầu vào đã là con trỏ
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
