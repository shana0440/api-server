package entities

// User is the example for showing how to use gorm
type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	TimeStamp
}

// TableName is for gorm to know the table name of the user entity
func (User) TableName() string {
	return "users"
}
