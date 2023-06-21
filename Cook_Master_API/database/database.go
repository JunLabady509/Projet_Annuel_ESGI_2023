package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/bcrypt"
)

const (
	DBDriver           = "mysql"
	DBName             = "gastroguru"
	DBUser             = "junior"
	DBPassword         = "junior"
	SubscriptionDbPath = "subscription.db"
	UserDbPath         = "user.db"
)

var (
	ErrRecordInvalid = errors.New("record is invalid")
	Db               *sql.DB
)

func InitDB() {
	var err error

	dataSourceName := fmt.Sprintf("%s:%s@/%s", DBUser, DBPassword, "")
	Db, err = sql.Open(DBDriver, dataSourceName)
	if err != nil {
		panic(err)
	}
	// Création de la base de données
	_, err = Db.Exec("CREATE DATABASE IF NOT EXISTS gastroguru")
	if err != nil {
		log.Fatal(err)
	}

	// Sélection de la base de données
	_, err = Db.Exec("USE " + DBName)
	if err != nil {
		log.Fatal(err)
	}
	if err = Db.Ping(); err != nil {
		panic(err)
	}
	log.Println("You connected to your database.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Db.Close()
		os.Exit(0)
	}()

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Compare un mot de passe Hashé avec un mot de passe non hashé
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateAllTables(db *sql.DB) {
	// Création de la table User
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		ID INT AUTO_INCREMENT,
		Name VARCHAR(255) NOT NULL,
		Role VARCHAR(255) NOT NULL,
		FirstName VARCHAR(255) NOT NULL,
		Address VARCHAR(255) NOT NULL,
		Email VARCHAR(255) NOT NULL,
		Password VARCHAR(255) NOT NULL,
		Phone VARCHAR(255) NOT NULL,
		Subscribed_ID INT NOT NULL,
		Loyality_ID INT NOT NULL,
		PRIMARY KEY (ID)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS workshops (
		id INT AUTO_INCREMENT,
		title VARCHAR(255),
		description TEXT, 
		instructor_id INT,
		location VARCHAR(100),
		capacity INT,
		price DECIMAL(5,2),
		start_time DATETIME,
		end_time DATETIME,
		PRIMARY KEY (id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS home_courses (
		id INT AUTO_INCREMENT,
		title VARCHAR(255),
		description TEXT, 
		instructor_id INT,
		client_id INT,
		location VARCHAR(100),
		price DECIMAL(5,2),
		start_time DATETIME,
		end_time DATETIME,
		PRIMARY KEY (id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS online_lessons (
		id INT AUTO_INCREMENT,
		title VARCHAR(255),
		description TEXT,
		instructor_id INT,
		price DECIMAL(5,2),
		video_link VARCHAR(255),
		uploaded_time DATETIME,
		PRIMARY KEY (id)
	);`)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
