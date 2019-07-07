package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDb() {
	log.Println("Initializing database...")

	db, _ = sql.Open("sqlite3", "./data.db")
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS servers (
			id INTEGER PRIMARY KEY,
			game_id VARCHAR(256),
			name VARCHAR(256),
			host VARCHAR(256),
			port INTEGER,
			num_players INTEGER,
			max_players INTEGER
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

func getServersByGameId(gid GameIDType) []GameServer {
	ret := make([]GameServer, 0)
	ret = append(ret, GameServer{Name: "Test"})
	return ret
}

func getServerById(id ServerIDType) GameServer {
	return _db[id]
}

func insertServer(server GameServer) (GameServer, error) {
	server.ID = nextID()
	_db[server.ID] = server
	return server, nil
}

func updateServerById(id ServerIDType, server GameServer) (GameServer, error) {
	curr, ok := _db[id]
	if !ok {
		return curr, errors.New("No server with ID type: " + string(id))
	}
	err := curr.UpdateFrom(&server)
	if err != nil {
		return curr, err
	}
	return curr, nil
}

func pingServerAlive(id ServerIDType) {

}
