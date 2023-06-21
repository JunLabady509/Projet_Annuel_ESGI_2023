package subscribe

import (
	"gastroguru/database"
	"log"
)

type Subscription struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Period      int    `json:"period"`
}

// Retrieve All Subscriptions from DB
func GetAllSubscriptions() ([]Subscription, error) {
	rows, err := database.Db.Query("SELECT * FROM subscriptions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subscriptions := []Subscription{}

	for rows.Next() {
		var s Subscription
		if err := rows.Scan(&s.ID, &s.Name, &s.Type, &s.Price, &s.Description, &s.Period); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, s)
	}

	return subscriptions, nil
}

//Retrieve Details about one Subscription
func GetSubscriptionByID() (*Subscription, error) {
	s := &Subscription{}
	if err := database.Db.QueryRow("SELECT * FROM subscriptions WHERE id = ?", s.ID).Scan(&s.ID, &s.Name, &s.Type, &s.Price, &s.Description, &s.Period); err != nil {
		return nil, err
	}
	return s, nil
}

// Delete removes a given subscription record from the database
func DeleteSubscription(id string) error {
	_, err := database.Db.Exec("DELETE FROM subscriptions WHERE id = ?", id)
	return err
}

// Save updates or creates a given subscription record in the database
func (s *Subscription) Save() error {
	if err := s.Validate(); err != nil {
		return err
	}

	res, err := database.Db.Exec("INSERT INTO subscriptions (Name, Type, Price, Description, Period) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		s.Name, s.Type, s.Price, s.Description, s.Period)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	s.ID = int(id)

	return nil
}

func (s *Subscription) Validate() error {
	if s.Name == "" {
		return database.ErrRecordInvalid
	}
	return nil
}
