package component

import (
	"gorm.io/gorm"
	"test/component/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
}

func NewAppCtx(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secretKey: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
