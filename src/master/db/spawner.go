package db

import (
	"errors"
	"log"

	"minornine.com/hotel/src/master/models"
	hotel_pb "minornine.com/hotel/src/proto"
	"minornine.com/hotel/src/shared"
)

func GetSpawners() []models.Spawner {
	return spawnerQuery("")
}

func GetSpawnersByGameId(gid shared.GameIDType) []models.Spawner {
	return spawnerQuery("WHERE game_id = ?", gid)
}

func InsertSpawner(spawner models.Spawner) error {
	stmt, err := db.Prepare(`
		INSERT INTO spawners (game_id, host, port, num_game_servers, max_game_servers)
		VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to insert spawner")
	}
	_, err = stmt.Exec(
		spawner.GameID, spawner.Host, spawner.Port, spawner.NumGameServers, spawner.MaxGameServers)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to insert spawner")
	}
	return nil
}

func UpdateSpawnerFromStatus(id int, status *hotel_pb.SpawnerStatus) error {
	stmt, err := db.Prepare(`
		UPDATE spawners
		SET game_id = ?,
			num_game_servers = ?,
			max_game_servers = ?
		WHERE
			id = ?
	`)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to update server.")
	}
	_, err = stmt.Exec(
		status.SupportedGameId, status.NumGameServers, status.MaxGameServers, id)
	if err != nil {
		log.Println(err)
		errors.New("Failed to update server.")
	}
	return nil
}

func DeleteSpawnerById(id int) error {
	stmt, err := db.Prepare(`
		DELETE FROM spawners
		WHERE id = ?
	`)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to delete spawner.")
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to delete spawner.")
	}
	return nil
}

func spawnerQuery(where string, args ...interface{}) []models.Spawner {
	ret := make([]models.Spawner, 0)
	q := "SELECT id, game_id, host, port, num_game_servers, max_game_servers FROM spawners"
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
	var s models.Spawner
	for rows.Next() {
		rows.Scan(&s.ID, &s.GameID, &s.Host, &s.Port, &s.NumGameServers, &s.MaxGameServers)
		ret = append(ret, s)
	}
	return ret
}
