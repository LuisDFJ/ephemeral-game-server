package db

import (
	"time"
	"errors"
	sq "database/sql"
	_ "github.com/glebarez/go-sqlite"
)

func (db *DB) GetLastLog() (int64, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.get_log()
}

func (db *DB) get_log() (int64,error) {
	sql := `SELECT timestamp FROM vps_log ORDER BY timestamp DESC LIMIT 1;`
	row := db.db.QueryRow(sql)
	var timestamp int64
	err := row.Scan(&timestamp)
	if errors.Is(err, sq.ErrNoRows) {
		return db.insert_log("host")
	}
	return timestamp, err
}

func (db *DB) InsertLog(user string) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	_,err := db.insert_log(user)
	return err
}

func (db *DB) insert_log( user string ) (int64, error) {
	sql := `INSERT INTO vps_log (username, timestamp) VALUES (?,?);`
	timestamp := time.Now().Unix()
	_, err := db.db.Exec(sql, user, timestamp)
	return timestamp, err
}
