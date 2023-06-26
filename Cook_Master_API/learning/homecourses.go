package learning

import (
	"errors"
	"fmt"
	"gastroguru/database"
	"log"
	"time"
)

type HomeCourse struct {
	Learning
	Duration time.Duration `json:"duration"`
	Capacity int           `json:"capacity"`
}

func GetAllHomeCourses() ([]HomeCourse, error) {
	rows, err := database.Db.Query("SELECT * FROM home_courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	homeCourses := []HomeCourse{}

	for rows.Next() {
		var h HomeCourse
		if err := rows.Scan(&h.ID, &h.Title, &h.Description, &h.Instructor_ID, &h.Price, &h.Start_Time, &h.End_Time, &h.Duration, &h.Capacity); err != nil {
			return nil, err
		}
		homeCourses = append(homeCourses, h)
	}
	return homeCourses, nil
}

func GetHomeCourse(id string) (*HomeCourse, error) {
	h := &HomeCourse{}
	if err := database.Db.QueryRow("SELECT * FROM home_courses WHERE id = ?", id).Scan(&h.ID, &h.Title, &h.Description, &h.Instructor_ID, &h.Price, &h.Start_Time, &h.End_Time, &h.Duration, &h.Capacity); err != nil {
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

	res, err := database.Db.Exec("INSERT INTO home_courses (title, description, instructor_id, price, start_time, end_time, duration, capacity) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", h.Title, h.Description, h.Instructor_ID, h.Price, h.Start_Time, h.End_Time, h.Duration, h.Capacity)
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
		return errors.New("title cannot be empty")
	}
	if h.Description == "" {
		return errors.New("description cannot be empty")
	}
	if h.Instructor_ID == 0 {
		return errors.New("instructor_ID cannot be empty")
	}
	if h.Price == 0 {
		return errors.New("price cannot be empty")
	}
	if h.Start_Time == "" || h.End_Time == "" {
		return errors.New("start_time and end_time must be valid dates")
	}
	if h.Duration == 0 {
		return errors.New("duration cannot be empty")
	}
	if h.Capacity == 0 {
		return errors.New("capacity cannot be empty")
	}
	return nil
}

func (h *HomeCourse) GetStringTimeAsTime(strTime string) (time.Time, error) {
	return time.Parse(time.RFC3339, strTime)
}
