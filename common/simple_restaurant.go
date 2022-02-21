package common

type SimpleRestaurant struct {
	SQLModel `json:",inline"`
	Name     string `json:"name" gorm:"column:name"`
	Address  string `json:"address" gorm:"column:addr"`
	Logo     *Image `json:"logo" gorm:"column:logo"`
	Cover    *Image `json:"cover" gorm:"cover"`
}

func (SimpleRestaurant) TableName() string {
	return "restaurants"
}
func (u *SimpleRestaurant) Mask(isAdmin bool) {
	u.GenUID(DBTypeRestaurant)
}
