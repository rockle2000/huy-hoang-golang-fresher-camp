package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"

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
	food := r.Group("/foods")
	{
		// create new food
		food.POST("", func(c *gin.Context) {
			var newFood Food

			if err := c.ShouldBind(&newFood); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			if err := db.Create(&newFood).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, newFood)
		})

		// Get all food
		food.GET("", func(c *gin.Context) {
			var listFood []Food
			type Filter struct {
				Status int `json:"status" form:"status"`
			}
			var filter Filter
			if err := c.ShouldBind(&filter); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			newDB := db
			if filter.Status > 0 {
				newDB = db.Where("status = ?", filter.Status)
			}
			if err := newDB.Find(&listFood).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, listFood)
		})

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
