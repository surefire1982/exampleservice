package entity

import "time"

// User entity
type User struct {
	UserID    string    `json:"userid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
