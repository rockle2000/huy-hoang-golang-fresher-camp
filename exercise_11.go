package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"test/component"
	"test/component/uploadprovider"
	"test/middleware"
	"test/modules/food/foodtransport/ginfood"
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
