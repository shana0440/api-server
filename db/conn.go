package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// Use to connect to postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/u-job/api-server/config"
)

// Conn is use to get the database connection
var Conn *gorm.DB

// CreateConn use to create the connection with database
func CreateConn(c *config.DB) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", c.Host, c.Port, c.User, c.Name, c.Password, c.SslMode)
	conn, err := gorm.Open("postgres", dsn)
	if err != nil {
		return err
	}
	conn.DB().SetMaxIdleConns(20)
	Conn = conn
	return nil
}

// Close use to close database connection
func Close() {
	Conn.Close()
}
