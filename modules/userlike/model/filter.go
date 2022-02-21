package userlikemodel

type Filter struct {
	UserId       int `json:"-" form:"user_id"`
	RestaurantId int `json:"-" form:"restaurant_id"`
}
