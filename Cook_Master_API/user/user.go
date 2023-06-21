package user

import (
	"fmt"
	"log"

	"gastroguru/database"
	db "gastroguru/database"

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
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Subscribed_ID int    `json:"subscribed_id"`
	Loyality_ID   int    `json:"loyality_id"`
}

// Partie à implémenter pour les permissions
// Il n'y a pas de table permissions dans la base de données
// Peut être qu'on sera amené à modifier la table users et la structure User
// pour implémenter les permissions
// à méditer

type Permission struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// All retrieves all users from the database
func All() ([]User, error) {
	rows, err := db.Db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Role,
			&u.FirstName, &u.Address, &u.Email, &u.Password,
			&u.Phone, &u.Subscribed_ID, &u.Loyality_ID); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// One returns a single user record from the database
func One(id string) (*User, error) {
	u := &User{}
	if err := db.Db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Role, &u.FirstName, &u.Address, &u.Email, &u.Phone, &u.Subscribed_ID, &u.Loyality_ID, &u.Password); err != nil {
		return nil, err
	}
	return u, nil
}

// Delete removes a given user record from the database
func Delete(id string) error {
	_, err := db.Db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

// Save updates or creates a given user record in the database
func (u *User) Save() error {
	if err := u.Validate(); err != nil {
		return err
	}
	var err error
	u.Password, err = database.HashPassword(u.Password)
	if err != nil {
		fmt.Println("Error Hashing Password :", err)
		log.Fatal(err)
	}
	res, err := db.Db.Exec("INSERT INTO users (FirstName, Name, Email, Password, Role, Address, Phone, Subscribed_ID, Loyality_ID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		u.FirstName, u.Name, u.Email, u.Password, u.Role, u.Address, u.Phone, u.Subscribed_ID, u.Loyality_ID)
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
		return db.ErrRecordInvalid
	}
	return nil
}

func (u *User) FindUserByEmail(email string) (*User, error) {
	u = &User{}
	if err := db.Db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&u.ID, &u.Name, &u.Role, &u.FirstName, &u.Address, &u.Email, &u.Phone, &u.Subscribed_ID, &u.Loyality_ID, &u.Password); err != nil {
		return nil, err
	}
	return u, nil
}
