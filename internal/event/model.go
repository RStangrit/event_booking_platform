package event

import "time"

// EventRequest defines the structure for creating or updating an event.
// Includes validation tags for input validation and JSON tags for serialization.
type EventRequest struct {
	ID          int      `json:"id" binding:"-"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Date        string   `json:"date" binding:"required,datetime=2006-01-02"`
	Location    string   `json:"location" binding:"required"`
	Capacity    int      `json:"capacity" binding:"required,min=1"`
	Price       *float64 `json:"price,omitempty"`
	CreatedBy   int      `json:"created_by" binding:"required"`
}

// EventResponse defines the structure for data sent to the client.
type EventResponse struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Date        string     `json:"date"`
	Location    string     `json:"location"`
	Capacity    int        `json:"capacity"`
	Price       *float64   `json:"price,omitempty"`
	CreatedBy   int        `json:"created_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// Event represents the domain model for an event, typically used with the database layer.
// Includes timestamps for creation and updates.
type Event struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Date        string     `json:"date"`
	Location    string     `json:"location"`
	Capacity    int        `json:"capacity"`
	Price       *float64   `json:"price,omitempty"`
	CreatedBy   int        `json:"created_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
