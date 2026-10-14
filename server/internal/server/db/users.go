package db

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) CreateUser(username string, password string) error {
	hash,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil { return err }
	return db.insert_user(username,string(hash))
}

func (db *DB) insert_user( username string, hash string ) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	sql := `INSERT INTO users (username, password) VALUES (?,?);`
	_, err := db.db.Exec(sql, username, hash)
	return err
}

var AuthErr = fmt.Errorf("auth error")

func (db *DB) AuthUser(username string, password string) error {
	hash := db.get_user_hash(username)
	if hash == "" { return AuthErr }
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil { return AuthErr }
	return nil
}

func (db *DB) get_user_hash( username string ) string {
	db.mu.Lock()
	defer db.mu.Unlock()
	sql := `SELECT password FROM users WHERE users.username = ?;`
	res := db.db.QueryRow(sql, username)
	var hash string
	res.Scan(&hash)
	return hash
}
