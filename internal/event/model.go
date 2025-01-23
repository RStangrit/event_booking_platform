package event

import "time"

type EventRequest struct {
	ID          int
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Date        string `binding:"required"`
	Location    string `binding:"required"`
	Capacity    int    `binding:"required"`
	Price       *float64
	Created_by  int `binding:"required"`
	Created_at  time.Time
	Updated_at  *time.Time
}

type EventResponse struct {
	ID          int
	Title       string
	Description string
	Date        string
	Location    string
	Capacity    int
	Price       *float64
	Created_by  int
	Created_at  time.Time
	Updated_at  *time.Time
}
