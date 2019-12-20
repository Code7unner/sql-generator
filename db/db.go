package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	CreateQuery = `
		DROP TABLE IF EXISTS public.categories CASCADE;
		DROP TABLE IF EXISTS public.users CASCADE;
		DROP TABLE IF EXISTS public.messages CASCADE;
		
		CREATE UNLOGGED TABLE IF NOT EXISTS public.messages (
			"id" uuid NOT NULL,
			"text" TEXT NOT NULL,
			"category_id" uuid NOT NULL,
			"posted_at" TIME NOT NULL,
			"author_id" uuid NOT NULL
		) WITH (
			OIDS=FALSE
		);
		
		CREATE UNLOGGED TABLE IF NOT EXISTS public.categories (
		   "id" uuid NOT NULL,
		   "name" varchar(255) NOT NULL,
		   "parent_id" uuid NOT NULL
		) WITH (
			OIDS=FALSE
		);
		
		CREATE UNLOGGED TABLE IF NOT EXISTS public.users (
			"id" uuid NOT NULL,
			"name" varchar(255) NOT NULL
		) WITH (
			OIDS=FALSE
		);
		
		ALTER TABLE public.users SET (autovacuum_enabled = false);
		ALTER TABLE public.categories SET (autovacuum_enabled = false);
		ALTER TABLE public.messages SET (autovacuum_enabled = false);
	`
)

type Postgres struct {
	DB *sql.DB
}

type Config struct {
	dbUser string
	dbPass string
	dbName string
}

func InitDatabase() (*Postgres, error) {
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
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, err
}

func (p *Postgres) CreateTablesQuery() error {
	transaction, err := p.DB.Begin()
	if err != nil {
		return err
	}

	_, err = transaction.Exec(CreateQuery)
	if err != nil {
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	return nil
}

