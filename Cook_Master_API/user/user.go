package user

import (
	"errors"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

// Les utilisateurs sont stockés dans une base de données Storm
type User struct {
	ID            bson.ObjectId `json:"id" storm:"id"`
	Name          string        `json:"name"`
	Role          string        `json:"role"`
	FirstName     string        `json:"first_name"`
	Address       string        `json:"address"`
	Email         string        `json:"email"`
	Phone         string        `json:"phone"`
	Subscribed_ID int           `json:"subscribed_id"`
	Loyality_ID   int           `json:"loyality_id"`
}

const (
	dbPath = "users.db"
)

// Errors utilisés par le package
var (
	ErrRecordInvalid = errors.New("record is invalid")
)

// All retrieves all users from the database
func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil

}

// One returns a single user record from the database
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete removes a given user record from the database
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}

	return db.DeleteStruct(u)
}

// Save updates or creates a given user record in the database
func (u *User) Save() error {
	if err := u.Validate(); err != nil {
		return err
	}

	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Save(u)
}

// Validate makes sure that the record contains valid data
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
