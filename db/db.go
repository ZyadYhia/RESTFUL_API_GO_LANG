package db

// _ is used to prevent editor from removing packages
// as some packages may be used indirectly
import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Error with creating connection") // panic() with GIN will not crash the run but extract log
	}
	DB.SetMaxOpenConns(10) // Max opening connections together
	DB.SetMaxIdleConns(5)  // Max opening connection if no ongoing communication with DB
	CreateTables()
	fmt.Println("Initialization DB Done!!")
}

func CreateTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
		)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println("Error Table: ", err)
		panic("Error with creating Users table") // panic() with GIN will not crash the run but extract log
	}
	createEventTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		fmt.Println("Error Table: ", err)
		panic("Error with creating Events table") // panic() with GIN will not crash the run but extract log
	}

	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		fmt.Println("Error Table: ", err)
		panic("Error with creating Registeration table") // panic() with GIN will not crash the run but extract log
	}
}
