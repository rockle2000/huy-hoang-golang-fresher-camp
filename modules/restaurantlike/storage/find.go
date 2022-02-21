package restaurantlikestorage

import (
	"context"
	"test/common"
	restaurantlikemodel "test/modules/restaurantlike/model"
)

func (s *sqlStore) Find(ctx context.Context, condition map[string]interface{}) (*restaurantlikemodel.Like, error) {
	db := s.db

	var oldData restaurantlikemodel.Like

	if err := db.
		Where(condition).
		First(&oldData).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &oldData, nil
}
