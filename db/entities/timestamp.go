package entities

import (
	"time"
)

// TimeStamp privode the created_at and updated_at two columns
type TimeStamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
