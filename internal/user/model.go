package users

import "time"

// to describe the data structure (for example, user).
type UserRequest struct {
	Id         int
	Name       string `binding:"required"`
	Email      string `binding:"required"`
	Password   string `binding:"required"`
	Role       string `binding:"required"`
	Created_at time.Time
	Updated_at *time.Time
}

type UserResponse struct {
	Id         int
	Name       string
	Email      string
	Created_at time.Time
}
