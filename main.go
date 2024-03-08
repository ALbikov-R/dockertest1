package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbHost  = "cockroachdb"
	dbPort  = 26257
	dbUser  = "root"
	dbName  = "test_db"
	sslMode = "disable"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbName, sslMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	defer db.Close()

	// Ensure the database exists and set the context (This step might be redundant if the dbName is correctly specified in the connection string)
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName))
	if err != nil {
		log.Fatal("error ensuring the database exists: ", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Users (
        ID SERIAL PRIMARY KEY,
        Name TEXT,
        Email TEXT UNIQUE NOT NULL
    )`)
	if err != nil {
		log.Fatal("error creating Users table: ", err)
	}

	fmt.Println("Successfully connected and created Users table if not exists.")
}
