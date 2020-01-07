package db

import (
	"errors"
	"log"
	"time"

	"minornine.com/hotel/src/master/models"
	"minornine.com/hotel/src/shared"
)

func DbGetGameServersByGameId(gid shared.GameIDType) []models.GameServer {
	return serverQuery("WHERE game_id = ?", gid)
}

func DbGetGameServerById(id models.ServerIDType) (models.GameServer, bool) {
	var ret models.GameServer
	servers := serverQuery("WHERE id = ?", id)
	if len(servers) > 0 {
		return servers[0], true
	}
	return ret, false
}

func DbInsertGameServer(server models.GameServer) (models.GameServer, error) {
	stmt, err := db.Prepare(`
		INSERT INTO servers (game_id, session_id, name, host, port, num_players, max_players, last_modified)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
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
	server.ID = models.ServerIDType(id)
	return server, nil
}

func DbUpdateGameServerById(id models.ServerIDType, server models.GameServer) (models.GameServer, error) {
	stmt, err := db.Prepare(`
		UPDATE servers
		SET name = ?,
			host = ?,
			port = ?,
			num_players = ?,
			max_players = ?,
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

func DbDeleteGameServerById(id models.ServerIDType) error {
	stmt, err := db.Prepare(`
		DELETE FROM servers
		WHERE id = ?
	`)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to delete server.")
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to delete server.")
	}
	return nil
}

func DeleteServersOlderThan(timestamp int64) error {
	stmt, err := db.Prepare(`
		DELETE FROM servers
		WHERE last_modified < ?
	`)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to delete servers.")
	}
	_, err = stmt.Exec(timestamp)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to delete servers.")
	}
	return nil
}
func getModifiedTime() int64 {
	return time.Now().Unix()
}

func serverQuery(where string, args ...interface{}) []models.GameServer {
	ret := make([]models.GameServer, 0)
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
	var s models.GameServer
	for rows.Next() {
		rows.Scan(&s.ID, &s.GameID, &s.SessionID, &s.Name, &s.Host, &s.Port, &s.NumPlayers, &s.MaxPlayers)
		ret = append(ret, s)
	}
	return ret
}
