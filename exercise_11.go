package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"test/component"
	"test/component/uploadprovider"
	"test/middleware"
	"test/modules/restaurant/restauranttransport/ginrestaurant"
	"test/modules/sanpham/sanphamtransport/ginsanpham"
	"test/modules/upload/uploadtransport/ginupload"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DBConnection")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}
func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {

	appCtx := component.NewAppCtx(db, upProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.POST("/upload", ginupload.Upload(appCtx))
	food := r.Group("/foods")
	{
		// create new sanpham
		food.POST("", ginsanpham.CreateFood(appCtx))
		// Get all sanpham
		food.GET("", ginsanpham.ListFood(appCtx))
		//Get sanpham by id
		food.GET("/:id", ginsanpham.GetFood(appCtx))
		//Update sanpham by id
		food.PATCH("/:id", ginsanpham.UpdateFood(appCtx))
		//Delete sanpham by id
		food.DELETE("/:id", ginsanpham.DeleteFood(appCtx))
	}

	restaurant := r.Group("/restaurants")
	{
		restaurant.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurant.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurant.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}
	return r.Run()
}
