package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// InitDB initializes and returns a database connection.
func connect(username, password, host, port, databaseName string) (*sql.DB, error) {
	// Build the data source name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, databaseName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL database!")
	return db, nil
}

func testConnection(username, password, host, port, databaseName string) {

	// Build the data source name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, databaseName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Connected to MySQL database!")
}
