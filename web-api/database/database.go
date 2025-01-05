package database

// I use standard database/sql database API with Postgres PGX adapter
// it's less performant, but provides flexibility over real drivers you use behind the scenes
// another option is to use "github.com/jackc/pgx" API directly, it's more performant
import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib" // Import the adapter
)

// DB global variable of Database to be used elsewhere in this application
var DB *sql.DB

func InitDatabase() {
	var password = "<SET YOUR DATABASE PASSWORD HERE>"
	var err error

	var connectionStringPostgres = fmt.Sprintf("postgres://postgres:%s@localhost:5432/gin-web-api", password)

	DB, err = sql.Open("pgx", connectionStringPostgres)
	if err != nil {
		// even if we panic here the server will not crash because of Gin
		panic("Connection with the database cannot be established")
	}

	// number of connections we keep opened while no one uses them
	DB.SetMaxIdleConns(5)

	// number of connections could simultaneously be opened
	DB.SetMaxOpenConns(10)

	// we state that connection should be closed once the function context ended (it's ended at "}" of this function)
	defer DB.Close()

	// migrations
	createTables()
}

func createTables() {
	var createTableEvents = `
		CREATE TABLE IF NOT EXISTS events (
                                      id SERIAL PRIMARY KEY,
                                      name TEXT NOT NULL,
                                      description TEXT NOT NULL,
                                      location TEXT NOT NULL,
                                      dateTime TIMESTAMP NOT NULL,
                                      user_id INTEGER
		)
	`

	_, err := DB.Exec(createTableEvents)
	if err != nil {
		panic(fmt.Sprintf("Migration has not been applied properly", err))
	}
}
