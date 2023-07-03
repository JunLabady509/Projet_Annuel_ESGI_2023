package user

import (
	"fmt"
	"log"

	"gastroguru/database"

	_ "github.com/go-sql-driver/mysql"
)

// Les utilisateurs sont stockés dans une base de données MySQL
type User struct {
	ID              int    `json:"id"`
	Last_Name       string `json:"last_name"`
	First_Name      string `json:"first_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Picture         string `json:"picture"`
	Role            string `json:"role"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	Subscription_ID int    `json:"subscription_id"`
	Loyalty_Points  int    `json:"loyalty_id"`
}

// All retrieves all users from the database
func All() ([]User, error) {
	rows, err := database.Db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Last_Name, &u.First_Name, &u.Email, &u.Password,
			&u.Picture, &u.Role, &u.Address, &u.Phone, &u.Subscription_ID, &u.Loyalty_Points); err != nil {
			fmt.Println("Error Scanning Rows :", err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// One returns a single user record from the database
func One(id string) (*User, error) {
	u := &User{}
	if err := database.Db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&u.ID, &u.Last_Name, &u.First_Name, &u.Email, &u.Password,
		&u.Picture, &u.Role, &u.Address, &u.Phone, &u.Subscription_ID, &u.Loyalty_Points); err != nil {
		return nil, err
	}
	return u, nil
}

// Delete removes a given user record from the database
func Delete(id string) error {
	_, err := database.Db.Exec("DELETE FROM users WHERE id = ?", id)
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
	res, err := database.Db.Exec("INSERT INTO users (Last_Name, First_Name, Email, Picture, Password, Role, Address, Phone, Subscription_ID, Loyalty_Points ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		u.Last_Name, u.First_Name, u.Email, u.Picture, u.Password, u.Role, u.Address, u.Phone, u.Subscription_ID, u.Loyalty_Points)
	if err != nil {
		fmt.Println("Error Inserting User :", err)
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = int(id)

	return nil
}

// Validate makes sure that the record contains valid data
func (u *User) Validate() error {
	if u.Last_Name == "" ||
		u.First_Name == "" || u.Email == "" ||
		u.Password == "" {
		return database.ErrRecordInvalid
	}
	return nil
}

func FindUserByEmail(email string) (*User, error) {
	u := &User{}
	if err := database.Db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&u.ID, &u.Last_Name, &u.First_Name, &u.Email, &u.Password, &u.Picture, &u.Role, &u.Address, &u.Phone, &u.Subscription_ID, &u.Loyalty_Points); err != nil {
		return nil, err
	}

	return u, nil
}
