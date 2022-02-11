package restaurantlikemodel

import "time"

type Like struct {
	RestaurantId int        `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int        `json:"user_id" gorm:"column:restaurant_id"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (l Like) TableName() string {
	return "restaurant_likes"
}
