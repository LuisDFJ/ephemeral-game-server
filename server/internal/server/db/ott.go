package db

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func (db *DB) GenerateOTT() (string, error) {
	buf := make([]byte,256)
	_,err := io.ReadFull(rand.Reader,buf)
	if err != nil { return "", err }
	ott := base64.StdEncoding.EncodeToString(buf)
	err = db.insert_ott(ott)
	if err != nil { return "", err }
	return ott, nil
}

func (db *DB) insert_ott( ott string ) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	sql := `INSERT INTO ott (token) VALUES (?);`
	_, err := db.db.Exec(sql, ott)
	return err
}

func (db *DB) ValidateOTT(ott string) int {
	return db.get_ott(ott)
}

func (db *DB) get_ott( ott string ) int {
	db.mu.Lock()
	defer db.mu.Unlock()
	sql := `SELECT id FROM ott WHERE ott.token = ?;`
	row := db.db.QueryRow(sql, ott)
	var id int
	row.Scan(&id)
	return id
}

func (db *DB) DeleteOTT(id int) error {
	return db.del_ott(id)
}

func (db *DB) del_ott( id int ) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	sql := `DELETE FROM ott WHERE ott.id = ?`
	_,err := db.db.Exec(sql, id)
	return err
}

