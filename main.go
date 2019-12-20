package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

type User struct {
	Id   uuid.UUID
	Name string
}

type Category struct {
	Id       uuid.UUID
	Name     string
	ParentId uuid.UUID
}

type Message struct {
	Id         uuid.UUID
	Text       string
	CategoryId uuid.UUID
	PostedAt   time.Time
	AuthorId   uuid.UUID
}

type Store interface {
	//CreateUsers() 
	//CreateCategories()
	//CreateMessages()
}
type Postgres struct {
	DB *sql.DB
}

func main()  {
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("Error initialize database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}

	defer db.Close()
}

type Config struct {
	dbUser string
	dbPass string
	dbName string
}

func initDatabase() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := Config{
		dbUser : os.Getenv("db_user"),
		dbPass : os.Getenv("db_pass"),
		dbName : os.Getenv("db_name"),
	}

	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		c.dbUser, c.dbPass, c.dbName)

	db, err := sql.Open("postgres", psqlInfo)

	return db, err
}
