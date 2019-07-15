package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

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
			session_id INTEGER,
			name VARCHAR(256),
			host VARCHAR(256),
			port INTEGER,
			num_players INTEGER,
			max_players INTEGER,
			last_modified INTEGER
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

// TODO: Make private once package split happens
func serverQuery(where string, args ...interface{}) []GameServer {
	ret := make([]GameServer, 0)
	q := "SELECT id, game_id, session_id, name, host, port, num_players, max_players FROM servers"
	if len(where) > 0 {
		q = q + " " + where
	}
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
		return ret
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		log.Println(err)
		return ret
	}
	var s GameServer
	for rows.Next() {
		rows.Scan(&s.ID, &s.GameID, &s.SessionID, &s.Name, &s.Host, &s.Port, &s.NumPlayers, &s.MaxPlayers)
		ret = append(ret, s)
	}
	return ret
}

func getServersByGameId(gid GameIDType) []GameServer {
	return serverQuery("WHERE game_id = ?", gid)
}

func getServerById(id ServerIDType) (GameServer, bool) {
	var ret GameServer
	servers := serverQuery("WHERE id = ?", id)
	if len(servers) > 0 {
		return servers[0], true
	}
	return ret, false
}

func insertServer(server GameServer) (GameServer, error) {
	stmt, err := db.Prepare(`
		INSERT INTO servers (game_id, session_id, name, host, port, num_players, max_players, last_modified)
		VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err)
		return server, errors.New("Failed to create server.")
	}
	res, err := stmt.Exec(
		server.GameID, server.SessionID, server.Name, server.Host, server.Port,
		server.NumPlayers, server.MaxPlayers, getModifiedTime())
	if err != nil {
		log.Println(err)
		return server, errors.New("Failed to create server.")
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return server, errors.New("Unknown fatal error retrieving server ID.")
	}
	server.ID = ServerIDType(id)
	return server, nil
}

func updateServerById(id ServerIDType, server GameServer) (GameServer, error) {
	stmt, err := db.Prepare(`
		UPDATE servers
		SET name = ?,
			host = ?,
			port = ?,
			num_players = ?,
			max_players = ?
			last_modified = ?
		WHERE
			id = ?
	`)
	if err != nil {
		log.Println(err)
		return server, errors.New("Failed to update server.")
	}
	_, err = stmt.Exec(
		server.Name, server.Host, server.Port, server.NumPlayers, server.MaxPlayers,
		getModifiedTime(), id)
	if err != nil {
		log.Println(err)
		return server, errors.New("Failed to update server.")
	}
	return server, nil
}

func updateServerAlive(id ServerIDType) error {
	stmt, err := db.Prepare(`
		UPDATE SERVERS
		SET last_modified = ?
		WHERE id = ?
	`)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to update server.")
	}
	_, err = stmt.Exec(getModifiedTime(), id)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to update server.")
	}
	return nil
}

func deleteServersOlderThan(timestamp int64) {
}

func getModifiedTime() int64 {
	return time.Now().Unix()
}
