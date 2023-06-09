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

// Vérifie si un utilisateur existe dans la base de données en fonction de son email
func UserExists(email string) (bool, error) {
	var count int
	err := Db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// La base de données est cryptée avec bcrypt, donc on ne peut vérifier si un mot de passe correspond à un email qu'en utilisant la fonction CheckPassword de bcrypt
func PasswordMatchesEmail(password string, email string) (bool, error) {
	var hashedPassword string
	err := Db.QueryRow("SELECT Password FROM users WHERE email = ?", email).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateAllTables(db *sql.DB) {
	// Création de la table User
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		email VARCHAR(255),
		password VARCHAR(255),
		picture VARCHAR(255),
		role VARCHAR(255),
		address VARCHAR(255),
		phone VARCHAR(255),
		subscription_id INT,
		loyalty_points INT,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS workshops (
		id INT AUTO_INCREMENT,
		title VARCHAR(255),
		description TEXT, 
		instructor_id INT,
		capacity INT,
		price DECIMAL(5,2),
		start_time DATETIME,
		end_time DATETIME,
		place VARCHAR(255),
		insight VARCHAR(255),
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
		price DECIMAL(5,2),
		start_time DATETIME,
		end_time DATETIME,
		instructor_id INT,
		Duration INT,
		Address VARCHAR(255),
		PRIMARY KEY (id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS online_lessons (
			id INT PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR(255),
			description TEXT,
			price DECIMAL(5, 2),
			start_Time DATETIME,
			end_Time DATETIME,
			instructor_ID INT,
			video_Link JSON,
			uploaded_Time DATETIME,
			insight TEXT
	);
	`)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS professional_trainings (
			ID INT AUTO_INCREMENT PRIMARY KEY,
			Title VARCHAR(255) NOT NULL,
			Description TEXT NOT NULL,
			Instructor_ID INT NOT NULL,
			Duration INT NOT NULL,
			Capacity INT NOT NULL,
			Schedule VARCHAR(255) NOT NULL,
			Certificate BOOL NOT NULL,
			Place VARCHAR(255) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS events (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    date DATETIME NOT NULL,
    duration INT NOT NULL,
    location VARCHAR(255),
    capacity INT,
    price DECIMAL(10, 2),
    image VARCHAR(255),
    type VARCHAR(255)
);`)

	if err != nil {
		log.Fatal(err)
	}

}
