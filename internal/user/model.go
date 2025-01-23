package users

import "time"

// to describe the data structure (for example, user).
type User struct {
	Id         int
	Name       string `binding:"required"`
	Email      string `binding:"required"`
	Password   string `binding:"required"`
	Role       string `binding:"required"`
	Created_at time.Time
	Updated_at time.Time
}
