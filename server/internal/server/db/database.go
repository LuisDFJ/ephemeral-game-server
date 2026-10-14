package db

import (
	"sync"
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
)

type DB struct {
	mu sync.Mutex
	db *sql.DB
}

func New() (*DB,error) {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil { return nil, err }
	d := &DB{db:db}
	err = d.populate_database()
	if err != nil { return nil, err }
	return d,nil
}

func (db *DB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.db.Close()
}

func (d *DB) populate_database() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	sql_ott := `CREATE TABLE IF NOT EXISTS ott (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		token TEXT NOT NULL UNIQUE
	);`
	sql_user := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	_,err := d.db.Exec(sql_ott)
	if err != nil { return err }
	_,err = d.db.Exec(sql_user)
	return err
}

