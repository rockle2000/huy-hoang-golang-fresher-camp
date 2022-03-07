package usermodel

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
	"test/common"
	"test/component/tokenprovider"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}
func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DBTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"-" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DBTypeUser)
}

func (u UserCreate) Validate() error {
	return validation.ValidateStruct(&u,
		// Email cannot be empty, and must be a valid email
		validation.Field(&u.Email, validation.Required, is.Email),
		// LastName cannot be empty, and contain only letter
		validation.Field(&u.LastName, validation.Required, is.UTFLetter),
		// FirstName cannot be empty, and contain only letter
		validation.Field(&u.FirstName, validation.Required, is.UTFLetter),
		// Phone cannot be empty, and must be a valid phone number
		validation.Field(&u.Phone, validation.Required, validation.Match(regexp.MustCompile("^(([+]?\\d{2})|\\d?)[\\s-]?[0-9]{2}[\\s-]?[0-9]{3}[\\s-]?[0-9]{4}$"))),
	)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrUserNameOrPasswordInvalid = common.NewCustomError(
		errors.New("invalid username or password"),
		"invalid username or password",
		"ErrUserNameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
