package postgresql

import (
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
)

var db *pg.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db = pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})
}

// DB returns postgres db
func DB() *pg.DB {
	return db
}
