package db

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDb initializes the database at the given path and ensures tables are created.
func InitDb(dataPath string) {
	dbPath := filepath.Join(dataPath, "data.db")
	log.Printf("Initializing database at %v...", dbPath)

	db, _ = sql.Open("sqlite3", dbPath)
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS servers (
			id INTEGER PRIMARY KEY,
			game_id VARCHAR(256),
			session_id INTEGER,
			name VARCHAR(256),
			host VARCHAR(256),
			port INTEGER,
			num_players INTEGER,
			max_players INTEGER,
			last_modified INTEGER
		);

		CREATE TABLE IF NOT EXISTS spawners (
			id INTEGER PRIMARY KEY, 
			game_id VARCHAR(256),
			host VARCHAR(256),
			port INTEGER,
			num_game_servers INTEGER,
			max_game_servers INTEGER
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		"CREATE INDEX IF NOT EXISTS idx_servers_game_id ON SERVERS(game_id)")
	if err != nil {
		log.Fatal(err)
	}
}
