package userbiz

import (
	"context"
	"go.opencensus.io/trace"
	"test/common"
	"test/component"
	"test/component/tokenprovider"
	"test/modules/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

//type TokenConfig interface {
//	GetAtExp() int
//	//GetRtExp() int
//}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// S1: Find user,email
// S2: Hash pass from input and compare with password in DB
// S3: Provider: issue JWT token for client
// S3.1:Access token and refresh token
// S4: Return token(s)

func (biz *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {

	ctx1, span1 := trace.StartSpan(ctx, "user.biz.login")
	user, err := biz.storeUser.FindUser(ctx1, map[string]interface{}{"email": data.Email})
	span1.End()
	if err != nil {
		return nil, usermodel.ErrUserNameOrPasswordInvalid
	}

	_, span2 := trace.StartSpan(ctx, "user.biz.login.gen-jwt")
	// Ma hoa voi password tu input va salt trong db
	passwordHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passwordHashed {
		span2.End()
		return nil, usermodel.ErrUserNameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	span2.End()
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}

	//account := usermodel.NewAccount(accessToken, refreshToken)
	return accessToken, nil
}
