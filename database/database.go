package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "postgresql://expense-tracker_owner:JkF3Vi9txQwe@ep-small-boat-a1rtpu4w.ap-southeast-1.aws.neon.tech/expense-tracker?sslmode=require"
	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	// Check if the database is reachable
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the DB: ", err)
	}

	// Query the database version
	row := DB.QueryRow("SELECT version()")
	var version string
	err = row.Scan(&version)
	if err != nil {
		log.Fatal("Failed to query database version: ", err)
	}

	fmt.Printf("Connected to PostgreSQL version: %s\n", version)
}
