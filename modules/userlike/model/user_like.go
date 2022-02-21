package userlikemodel

import (
	"test/common"
	"time"
)

const EntityName = "UserLike"

type Like struct {
	UserId       int                      `json:"user_id" gorm:"column:user_id"`
	RestaurantId int                      `json:"restaurant_id" gorm:"column:restaurant_id"`
	CreatedAt    *time.Time               `json:"created_at" gorm:"column:created_at"`
	Restaurant   *common.SimpleRestaurant `json:"restaurant" gorm:"preload:false"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}
