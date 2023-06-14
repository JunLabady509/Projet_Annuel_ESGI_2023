package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	dataSourceName := fmt.Sprintf("%s:%s@/%s", DBUser, DBPassword, DBName)
	Db, err = sql.Open(DBDriver, dataSourceName)
	if err != nil {
		panic(err)
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

func CreateCoursesTable(db *sql.DB) {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS courses (
		id INT AUTO_INCREMENT,
		title VARCHAR(255),
		description TEXT,
		type ENUM('on_site', 'at_home', 'online'), 
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

}
