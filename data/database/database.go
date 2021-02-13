package database

import (
	"fmt"
	"log"
	"os"

	"github.com/asdine/storm/v3"
	"github.com/joho/godotenv"
)

// Database ..
type Database struct {
	Conn storm.DB
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// New returns a new instance of the database
func New() *Database {
	env := os.Getenv("ENVIRONMENT")
	connection, err := storm.Open(fmt.Sprintf("./data/database/%s.db", env))
	if err != nil {
		log.Fatal(err.Error())
	}

	return &Database{
		Conn: *connection,
	}
}
