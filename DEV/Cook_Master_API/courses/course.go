package course

import (
	"errors"
	"gastroguru/database"
	"log"
	"time"
)

type Course struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Type          string  `json:"type"`
	Instructor_ID int     `json:"instructor_id"`
	Location      string  `json:"location"`
	Capacity      int     `json:"capacity"`
	Price         float64 `json:"price"`
	Start_Time    string  `json:"start_time"`
	End_Time      string  `json:"end_time"`
}

func GetAllCourses() ([]Course, error) {
	database.CreateCoursesTable(database.Db)

	rows, err := database.Db.Query("SELECT * FROM courses")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	courses := []Course{}

	for rows.Next() {
		var c Course
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.Type,
			&c.Instructor_ID, &c.Location, &c.Capacity, &c.Price, &c.Start_Time, &c.End_Time); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func GetCourse(id string) (*Course, error) {
	c := &Course{}
	if err := database.Db.QueryRow("SELECT * FROM courses WHERE id = ?", id).Scan(&c.ID, &c.Title, &c.Description, &c.Type, &c.Instructor_ID, &c.Location, &c.Capacity, &c.Price, &c.Start_Time, &c.End_Time); err != nil {
		return nil, err
	}
	return c, nil
}

func DeleteCourse(id string) error {
	_, err := database.Db.Exec("DELETE FROM courses WHERE id = ?", id)
	return err
}

func (c *Course) Save() error {
	if err := c.Validate(); err != nil {
		return err
	}

	res, err := database.Db.Exec("INSERT INTO courses (Title, Description, Type, Instructor_ID, Location, Capacity, Price, Start_Time, End_Time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		c.Title, c.Description, c.Type, c.Instructor_ID, c.Location, c.Capacity, c.Price, c.Start_Time, c.End_Time)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	c.ID = int(id)
	return nil
}

func (c *Course) Validate() error {
	var startTIME, endTIME time.Time
	startTIME, _ = c.GetStringTimeAsTime(c.Start_Time)
	endTIME, _ = c.GetStringTimeAsTime(c.End_Time)
	if c.Title == "" {
		return errors.New("title cannot be empty")
	}
	if c.Description == "" {
		return errors.New("description cannot be empty")
	}
	if c.Type != "on_site" && c.Type != "at_home" && c.Type != "online" {
		return errors.New("type can only be 'on_site', 'at_home', or 'online'")
	}
	if c.Instructor_ID <= 0 {
		return errors.New("instructor_id must be greater than 0")
	}
	if c.Location == "" {
		return errors.New("location cannot be empty")
	}
	if c.Capacity <= 0 {
		return errors.New("capacity must be greater than 0")
	}
	if c.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if startTIME.IsZero() || endTIME.IsZero() {
		return errors.New("start_time and end_time must be valid dates")
	}
	if endTIME.Before(startTIME) || endTIME.Equal(startTIME) {
		return errors.New("end_time must be after start_time")
	}
	return nil
}

func (c *Course) GetStringTimeAsTime(strTime string) (time.Time, error) {
	return time.Parse(time.RFC3339, strTime)
}
