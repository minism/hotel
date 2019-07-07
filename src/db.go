package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initDb() {
	log.Println("Initializing database...")

	db, _ := sql.Open("sqlite3", "./data.db")
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS servers (
			id INTEGER PRIMARY KEY
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
