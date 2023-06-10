package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Les utilisateurs sont stockés dans une base de données MySQL
type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Role          string `json:"role"`
	FirstName     string `json:"first_name"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Subscribed_ID int    `json:"subscribed_id"`
	Loyality_ID   int    `json:"loyality_id"`
}

const (
	DBDriver   = "mysql"
	DBName     = "gastroguru"
	DBUser     = "junior"
	DBPassword = "junior"
	dbPath     = "users.db"
)

var (
	ErrRecordInvalid = errors.New("record is invalid")
	db               *sql.DB
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@/%s", DBUser, DBPassword, DBName)
	db, err = sql.Open(DBDriver, dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

// All retrieves all users from the database
func All() ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Role,
			&u.FirstName, &u.Address, &u.Email,
			&u.Phone, &u.Subscribed_ID,
			&u.Loyality_ID); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// One returns a single user record from the database
func One(id string) (*User, error) {
	u := &User{}
	if err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Role, &u.FirstName, &u.Address, &u.Email, &u.Phone, &u.Subscribed_ID, &u.Loyality_ID); err != nil {
		return nil, err
	}
	return u, nil
}

// Delete removes a given user record from the database
func Delete(id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

// Save updates or creates a given user record in the database
func (u *User) Save() error {
	if err := u.Validate(); err != nil {
		return err
	}

	res, err := db.Exec("INSERT INTO users (FirstName, Name, Email, Role, Address, Phone, Subscribed_ID, Loyality_ID) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		u.FirstName, u.Name, u.Email, u.Role, u.Address, u.Phone, u.Subscribed_ID, u.Loyality_ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	u.ID = int(id)

	return nil
}

// Validate makes sure that the record contains valid data
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
