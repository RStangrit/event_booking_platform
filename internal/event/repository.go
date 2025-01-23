package event

import (
	"errors"
	"main/pkg/database"
	"main/pkg/util"
)

// to interact with the database.
func (event *EventRequest) Save() error {
	unique, _ := IsTitleUnique(event.Title)
	if !unique {
		return errors.New("event title already exists")
	}

	query := "INSERT INTO events (title, description, date, location, capacity, price, created_by, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	statement, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	currentTime := util.GetCurrentTime()

	result, err := statement.Exec(event.Title, event.Description, event.Date, event.Location, event.Capacity, event.Price, event.Created_by, currentTime)
	if err != nil {
		return err
	}
	eventId, err := result.LastInsertId()
	event.ID = int(eventId)
	return err
}

func checkTitlePresence(title string) (count int, err error) {
	query := "SELECT COUNT(*) FROM events WHERE title = ?"
	err = database.DB.QueryRow(query, title).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, err
}
