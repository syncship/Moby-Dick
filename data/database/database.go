package database

import (
	"fmt"
	"log"
	"path"
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
	if env == "" {
		env = "PROD"
	}

	p := path.Join(os.Getenv("DB_DIR"), fmt.Sprintf("%s.db", env))
	connection, err := storm.Open(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &Database{
		Conn: *connection,
	}
}
