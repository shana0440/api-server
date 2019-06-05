package repositories

import (
	"github.com/u-job/api-server/db"
	"github.com/u-job/api-server/db/entities"
)

// CreateUser use to save user to user table
func CreateUser(user entities.User) error {
	return db.Conn.Save(&user).Error
}

// GetUser use to find the user by user id
func GetUser(id int64) (*entities.User, error) {
	var user *entities.User
	err := db.Conn.First(user, "id = ?", id).Error
	return user, err
}
