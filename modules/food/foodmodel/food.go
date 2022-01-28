package foodmodel

import (
	"errors"
	"strings"
)

type Food struct {
	Id          int    `json:"id,omitempty" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (Food) TableName() string {
	return "products"
}

type FoodUpdate struct {
	Name        *string `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
}

func (FoodUpdate) TableName() string {
	return "products"
}

type FoodCreate struct {
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (res *FoodCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return errors.New("food's name cannot be null")
	}
	return nil
}

func (FoodCreate) TableName() string {
	return "products"
}
