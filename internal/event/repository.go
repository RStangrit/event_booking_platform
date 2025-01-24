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

	result, err := statement.Exec(event.Title, event.Description, event.Date, event.Location, event.Capacity, &event.Price, event.Created_by, currentTime)
	if err != nil {
		return err
	}
	eventId, err := result.LastInsertId()
	event.ID = int(eventId)
	return err
}

func getAllEvents() ([]EventResponse, error) {
	var events []EventResponse
	query := "SELECT id, title, description, date, location, capacity, price, created_by, created_at, updated_at FROM events"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event EventResponse
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location, &event.Capacity, &event.Price, &event.Created_by, &event.Created_at, &event.Updated_at)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func getOneEvent(eventId int64) (*EventResponse, error) {
	query := "SELECT id, title, description, date, location, capacity, price, created_by, created_at FROM events WHERE id = ?"
	row := database.DB.QueryRow(query, eventId)

	var event EventResponse
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location, &event.Capacity, &event.Price, &event.Created_by, &event.Created_at)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event EventResponse) Update() error {
	query := `
	UPDATE events
	SET title = ?, description = ?, date = ?, location = ?, capacity = ?, price = ?, created_by = ?, updated_at = ?
	WHERE id = ?
	`
	statement, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	currentTime := util.GetCurrentTime()

	_, err = statement.Exec(event.Title, event.Description, event.Date, event.Location, event.Capacity, event.Price, event.Created_by, currentTime, event.ID)
	return err
}

func (event EventResponse) Delete() error {
	query := "DELETE FROM events WHERE id = ? "
	statement, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(event.ID)
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
