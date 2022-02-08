package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"test/component"
	"test/middleware"
	"test/modules/food/foodtransport/ginfood"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DBConnection")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}
func runService(db *gorm.DB) error {
	appCtx := component.NewAppCtx(db)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	food := r.Group("/foods")
	{
		// create new food
		food.POST("", ginfood.CreateFood(appCtx))
		// Get all food
		food.GET("", ginfood.ListFood(appCtx))
		//Get food by id
		food.GET("/:id", ginfood.GetFood(appCtx))
		//Update food by id
		food.PATCH("/:id", ginfood.UpdateFood(appCtx))
		//Delete food by id
		food.DELETE("/:id", ginfood.DeleteFood(appCtx))
	}
	return r.Run()
}
