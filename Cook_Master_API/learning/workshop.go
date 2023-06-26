package learning

import (
	"errors"
	"fmt"
	"gastroguru/database"
	"log"
	"time"
)

type Workshop struct {
	Learning
	Capacity int    `json:"capacity"`
	Place    string `json:"place"`
}

func GetAllWorkshops() ([]Workshop, error) {
	rows, err := database.Db.Query("SELECT * FROM workshops")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	workshops := []Workshop{}

	for rows.Next() {
		var w Workshop
		if err := rows.Scan(&w.ID, &w.Title, &w.Description, &w.Instructor_ID, &w.Capacity, &w.Price, &w.Start_Time, &w.End_Time, &w.Place); err != nil {
			return nil, err
		}
		workshops = append(workshops, w)
	}

	return workshops, nil
}

func GetWorkshop(id string) (*Workshop, error) {
	w := &Workshop{}
	if err := database.Db.QueryRow("SELECT * FROM workshops WHERE id = ?", id).Scan(&w.ID, &w.Title, &w.Description, &w.Instructor_ID, &w.Capacity, &w.Price, &w.Start_Time, &w.End_Time, &w.Place); err != nil {
		return nil, err
	}
	return w, nil
}

func DeleteWorkshop(id string) error {
	_, err := database.Db.Exec("DELETE FROM workshops WHERE id = ?", id)
	return err
}

func (w *Workshop) Save() error {
	if err := w.Validate(); err != nil {
		return err
	}

	res, err := database.Db.Exec("INSERT INTO workshops (Title, Description, Instructor_ID, Capacity, Price, Start_Time, End_Time, Place) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		w.Title, w.Description, w.Instructor_ID, w.Capacity, w.Price, w.Start_Time, w.End_Time, w.Place)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	w.ID = int(id)
	return nil
}

func (w *Workshop) Validate() error {
	if w.Title == "" {
		return errors.New("title cannot be empty")
	}
	if w.Description == "" {
		return errors.New("description cannot be empty")
	}
	if w.Instructor_ID == 0 {
		return errors.New("instructor_id must be greater than 0")
	}
	if w.Place == "" {
		return errors.New("location cannot be empty")
	}
	if w.Capacity <= 0 {
		return errors.New("capacity must be greater than 0")
	}
	if w.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if w.Start_Time == "" || w.End_Time == "" {
		return errors.New("start_time and end_time must be valid dates")
	}
	return nil
}

func GetStringTimeAsTime(strTime string) (time.Time, error) {
	return time.Parse(time.RFC3339, strTime)
}
