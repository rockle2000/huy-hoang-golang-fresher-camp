package userlikestorage

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"test/common"
	userlikemodel "test/modules/userlike/model"
	"time"
)

const TimeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetRestaurantsLikedByUser(
	ctx context.Context,
	condition map[string]interface{},
	filter *userlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleRestaurant, error) {
	db := s.db
	var result []userlikemodel.Like

	db = db.Table(userlikemodel.Like{}.TableName()).Where(condition)
	if v := filter; v != nil {
		if v.UserId > 0 {
			db = db.Where("user_id = ?", v.UserId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("Restaurant")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(TimeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))

	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	restaurants := make([]common.SimpleRestaurant, len(result))

	for i, item := range result {
		result[i].Restaurant.CreatedAt = item.CreatedAt
		result[i].Restaurant.UpdatedAt = nil
		restaurants[i] = *result[i].Restaurant
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(TimeLayout))))
			paging.NextCursor = cursorStr
		}
	}
	return restaurants, nil
}
