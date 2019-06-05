package entities

import (
	"github.com/jinzhu/gorm"
)

// User is the example for showing how to use gorm
type User struct {
	gorm.Model
	ID    int64
	Email string
}

// TableName is for gorm to know the table name of the user entity
func (User) TableName() string {
	return "users"
}
