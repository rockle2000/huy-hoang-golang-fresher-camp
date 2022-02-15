package common

const (
	DBTypeRestaurant = 1
	DBTypeFood       = 2
	DBTypeCategory   = 3
	DBTypeUser       = 4
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
