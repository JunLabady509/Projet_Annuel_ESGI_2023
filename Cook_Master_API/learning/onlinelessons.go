package learning

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gastroguru/database"
	"log"
)

type OnlineLesson struct {
	Learning
	Video_Link    []string `json:"video_link"`
	Uploaded_Time string   `json:"uploaded_time"`
	Insight       string   `json:"insight"`
}

func GetAllOnlineLessons() ([]OnlineLesson, error) {
	rows, err := database.Db.Query("SELECT * FROM online_lessons")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	onlineLessons := []OnlineLesson{}

	for rows.Next() {
		var o OnlineLesson

		var tempVideoLink []byte

		if err := rows.Scan(&o.ID, &o.Title, &o.Description, &o.Price, &o.Start_Time, &o.End_Time, &o.Instructor_ID, &tempVideoLink, &o.Uploaded_Time, &o.Insight); err != nil {
			return nil, err
		}

		// Convertir tempVideoLink en []string en utilisant un délimiteur
		Video_Link := bytes.Split(tempVideoLink, []byte{'|'})
		Video_LinkStrings := make([]string, len(Video_Link))
		for i, link := range Video_Link {
			Video_LinkStrings[i] = string(link)
		}
		o.Video_Link = Video_LinkStrings
		onlineLessons = append(onlineLessons, o)
	}

	return onlineLessons, nil
}

func GetOnlineLesson(id string) (*OnlineLesson, error) {
	o := &OnlineLesson{}
	var videoLinkJSON string
	if err := database.Db.QueryRow("SELECT * FROM online_lessons WHERE id = ?", id).Scan(&o.ID, &o.Title, &o.Description, &o.Instructor_ID, &o.Price, &o.Start_Time, &o.End_Time, &videoLinkJSON, &o.Uploaded_Time, &o.Insight); err != nil {
		return nil, err
	}

	err := json.Unmarshal([]byte(videoLinkJSON), &o.Video_Link)
	if err != nil {
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

	tempVideoLink, err := json.Marshal(o.Video_Link)
	if err != nil {
		return err
	}

	res, err := database.Db.Exec("INSERT INTO online_lessons (title, description, instructor_id, price, start_time, end_time, video_link, uploaded_time, insight) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		o.Title, o.Description, o.Instructor_ID, o.Price, o.Start_Time, o.End_Time, tempVideoLink, o.Uploaded_Time, o.Insight)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	o.ID = int(id)
	return nil
}

func (o *OnlineLesson) Validate() error {
	if o.Title == "" {
		return errors.New("title cannot be empty")
	}
	if o.Description == "" {
		return errors.New("description cannot be empty")
	}
	if o.Instructor_ID == 0 {
		return errors.New("instructor_id must be greater than 0")
	}
	if o.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if o.Start_Time == "" || o.End_Time == "" {
		return errors.New("start_time and end_time must be valid dates")
	}
	if o.Uploaded_Time == "" {
		return errors.New("uploaded_time cannot be empty")
	}
	if len(o.Video_Link) == 0 {
		return errors.New("video_link cannot be empty")
	}

	return nil
}
