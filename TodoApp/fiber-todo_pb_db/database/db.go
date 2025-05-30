package database
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)
var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to the database successfully")
}