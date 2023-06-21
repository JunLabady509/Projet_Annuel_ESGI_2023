package homecourses

import (
	"errors"
	"fmt"
	"gastroguru/database"
	"log"
)

type HomeCourse struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Instructor_ID int     `json:"instructor_id"`
	Client_ID     int     `json:"client_id"`
	Location      string  `json:"location"`
	Price         float64 `json:"price"`
	Start_Time    string  `json:"start_time"`
	End_Time      string  `json:"end_time"`
}

func GetAllHomeCourses() ([]HomeCourse, error) {

	rows, err := database.Db.Query("SELECT * FROM home_courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	home_courses := []HomeCourse{}

	for rows.Next() {
		var h HomeCourse
		if err := rows.Scan(&h.ID, &h.Title, &h.Description,
			&h.Instructor_ID, &h.Client_ID, &h.Location, &h.Price, &h.Start_Time, &h.End_Time); err != nil {
			return nil, err
		}
		home_courses = append(home_courses, h)
	}
	return home_courses, nil
}

func GetHomeCourse(id string) (*HomeCourse, error) {
	h := &HomeCourse{}
	if err := database.Db.QueryRow("SELECT * FROM home_courses WHERE id = ?", id).Scan(&h.ID, &h.Title, &h.Description, &h.Instructor_ID, &h.Client_ID, &h.Location, &h.Price, &h.Start_Time, &h.End_Time); err != nil {
		return nil, err
	}
	return h, nil
}

func DeleteHomeCourse(id string) error {
	_, err := database.Db.Exec("DELETE FROM home_courses WHERE id = ?", id)
	return err
}

func (h *HomeCourse) Save() error {
	if err := h.Validate(); err != nil {
		return err
	}

	res, err := database.Db.Exec("INSERT INTO home_courses VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", h.ID, h.Title, h.Description, h.Instructor_ID, h.Client_ID, h.Location, h.Price, h.Start_Time, h.End_Time)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	h.ID = int(id)

	return nil
}

func (h *HomeCourse) Validate() error {
	if h.Title == "" {
		return errors.New("ditle cannot be empty")
	}
	if h.Description == "" {
		return errors.New("description cannot be empty")
	}
	if h.Instructor_ID == 0 {
		return errors.New("instructor_ID cannot be empty")
	}
	if h.Client_ID == 0 {
		return errors.New("client_ID cannot be empty")
	}
	if h.Location == "" {
		return errors.New("location cannot be empty")
	}
	if h.Price == 0 {
		return errors.New("price cannot be empty")
	}
	if h.Start_Time == "" {
		return errors.New("Start_Time cannot be empty")
	}
	if h.End_Time == "" {
		return errors.New("End_Time cannot be empty")
	}
	return nil
}
