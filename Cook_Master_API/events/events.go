package events

import (
	"errors"
	"gastroguru/database"
	"os/user"
	"time"
)

type Event struct {
	ID           int         `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Date         time.Time   `json:"date"`
	Duration     int         `json:"duration"`
	Location     string      `json:"location"`
	Participants []user.User `json:"participants"`
	Capacity     int         `json:"capacity"`
	Price        float64     `json:"price"`
	Image        string      `json:"image"`
	Type         string      `json:"type"`
}

func GetAllEvents() ([]Event, error) {
	rows, err := database.Db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		tempDate := []uint8{}
		if err := rows.Scan(&e.ID, &e.Title, &e.Description,
			&tempDate, &e.Duration, &e.Location, &e.Capacity,
			&e.Price, &e.Image, &e.Type); err != nil {
			return nil, err
		}
		e.Date = convertDateFromDBtoStruct(tempDate)
		events = append(events, e)
	}

	return events, nil
}

func GetEvent(id string) (*Event, error) {
	e := &Event{}
	if err := database.Db.QueryRow("SELECT * FROM events WHERE id = ?", id).Scan(&e.ID,
		&e.Title, &e.Description, &e.Date,
		&e.Duration, &e.Location, &e.Capacity,
		&e.Price, &e.Image, &e.Type); err != nil {
		return nil, err
	}
	return e, nil
}

func DeleteEvent(id string) error {
	_, err := database.Db.Exec("DELETE FROM events WHERE id = ?", id)
	return err
}

func (e *Event) Save() error {
	if err := e.Validate(); err != nil {
		return err
	}

	_, err := database.Db.Exec("INSERT INTO events (id, title, description, date, duration, location, capacity, price, image, type) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		e.ID, e.Title, e.Description, e.Date, e.Duration, e.Location, e.Capacity, e.Price, e.Image, e.Type)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Validate() error {
	if e.Title == "" {
		return errors.New("title cannot be empty")
	}
	if e.Date.IsZero() {
		return errors.New("date must be valid")
	}
	if e.Duration <= 0 {
		return errors.New("duration must be greater than 0")
	}

	return nil
}

func GetStringAsTime(strTime string) (time.Time, error) {
	return time.Parse(time.RFC3339, strTime)
}

func convertDateFromDBtoStruct(date []uint8) time.Time {
	dateString := string(date)
	t, err := time.Parse("2006-01-02 15:04:05", dateString)
	if err != nil {
		panic(err)
	}
	return t
}
