package restaurantmodel

import (
	"errors"
	"strings"
	"test/common"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name"`
	UserId          int                `json:"-" gorm:"column:owner_id"`
	Address         string             `json:"address" gorm:"column:addr"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo"`
	Cover           *common.Image      `json:"cover" gorm:"cover"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false"`
	LikeCount       int                `json:"liked_count" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Name            *string       `json:"name" gorm:"column:name"`
	Address         *string       `json:"address" gorm:"column:addr"`
	Logo            *common.Image `json:"logo" gorm:"column:logo"`
	Cover           *common.Image `json:"cover" gorm:"column:cover"`
	CityId          int           `json:"cityId" gorm:"column:city_id"`
}

func (RestaurantUpdate) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name"`
	Address         string        `json:"address" gorm:"column:addr"`
	OwnerId         int           `json:"-" gorm:"column:owner_id"`
	Logo            *common.Image `json:"logo" gorm:"column:logo"`
	Cover           *common.Image `json:"cover" gorm:"cover"`
	CityId          int           `json:"cityId" gorm:"column:city_id"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "restaurant name cannot be null", "ErrNamCannotBeEmpty")
)

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if res.Name == "" {
		return errors.New("restaurant's name cannot be null")
	}
	return nil
}

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DBTypeRestaurant)
	if u := data.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}
