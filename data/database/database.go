package database

import (
	"log"

	"github.com/asdine/storm/v3"
)

// Database ..
type Database struct {
	Conn storm.DB
}

// New returns a new instance of the database
func New() *Database {
	connection, err := storm.Open("./data/database/dev.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	return &Database{
		Conn: *connection,
	}
}
