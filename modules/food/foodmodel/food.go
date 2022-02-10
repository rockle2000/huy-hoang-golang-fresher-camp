package foodmodel

import (
	"errors"
	"strings"
	"test/common"
)

const EntityName = "Food"

type Food struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Description     string         `json:"description" gorm:"column:description"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (Food) TableName() string {
	return "sanpham"
}

type FoodUpdate struct {
	Name        *string        `json:"name" gorm:"column:name"`
	Description *string        `json:"description" gorm:"column:description"`
	Logo        *common.Image  `json:"logo" gorm:"column:logo"`
	Cover       *common.Images `json:"cover" gorm:"column:cover"`
}

func (FoodUpdate) TableName() string {
	return "sanpham"
}

type FoodCreate struct {
	Name        string         `json:"name" gorm:"column:name"`
	Description string         `json:"description" gorm:"column:description"`
	Logo        *common.Image  `json:"logo" gorm:"column:logo"`
	Cover       *common.Images `json:"cover" gorm:"column:cover"`
}

func (res *FoodCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return errors.New("food's name cannot be null")
	}
	return nil
}

func (FoodCreate) TableName() string {
	return "sanpham"
}
