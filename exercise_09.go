package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"test/component"
	"test/modules/food/foodtransport/ginfood"

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
func (FoodUpdate) TableName() string {
	return "products"
}

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
	r := gin.Default()

	appCtx := component.NewAppCtx(db)
	food := r.Group("/foods")
	{
		// create new food
		food.POST("", ginfood.CreateFood(appCtx))

		// Get all food
		food.GET("", ginfood.ListFood(appCtx))

		//Get food by id
		food.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			var item Food
			if err := db.Where("id = ?", id).First(&item).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, item)
		})

		//Update food by id
		food.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			var item FoodUpdate
			if err := c.ShouldBind(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			if err := db.Where("id = ?", id).Updates(&item); err.RowsAffected < 1 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": errors.New("update failed"),
				})
				return
			}
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Update successfully",
			})
		})

		//Delete food by id
		food.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			if err := db.Table(Food{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Delete successfully",
			})
		})
	}
	return r.Run()
}
