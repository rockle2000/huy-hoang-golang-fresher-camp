package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"test/common"
	"test/component"
	"test/component/uploadprovider"
	"test/middleware"
	"test/modules/restaurant/restauranttransport/ginrestaurant"
	"test/modules/restaurantlike/transport/ginrestaurantlike"
	"test/modules/sanpham/sanphamtransport/ginsanpham"
	"test/modules/upload/uploadtransport/ginupload"
	"test/modules/user/usertransport/ginuser"

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

	secretKey := os.Getenv("SYSTEM_SECRET")
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}
}
func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {

	appCtx := component.NewAppCtx(db, upProvider, secretKey)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appCtx))

	v1.POST("/login", ginuser.Login(appCtx))
	v1.POST("/register", ginuser.Register(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	food := v1.Group("/foods")
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

	restaurant := v1.Group("/restaurants", middleware.RequiredAuth(appCtx))
	{
		restaurant.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurant.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurant.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurant.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))
	}

	v1.GET("/encode-uid", gin.HandlerFunc(func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}
		var a reqData
		c.ShouldBind(&a)
		c.JSON(http.StatusOK, gin.H{
			"id": common.NewUID(uint32(a.RealId), a.DbType, 1),
		})

	}))
	return r.Run()
}
