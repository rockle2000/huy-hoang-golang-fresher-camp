package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/common"
	"test/component"
	"test/component/hasher"
	"test/component/tokenprovider/jwt"
	"test/modules/user/userbiz"
	"test/modules/user/usermodel"
	"test/modules/user/userstorage"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin
		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleResponseSuccess(account))
	}
}
