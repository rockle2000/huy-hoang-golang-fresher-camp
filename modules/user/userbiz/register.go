package userbiz

import (
	"context"
	"test/common"
	usermodel "test/modules/user/usermodel"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{registerStorage: registerStorage, hasher: hasher}
}

func (biz *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {

	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}
	user, _ := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	//Email already existed
	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
