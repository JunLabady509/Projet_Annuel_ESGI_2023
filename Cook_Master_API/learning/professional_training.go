package learning

import (
	"errors"
	"fmt"
	"gastroguru/database"
	"log"
	"time"
)

type ProfessionalTraining struct {
	Learning
	Duration    time.Duration `json:"duration"`
	Capacity    int           `json:"capacity"`
	Schedule    string        `json:"schedule"`
	Certificate bool          `json:"certificate"`
	Place       string        `json:"place"`
}

func GetAllProfessionalTrainings() ([]ProfessionalTraining, error) {
	rows, err := database.Db.Query("SELECT * FROM professional_trainings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	trainings := []ProfessionalTraining{}

	for rows.Next() {
		var t ProfessionalTraining
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Instructor_ID, &t.Duration, &t.Capacity, &t.Schedule, &t.Certificate, &t.Place); err != nil {
			return nil, err
		}
		trainings = append(trainings, t)
	}

	return trainings, nil
}

func GetProfessionalTraining(id string) (*ProfessionalTraining, error) {
	t := &ProfessionalTraining{}
	if err := database.Db.QueryRow("SELECT * FROM professional_trainings WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Description, &t.Instructor_ID, &t.Duration, &t.Capacity, &t.Schedule, &t.Certificate, &t.Place); err != nil {
		return nil, err
	}
	return t, nil
}

func DeleteProfessionalTraining(id string) error {
	_, err := database.Db.Exec("DELETE FROM professional_trainings WHERE id = ?", id)
	return err
}

func (t *ProfessionalTraining) Save() error {
	if err := t.Validate(); err != nil {
		return err
	}

	res, err := database.Db.Exec("INSERT INTO professional_trainings (Title, Description, Instructor_ID, Duration, Capacity, Schedule, Certificate, Place) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		t.Title, t.Description, t.Instructor_ID, t.Duration, t.Capacity, t.Schedule, t.Certificate, t.Place)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	t.ID = int(id)
	return nil
}

func (t *ProfessionalTraining) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	if t.Description == "" {
		return errors.New("description cannot be empty")
	}
	if t.Instructor_ID == 0 {
		return errors.New("instructor_id must be greater than 0")
	}
	if t.Duration <= 0 {
		return errors.New("duration must be greater than 0")
	}
	if t.Capacity <= 0 {
		return errors.New("capacity must be greater than 0")
	}
	if t.Schedule == "" {
		return errors.New("schedule cannot be empty")
	}
	if t.Place == "" {
		return errors.New("place cannot be empty")
	}
	return nil
}

func (t *ProfessionalTraining) GetStringDurationAsDuration(strDuration string) (time.Duration, error) {
	return time.ParseDuration(strDuration)
}
