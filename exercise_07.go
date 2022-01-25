package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Food struct {
	Id          int    `json:"id,omitempty" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

type FoodUpdate struct {
	Name        *string `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
}

func (Food) TableName() string {
	return "products"
}

func main() {
	dsn := os.Getenv("DBConnection")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	// INSERT
	newFood := Food{Name: "Test", Description: "Desc của Test"}
	if err := db.Create(&newFood); err != nil {
		fmt.Println(err)
	}
	fmt.Println(newFood)

	// SELECT
	var foods []Food
	db.Where("status = ?", 1).Find(&foods)
	fmt.Println("List Foods: ", foods)

	var food Food
	if err := db.Where("id = ?", 4).First(&food); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Food: ", food)

	//UPDATE
	food.Description = "ABC"
	db.Table(Food{}.TableName()).Updates(&food)

	//DELETE
	db.Table(Food{}.TableName()).Where("id = ?", 6).Delete(nil)

}
