package user

import "gorm.io/gorm"

type user struct {
	gorm.Model
	Name string
}
