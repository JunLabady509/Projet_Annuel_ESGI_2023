package learning

import (
	"fmt"
	"gastroguru/database"
	"log"
)


type OnlineLesson struct {
	Learning
	Video_Link    string `json:"video_link"`
	Uploaded_Time string `json:"uploaded_time"`
}

func GetAllOnlineLessons() ([]OnlineLesson, error) {
	rows, err := database.Db.Query("SELECT * FROM online_lessons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	online_lessons := []OnlineLesson{}

	for rows.Next() {
		var o OnlineLesson
		if err := rows.Scan(&o.ID, &o.Title, &o.Description,
			&o.Instructor_ID, &o.Price, &o.Video_Link, &o.Uploaded_Time); err != nil {
			return nil, err
		}
		online_lessons = append(online_lessons, o)
	}
	return online_lessons, nil
}

func GetOnlineLesson(id string) (*OnlineLesson, error) {
	o := &OnlineLesson{}
	if err := database.Db.QueryRow("SELECT * FROM online_lessons WHERE id = ?", id).Scan(&o.ID, &o.Title, &o.Description, &o.Instructor_ID, &o.Price, &o.Video_Link, &o.Uploaded_Time); err != nil {
		return nil, err
	}
	return o, nil
}

func DeleteOnlineLesson(id string) error {
	_, err := database.Db.Exec("DELETE FROM online_lessons WHERE id = ?", id)
	return err
}

func (o *OnlineLesson) Save() error {
	if err := o.Validate(); err != nil {
		return err
	}
	res, err := database.Db.Exec("INSERT INTO online_lessons (title, description, instructor_id, price, video_link, uploaded_time) VALUES (?, ?, ?, ?, ?, ?)", o.Title, o.Description, o.Instructor_ID, o.Price, o.Video_Link, o.Uploaded_Time)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}
	o.ID = int(id)

	return nil
}

func (o *OnlineLesson) Validate() error {
	if o.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if o.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	if o.Instructor_ID == 0 {
		return fmt.Errorf("instructor ID cannot be empty")
	}
	if o.Price == 0 {
		return fmt.Errorf("price cannot be empty")
	}
	if o.Video_Link == "" {
		return fmt.Errorf("video Link cannot be empty")
	}
	if o.Uploaded_Time == "" {
		return fmt.Errorf("uploaded Time cannot be empty")
	}
	return nil
}
