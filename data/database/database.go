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
	godotenv.Load()

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatalln("Could not find value for 'DISCORD_TOKEN', make sure you have a .env file or have setup the environment variables properly.")
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
