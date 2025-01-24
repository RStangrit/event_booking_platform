package users

import (
	"time"
)

// UserRequest defines the data structure for creating or updating a user.
// It includes validation tags for input validation and JSON tags for serialization.
type UserRequest struct {
	Id       int    `json:"id" binding:"-"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=admin user guest"`
}

// UserResponse defines the structure for user data sent to the client.
type UserResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// User defines the domain model for a user, typically used with the database layer.
// It includes timestamps for creation and updates.
type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Omit password in JSON responses
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
