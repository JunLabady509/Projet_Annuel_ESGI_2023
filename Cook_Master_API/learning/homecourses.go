package learning

import (
	"errors"
	"gastroguru/database"
	"log"
	"time"
)

type HomeCourse struct {
	Learning
	Duration time.Duration `json:"duration"`
	Address  string        `json:"address"`
}

func GetAllHomeCourses() ([]HomeCourse, error) {
	rows, err := database.Db.Query("SELECT * FROM home_courses")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	homeCourses := []HomeCourse{}

	for rows.Next() {
		var hc HomeCourse
		if err := rows.Scan(&hc.ID, &hc.Title, &hc.Description, &hc.Price, &hc.Start_Time, &hc.End_Time, &hc.Instructor_ID, &hc.Duration, &hc.Address); err != nil {
			return nil, err
		}
		homeCourses = append(homeCourses, hc)
	}

	return homeCourses, nil
}

func GetHomeCourse(id string) (*HomeCourse, error) {
	hc := &HomeCourse{}
	if err := database.Db.QueryRow("SELECT * FROM home_courses WHERE id = ?", id).Scan(&hc.ID, &hc.Title, &hc.Description, &hc.Price, &hc.Start_Time, &hc.End_Time, &hc.Instructor_ID, &hc.Duration, &hc.Address); err != nil {
		return nil, err
	}
	return hc, nil
}

func DeleteHomeCourse(id string) error {
	_, err := database.Db.Exec("DELETE FROM home_courses WHERE id = ?", id)
	return err
}

func (hc *HomeCourse) Save() error {
	if err := hc.Validate(); err != nil {
		return err
	}

	res, err := database.Db.Exec("INSERT INTO home_courses (Title, Description, Price, Start_Time, End_Time, Instructor_ID, Duration, Address) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		hc.Title, hc.Description, hc.Price, hc.Start_Time, hc.End_Time, hc.Instructor_ID, hc.Duration, hc.Address)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	hc.ID = int(id)
	return nil
}

func (hc *HomeCourse) Validate() error {
	if hc.Title == "" {
		return errors.New("title cannot be empty")
	}
	if hc.Description == "" {
		return errors.New("description cannot be empty")
	}
	if hc.Instructor_ID == 0 {
		return errors.New("instructor_id must be greater than 0")
	}
	if hc.Address == "" {
		return errors.New("address cannot be empty")
	}
	if hc.Duration <= 0 {
		return errors.New("duration must be greater than 0")
	}
	if hc.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if hc.Start_Time == "" || hc.End_Time == "" {
		return errors.New("start_time and end_time must be valid dates")
	}
	return nil
}

func ConvertIntToDuration(hours int) time.Duration {
	duration := time.Duration(hours) * time.Hour
	return duration
}
